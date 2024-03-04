package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	sdk "github.com/hadanhtuan/go-sdk"
)

type AWSClient struct {
	AwsCfg awsConfig.Config
	AwsEnv sdk.AWSEnv
}

// TODO: Global variable for internal package, cannot export to outside
var (
	aws *AWSClient
)

func ConnectAWS() *AWSClient {
	if aws != nil {
		return aws
	}
	aws = new(AWSClient)

	sdk.ParseENV(&aws.AwsEnv)

	awsCfg, _ := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(aws.AwsEnv.Region))

	aws.AwsCfg = awsCfg
	return aws
}
