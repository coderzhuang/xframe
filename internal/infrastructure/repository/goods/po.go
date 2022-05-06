package goods

import (
	"time"
)

type Goods struct {
	ID        uint64
	Name      string
	No        string `copier:"GoodsNo"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Goods) TableName() string {
	return "tab_goods"
}
