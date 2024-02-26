package config

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
)

// All config
type Config struct {
	HttpServer HttpServer
	Cors       cors.Config
	GRPC       GrpcClient
	DBOrm      DBOrm
}

type HttpServer struct {
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

type DBOrm struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	DBName   string `mapstructure:"DB_NAME"`
	DBUser   string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PWD"`
}

type GrpcClient struct {
	UserServiceHost string `mapstructure:"GRPC_USER_HOST"`
	UserServicePort string `mapstructure:"GRPC_USER_PORT"`

	ChatServiceHost string `mapstructure:"GRPC_CHAT_HOST"`
	ChatServicePort string `mapstructure:"GRPC_CHAT_PORT"`
}

type AWSConfig struct {
	Region string `mapstructure:"AWS_REGION"`
	KMSKey string `mapstructure:"AWS_KMS_KEY"`
}

var AWS = &AWSConfig{}

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
	err = ParseENV(&config.DBOrm)
	if err != nil {
		fmt.Printf("Error parsing database env. Error Detail %s", err.Error())
		return
	}
	err = ParseENV(&config.HttpServer)
	if err != nil {
		fmt.Printf("Error parsing grpc http. Error Detail %s", err.Error())
		return
	}
	config.Cors = GetCorsConfig()

	return
}

func ParseENV[T interface{}](object T) (error) {
	err := viper.Unmarshal(object)
	if err != nil {
		return err
	}
	return nil
}

func GetCorsConfig() cors.Config {
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	configCors.AllowHeaders = []string{
		"Origin",
		"Authorization",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"X-User-Agent",
	}
	configCors.ExposeHeaders = []string{
		"Origin",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"X-User-Agent",
	}
	configCors.AllowCredentials = true

	return configCors
}
