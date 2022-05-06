package goods

import (
	"context"
	"fmt"
	"testing"
	"xframe/internal/infrastructure/repository/goods"

	"xframe/pkg"
)

func TestInfo(t *testing.T) {
	rdb := pkg.NewRedis()
	s := New(rdb, &goods.MockInfo{})
	res, _ := s.Info(context.Background(), 1)
	fmt.Println(res)
}
