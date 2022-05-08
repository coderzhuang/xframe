package cron

import (
	"github.com/robfig/cron/v3"
	"xframe/internal/service/goods"
	"xframe/pkg/provider/cron_service"
)

func InitCron(
	goodsService *goods.Goods,
) cron_service.CronClosure {
	return func(c *cron.Cron) {
		//_, _ = c.AddFunc("@every 1s", func() {
		//	res, _ := goodsService.Info(context.Background(), 1)
		//	fmt.Printf("%+v\n", res)
		//})
	}
}
