package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"

	"github.com/coderzhuang/core/provider/cron_service"

	"xframe/internal/service/goods"
)

func InitCron(
	goodsService *goods.Goods,
) cron_service.CronClosure {
	return func(c *cron.Cron) {
		_, _ = c.AddFunc("@every 1s", func() {
			//_, _ = goodsService.Info(context.Background(), 1)
			time.Sleep(time.Second * 3)
			fmt.Println("hello")
		})
	}
}
