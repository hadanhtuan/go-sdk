package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hadanhtuan/go-sdk/common"
	"github.com/matelang/jwt-go-aws-kms/v2/jwtkms"
)

// TODO: don't need to provide access key, when deploy to EC2 need to associate role
func NewJWT(payload *common.JWTPayload) (string, error) {
	awsCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(aws.Region))
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	payload.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expirationTime)

	jwtToken := jwt.NewWithClaims(jwtkms.SigningMethodECDSA256, payload)

	kmsConfig := jwtkms.NewKMSConfig(kms.NewFromConfig(awsCfg), aws.KMSKey, false) // TODO: not multi region

	str, err := jwtToken.SignedString(kmsConfig.WithContext(context.Background()))

	if err != nil {
		return "", err
	}

	return str, nil

}

func VerifyJWT(token string) (*common.JWTPayload, error) {
	awsCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(aws.Region))
	if err != nil {
		return nil, err
	}

	kmsConfig := jwtkms.NewKMSConfig(kms.NewFromConfig(awsCfg), aws.KMSKey, false) 

	payload := common.JWTPayload{}

	_, err = jwt.ParseWithClaims(token, &payload, func(token *jwt.Token) (interface{}, error) {
		return kmsConfig, nil
	})
	if err != nil {
		return nil, err
	}
	return &payload, nil

}
