package cron_service

import (
	"github.com/robfig/cron/v3"
)

type CronClosure func(*cron.Cron)

func NewRouter(fn CronClosure) *cron.Cron {
	c := cron.New()
	fn(c)
	return c
}
