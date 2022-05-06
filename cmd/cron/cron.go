package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func InitCron() *cron.Cron {
	cronjob := cron.New()
	_, _ = cronjob.AddFunc("@every 1s", func() {
		fmt.Println("asdasda")
		//time.Sleep(time.Second * 6)
	})
	return cronjob
}
