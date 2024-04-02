package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/common"
	sdkConfig "github.com/hadanhtuan/go-sdk/config"
	"github.com/matelang/jwt-go-aws-kms/v2/jwtkms"
)

// TODO: don't need to provide access key, when deploy to EC2 need to associate role
func NewJWT(payload *common.JWTPayload) (*common.JWTToken, error) {
	AWS := GetConnection()

	if payload.RegisteredClaims.ExpiresAt == nil {
		expiresAt := time.Now().Add(3 * 24 * time.Hour)
		payload.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expiresAt)
	}

	jwtToken := jwt.NewWithClaims(jwtkms.SigningMethodECDSA256, payload)

	kmsConfig := jwtkms.NewKMSConfig(kms.NewFromConfig(AWS.AwsCfg), sdkConfig.AppConfig.AWS.KMSKey, false) // TODO: false = not multi region

	accessToken, err := jwtToken.SignedString(kmsConfig.WithContext(context.Background()))
	refreshToken := sdk.HashKey([]string{accessToken})

	if err != nil {
		return nil, err
	}

	// Access token expires: 3 day. Refresh token expires: 1 day
	return &common.JWTToken{
		AccessToken:      accessToken,
		AccessExpiresAt:  payload.RegisteredClaims.ExpiresAt.Time.Unix(),
		RefreshToken:     refreshToken,
		RefreshExpiresAt: payload.RegisteredClaims.ExpiresAt.Time.Add(1 * 24 * time.Hour).Unix(),
	}, nil

}

func VerifyJWT(token string) (*common.JWTPayload, error) {
	AWS := GetConnection()

	kmsConfig := jwtkms.NewKMSConfig(kms.NewFromConfig(AWS.AwsCfg), sdkConfig.AppConfig.AWS.KMSKey, false)

	payload := common.JWTPayload{}

	_, err := jwt.ParseWithClaims(token, &payload, func(token *jwt.Token) (interface{}, error) {
		return kmsConfig, nil
	})
	if err != nil {
		return nil, err
	}
	return &payload, nil

}
