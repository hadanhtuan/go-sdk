package config

import (
	"github.com/gin-contrib/cors"
)

func GetCorsConfig() cors.Config {
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	configCors.AllowHeaders = []string{
		"Origin",
		"Authorization",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"X-User-Agent",
	}
	configCors.ExposeHeaders = []string{
		"Origin",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"X-User-Agent",
	}
	configCors.AllowCredentials = true

	return configCors
}
