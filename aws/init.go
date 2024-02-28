package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	sdkConfig "github.com/hadanhtuan/go-sdk/config"
)

type AWSEnv struct {
	Region string `mapstructure:"AWS_REGION"`
	KMSKey string `mapstructure:"AWS_KMS_KEY"`
}

// TODO: Global variable for internal package, cannot export to outside
var (
	awsEnv AWSEnv
	awsCfg awsConfig.Config
)

func ConnectAWS() {
	sdkConfig.ParseENV(&awsEnv)

	awsCfg, _ = config.LoadDefaultConfig(context.Background(),
		config.WithRegion(awsEnv.Region))
}
