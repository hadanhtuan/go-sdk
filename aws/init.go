package aws

import (
	"github.com/hadanhtuan/go-sdk/config"
)

type AWSConfig struct {
	Region string `mapstructure:"AWS_REGION"`
	KMSKey string `mapstructure:"AWS_KMS_KEY"`
}

// TODO: Can only use aws ENV in local package, cannot export to outside
var aws AWSConfig  

func ConnectAWS() {
	config.ParseENV(&aws)
}
