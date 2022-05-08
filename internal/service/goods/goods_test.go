package goods

import (
	"context"
	"fmt"
	"testing"
	"xframe/internal/repository/goods"
	"xframe/pkg/redis"
)

func TestInfo(t *testing.T) {
	rdb := redis.New()
	s := New(rdb, &goods.MockInfo{})
	res, _ := s.Info(context.Background(), 1)
	fmt.Println(res)
}
