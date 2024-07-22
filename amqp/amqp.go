package amqp

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"sync"
	"time"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type CallbackFunc func([]byte)

type ExchangeValue string
type ExchangeEnt struct {
	Topic  ExchangeValue
	Fanout ExchangeValue
	Direct ExchangeValue
}

var ExchangeType = &ExchangeEnt{
	Topic:  "topic",
	Fanout: "fanout",
	Direct: "direct",
}

type Rabbit struct {
	Connection      *amqp.Connection
	ProducerChannel *amqp.Channel
	ConsumerChannel *amqp.Channel
	BindingRoutes   map[string]CallbackFunc
	Exchange        string
	Queue           string
}

var (
	AmqpClient *Rabbit
)

func ParseENV[T interface{}](object T) error {
	err := viper.Unmarshal(object)
	if err != nil {
		return err
	}
	return nil
}

// step 1
func ConnectRabbit(username, password, host, port, exchange, queue string, exchangeType ExchangeValue) *Rabbit {
	if AmqpClient != nil {
		return AmqpClient
	}

	connectString := fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		username,
		password,
		host,
		port,
	)

	cnf := amqp.Config{
		Properties: amqp.Table{"connection_name": exchange},
	}

	conn, err := amqp.DialConfig(connectString, cnf)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	AmqpClient = &Rabbit{
		Connection:      conn,
		ProducerChannel: InitProducerChannel(conn),
		ConsumerChannel: InitConsumerChannel(conn, exchange, queue, exchangeType), //Service listen event on its exchange, other services send event to that exchange
		Exchange:        exchange,
		Queue:           queue,
	}

	fmt.Println("[ 🚀 ] Connected Successfully to RabbitMQ")
	return AmqpClient
}

func GetConnection() *Rabbit {
	if AmqpClient != nil {
		return AmqpClient
	}
	panic("Rabbit not connected")
}

// step 2
func InitConsumerChannel(connection *amqp.Connection, exchange, queue string, exchangeType ExchangeValue) *amqp.Channel {
	var wg sync.WaitGroup

	channel, err := connection.Channel()
	if err != nil {
		slog.Info(err.Error())
		return nil
	}

	wg.Add(2)
	go DeclareExchange(exchange, exchangeType, channel, &wg)
	go DeClareQueue(queue, channel, &wg)

	wg.Wait()

	return channel
}

func InitProducerChannel(connection *amqp.Connection) *amqp.Channel {
	channel, err := connection.Channel()
	if err != nil {
		slog.Info(err.Error())
		return nil
	}

	return channel
}

// step 3: create exchange with fanout/topic/direct/headers
func DeclareExchange(name string, exchangeType ExchangeValue, channel *amqp.Channel, wg *sync.WaitGroup) error {
	defer wg.Done()

	return channel.ExchangeDeclare(
		name,                 // name
		string(exchangeType), // type
		true,                 // durable(make sure if consumer die, msg isn't lost, need to define in both consumer & producer)
		false,                // auto-deleted
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)
}

// step 3
func DeClareQueue(name string, channel *amqp.Channel, wg *sync.WaitGroup) (err error) {
	defer wg.Done()

	_, err = channel.QueueDeclare(
		name,  // name
		false, // durable
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,   // args
	)

	return
}

// ---------PRIVATE-----------------------------
func (r *Rabbit) PublishMessage(exchange, routingKey string, message []byte) error {
	return r.ProducerChannel.Publish(
		exchange,   // exchange name
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent, //(make sure if consumer die, msg isn't lost, save msg to disk)
			Timestamp:    time.Now(),
			Body:         message,
		},
	)
}

// step 4: bind exchange to queue
func (r *Rabbit) BindingQueue(exchange, queue, routingKey string, callback CallbackFunc) (err error) {
	if exchange == "" {
		return
	}

	if r.BindingRoutes == nil {
		r.BindingRoutes = make(map[string]CallbackFunc)
	}

	err = r.ConsumerChannel.QueueBind(
		queue,      // queue name
		routingKey, // routing key
		exchange,   // exchange name
		false,      // noWait
		nil,        // args
	)

	if err != nil {
		return
	}

	r.BindingRoutes[routingKey] = callback

	return
}

func (r *Rabbit) StartConsume(queue string) (data any, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	msgs, err := r.ConsumerChannel.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   //args
	)

	if err != nil {
		return nil, err
	}

	go r.ConsumeData(ctx, msgs)

	return
}

func (r *Rabbit) ConsumeData(ctx context.Context, messages <-chan amqp.Delivery) {
	for msg := range messages {
		callback := r.BindingRoutes[msg.RoutingKey]

		if callback != nil {
			callback(msg.Body)
		}
	}
}