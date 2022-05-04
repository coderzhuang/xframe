package goods

import (
	"context"
	"fmt"
	"test/config"
	"test/pkg"
	"test/service/goods/entity"
	"testing"
	"time"
)

type MockInfo struct{}

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

func TestInfo(t *testing.T) {
	config.Init()
	rdb := pkg.NewRedis()
	s := New(rdb, &MockInfo{})
	res, _ := s.Info(context.Background(), 1)
	fmt.Println(res)
}
