package goods

import "xframe/service/goods"

type HandlerGoods struct {
	ServiceGoods *goods.Goods
}

func New(s *goods.Goods) *HandlerGoods {
	return &HandlerGoods{ServiceGoods: s}
}
