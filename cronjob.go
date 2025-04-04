package sdk

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Task = func()

type CronJob struct {
	Task   Task
	delay  int
	period int
}

func (app *App) NewCronJob() *CronJob {
	var cron = &CronJob{}
	app.CronJobList = append(app.CronJobList, cron)
	return cron
}

// fn is the function, delay and period is second
func (cron *CronJob) SetCronJob(fn Task, delay, period int) {
	cron.Task = fn
	cron.delay = delay
	cron.period = period
}

func (cron *CronJob) Execute() {
	// delay
	time.Sleep(time.Duration(cron.delay) * time.Second)

	// run first time
	cron.Task()
	tick := time.NewTicker(time.Second * time.Duration(cron.period))
	go func(tick *time.Ticker) {
		for range tick.C {
			cron.Task()
		}
	}(tick)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	tick.Stop()
	os.Exit(0)
}
