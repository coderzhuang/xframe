package goods

import "test/service/goods"

type HandlerGoods struct {
	ServiceGoods *goods.Goods
}

func New(s *goods.Goods) *HandlerGoods {
	return &HandlerGoods{ServiceGoods: s}
}
