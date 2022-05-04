package repository

import (
	"context"
	"xframe/service/goods/entity"
)

type IGoodsRepository interface {
	Add(context.Context, entity.Goods) error
	Info(context.Context, int) (*entity.Goods, error)
}
