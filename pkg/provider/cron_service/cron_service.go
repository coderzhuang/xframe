package cron_service

import (
	"github.com/robfig/cron/v3"
	"xframe/pkg/application"
)

type CronService struct {
	c *cron.Cron
}

func New(c *cron.Cron) application.Service {
	return &CronService{c: c}
}

func (s *CronService) Run() {
	s.c.Start()
}

func (s *CronService) Shutdown() {
	s.c.Stop()
}
