package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	sdk "github.com/hadanhtuan/go-sdk"
)

// TODO: Global variable for internal package, cannot export to outside
var (
	awsEnv sdk.AWSEnv
	awsCfg awsConfig.Config
)

func ConnectAWS() {
	sdk.ParseENV(&awsEnv)

	awsCfg, _ = config.LoadDefaultConfig(context.Background(),
		config.WithRegion(awsEnv.Region))
}
