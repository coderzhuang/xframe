package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
	"xframe/config"
)

var App = &cli.App{
	Version: fmt.Sprintf("%s|%s|%s|%s",
		runtime.GOOS, runtime.GOARCH, config.BuildVersion, config.BuildAt),
	Action: Run,
}

func Run(c *cli.Context) error {
	time.Local, _ = time.LoadLocation("Asia/Shanghai")

	cronjob := InitCron()
	cronjob.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx := cronjob.Stop()
	log.Println("Shutting down cron...")
	select {
	case <-time.After(10 * time.Second):
		log.Fatal("Cron forced to shutdown...")
	case <-ctx.Done():
		log.Println("Cron exiting...")
	}
	return nil
}

func main() {
	_ = App.Run(os.Args)
}
