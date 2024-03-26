package config

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
)

// ENV of http, cors, grpc, aws, postgresql, redis, elastic, rabbitmq
type Config struct {
	HttpServer HttpEnv
	Cors       cors.Config
	GRPC       GrpcClient
	ORM        ORMEnv
	Cache      CacheEnv
	AWS        AWSEnv
	AMQP       AMQPEnv
	ES         ESEnv
	Stripe     StripeENV
}

// Global variable for using config in SDK
var (
	AppConfig *Config
)

type HttpEnv struct {
	AppPort       string `mapstructure:"PORT"`
	ENV           string `mapstructure:"ENV" json:"ENV"`
	TrustedDomain string `mapstructure:"TRUSTED_DOMAIN"`
	LogPath       string `mapstructure:"LOG_PATH"`

	ApiPath     string `mapstructure:"API_PATH"`
	SwaggerPath string `mapstructure:"SWAGGER_PATH"`

	LimitRequest            int `mapstructure:"LIMIT_REQUEST"`
	LimitRequestPerSecond   int `mapstructure:"LIMIT_REQUEST_PER_SECOND"`
	RequestTimeoutPerSecond int `mapstructure:"REQUEST_TIMEOUT_PER_SECOND"`
}

type GrpcClient struct {
	UserServiceHost string `mapstructure:"GRPC_USER_HOST"`
	UserServicePort string `mapstructure:"GRPC_USER_PORT"`

	ChatServiceHost string `mapstructure:"GRPC_CHAT_HOST"`
	ChatServicePort string `mapstructure:"GRPC_CHAT_PORT"`

	PropertyServiceHost string `mapstructure:"GRPC_PROPERTY_HOST"`
	PropertyServicePort string `mapstructure:"GRPC_PROPERTY_PORT"`

	LocationServiceHost string `mapstructure:"GRPC_LOCATION_HOST"`
	LocationServicePort string `mapstructure:"GRPC_LOCATION_PORT"`

	SearchServiceHost string `mapstructure:"GRPC_SEARCH_HOST"`
	SearchServicePort string `mapstructure:"GRPC_SEARCH_PORT"`

	PaymentServiceHost string `mapstructure:"GRPC_PAYMENT_HOST"`
	PaymentServicePort string `mapstructure:"GRPC_PAYMENT_PORT"`
}

type ORMEnv struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	DBName   string `mapstructure:"DB_NAME"`
	DBUser   string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PWD"`
}

type CacheEnv struct {
	CacheHost string `mapstructure:"CACHE_HOST"`
	CachePort string `mapstructure:"CACHE_PORT"`
	CachePass string `mapstructure:"CACHE_PWD"`
	CacheDB   int    `mapstructure:"CACHE_DB"`
}

type AWSEnv struct {
	Region string `mapstructure:"AWS_REGION"`
	KMSKey string `mapstructure:"AWS_KMS_KEY"`
}

type AMQPEnv struct {
	Host string `mapstructure:"AMQP_HOST"`
	Port string `mapstructure:"AMQP_PORT"`
	User string `mapstructure:"AMQP_USER"`
	Pass string `mapstructure:"AMQP_PWD"`
}

type ESEnv struct {
	Host     string `mapstructure:"ES_HOST"`
	Port     int    `mapstructure:"ES_PORT"`
	Username string `mapstructure:"ES_USER"`
	Password string `mapstructure:"ES_PWD"`
}

type StripeENV struct {
	PublishKey string `mapstructure:"STRIPE_PUBLISH_KEY"`
	SecretKey  string `mapstructure:"STRIPE_SECRET_KEY"`
}

func InitConfig(path string) (config *Config, err error) {
	config = new(Config)
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configPath := fmt.Sprintf("%s%s/.env", wd, path)
	viper.SetConfigFile(configPath)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = ParseENV(&config.GRPC)
	if err != nil {
		fmt.Printf("Error parsing grpc env. Error Detail %s", err.Error())
		return
	}
	err = ParseENV(&config.Cache)
	if err != nil {
		fmt.Printf("Error parsing cache env. Error Detail %s", err.Error())
		return
	}
	err = ParseENV(&config.ORM)
	if err != nil {
		fmt.Printf("Error parsing orm env. Error Detail %s", err.Error())
		return
	}
	err = ParseENV(&config.AMQP)
	if err != nil {
		fmt.Printf("Error parsing amqp env. Error Detail %s", err.Error())
		return
	}
	err = ParseENV(&config.ES)
	if err != nil {
		fmt.Printf("Error parsing elastic search http. Error Detail %s", err.Error())
		return
	}
	err = ParseENV(&config.AWS)
	if err != nil {
		fmt.Printf("Error parsing aws env. Error Detail %s", err.Error())
		return
	}
	err = ParseENV(&config.HttpServer)
	if err != nil {
		fmt.Printf("Error parsing grpc http. Error Detail %s", err.Error())
		return
	}
	err = ParseENV(&config.Stripe)
	if err != nil {
		fmt.Printf("Error parsing stripe env. Error Detail %s", err.Error())
		return
	}
	config.Cors = GetCorsConfig()

	AppConfig = config

	return
}

func ParseENV[T interface{}](object T) error {
	err := viper.Unmarshal(object)
	if err != nil {
		return err
	}
	return nil
}
