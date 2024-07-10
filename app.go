package sdk

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk/config"
	"google.golang.org/grpc"
)

type App struct {
	Router         *gin.Engine
	Config         *config.Config
	Handler        map[string]interface{}
	CronJobList    []*CronJob
	GRPCServerList []*GRPCServer
}

// TODO: should add amqp, grpc connection, databases connection into App

// Create App.Start() func that run the app

// Run all cronjobs
func (app *App) Start() error {
	var wg = sync.WaitGroup{}

	// start servers
	if len(app.GRPCServerList) > 0 {
		for _, s := range app.GRPCServerList {
			wg.Add(1)
			go s.Start(&wg)
		}
		fmt.Println("[ 🚀 ] GRPC Server started.")
	}

	if len(app.CronJobList) > 0 {
		for _, cr := range app.CronJobList {
			wg.Add(1)
			go cr.Execute()
		}
		fmt.Println("[ 🚀 ] Cronjobs started.")
	}
	wg.Wait()

	return nil
}
