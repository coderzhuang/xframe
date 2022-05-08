package application

import (
	"go.uber.org/dig"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Service interface {
	Run()
	Shutdown()
}

type ServerGroup struct {
	dig.In

	Services []Service `group:"server"`
}

type Application struct {
	Services []Service
}

func New(sg ServerGroup) *Application {
	return &Application{Services: sg.Services}
}

func (a *Application) Start() {
	for _, service := range a.Services {
		go service.Run()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	for _, service := range a.Services {
		service.Shutdown()
	}
	log.Println("Services Shut down")
}
