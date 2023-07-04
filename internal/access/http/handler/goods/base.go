package goods

import "xframe/internal/service/goods"

type HandlerGoods struct {
	ServiceGoods *goods.Goods
}

// @service
func New(s *goods.Goods) *HandlerGoods {
	return &HandlerGoods{ServiceGoods: s}
}
