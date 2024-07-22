package aws

import (
	"context"
	"fmt"
	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AWSClient struct {
	AwsCfg awsConfig.Config
	KMSKey string
}

var (
	AWS *AWSClient
)

func ConnectAWS(region string, KMSKey string) *AWSClient {
	AWS = &AWSClient{
		KMSKey: KMSKey,
	}

	awsCfg, _ := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region))

	AWS.AwsCfg = awsCfg

	fmt.Println("[ 🚀 ] Connected Successfully to AWS")
	return AWS
}

func GetConnection() *AWSClient {
	if AWS == nil {
		panic("Cannot connect to AWS")
	}
	return AWS
}
