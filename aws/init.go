package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	sdkConfig "github.com/hadanhtuan/go-sdk/config"
)

type AWSEnv struct {
	Region string `mapstructure:"AWS_REGION"`
	KMSKey string `mapstructure:"AWS_KMS_KEY"`
}

// TODO: Global variable for internal package, cannot export to outside
var awsEnv AWSEnv
var awsCfg awsConfig.Config

func ConnectAWS() {
	sdkConfig.ParseENV(&awsEnv)

	awsCfg, _ = config.LoadDefaultConfig(context.Background(),
		config.WithRegion(awsEnv.Region))
}
