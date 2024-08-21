package sdk

import (
	"fmt"
	"sync"
)

type App struct {
	Name           string
	CronJobList    []*CronJob
	GRPCServerList []*GRPCServer
	HTTPServerList []*HTTPServer
}

// App func run cronjob, start grpc server, http server
func (app *App) Start() error {
	var wg = sync.WaitGroup{}

	// start GRPC servers
	if len(app.GRPCServerList) > 0 {
		for _, s := range app.GRPCServerList {
			wg.Add(1)
			go s.Start(&wg)
		}
		fmt.Println("[ 🚀 ] GRPC Server started.")
	}

	// start HTTP servers
	if len(app.HTTPServerList) > 0 {
		for _, s := range app.HTTPServerList {
			wg.Add(1)
			go s.Start(&wg)
		}
		fmt.Println("[ 🚀 ] HTTP Server started.")
	}

	// start cronjob
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
