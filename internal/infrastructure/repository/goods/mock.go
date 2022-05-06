package goods

import (
	"context"
	"time"
	"xframe/internal/service/goods/entity"
)

type MockInfo struct{}

func (*MockInfo) i() {}

func (*MockInfo) Add(context.Context, entity.Goods) error {
	return nil
}

func (*MockInfo) Info(context.Context, int) (*entity.Goods, error) {
	t, _ := time.ParseInLocation("2006-01-02 03:04:05", "2000-10-01 00:00:00", time.Local)
	return &entity.Goods{
		Name:      "mock_test",
		GoodsNo:   "100000",
		CreatedAt: t,
	}, nil
}
