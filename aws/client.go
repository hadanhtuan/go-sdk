package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	sdkConfig "github.com/hadanhtuan/go-sdk/config"
)

type AWSClient struct {
	AwsCfg awsConfig.Config
}

// TODO: Global variable for internal package, cannot export to outside
var (
	AWS *AWSClient
)

func ConnectAWS() *AWSClient {
	AWS = new(AWSClient)

	awsCfg, _ := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(sdkConfig.AppConfig.AWS.Region))

	AWS.AwsCfg = awsCfg
	return AWS
}

func GetConnection() *AWSClient {
	if AWS != nil {
		return AWS
	}
	return ConnectAWS()
}
