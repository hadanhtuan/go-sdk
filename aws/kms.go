package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hadanhtuan/go-sdk/common"
	"github.com/matelang/jwt-go-aws-kms/v2/jwtkms"
)

// TODO: don't need to provide access key, when deploy to EC2 need to associate role
func NewJWT(payload *common.JWTPayload) (*common.JWTToken, error) {

	if payload.RegisteredClaims.ExpiresAt == nil {
		expiresAt := time.Now().Add(3 * 24 * time.Hour)
		payload.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expiresAt)
	}

	jwtToken := jwt.NewWithClaims(jwtkms.SigningMethodECDSA256, payload)

	kmsConfig := jwtkms.NewKMSConfig(kms.NewFromConfig(awsCfg), awsEnv.KMSKey, false) // TODO: false = not multi region

	str, err := jwtToken.SignedString(kmsConfig.WithContext(context.Background()))

	if err != nil {
		return nil, err
	}

	return &common.JWTToken{
		Token:     str,
		ExpiresAt: payload.RegisteredClaims.ExpiresAt.Time,
	}, nil

}

func VerifyJWT(token string) (*common.JWTPayload, error) {

	kmsConfig := jwtkms.NewKMSConfig(kms.NewFromConfig(awsCfg), awsEnv.KMSKey, false)

	payload := common.JWTPayload{}

	_, err := jwt.ParseWithClaims(token, &payload, func(token *jwt.Token) (interface{}, error) {
		return kmsConfig, nil
	})
	if err != nil {
		return nil, err
	}
	return &payload, nil

}
