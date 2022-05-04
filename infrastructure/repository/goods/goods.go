package goods

import (
	"context"
	"github.com/jinzhu/copier"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"gorm.io/gorm"
	"xframe/consts"
	"xframe/infrastructure/repository"
	"xframe/service/goods/entity"
)

type Repository struct {
	Db *gorm.DB
}

func New(db *gorm.DB) repository.IGoodsRepository {
	return &Repository{Db: db}
}

func (d *Repository) Add(ctx context.Context, data entity.Goods) error {
	_, span := otel.GetTracerProvider().Tracer(consts.Name).
		Start(ctx, "Repository-Add")
	defer span.End()

	goodsPo := Goods{}
	if err := copier.Copy(&goodsPo, &data); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return d.Db.WithContext(ctx).Create(&goodsPo).Error
}

func (d *Repository) Info(ctx context.Context, id int) (res *entity.Goods, err error) {
	_, span := otel.GetTracerProvider().Tracer(consts.Name).
		Start(ctx, "Repository-Info")
	defer span.End()

	res = &entity.Goods{}
	goodsPo := Goods{}
	err = d.Db.WithContext(ctx).First(&goodsPo, id).Error
	if err = copier.Copy(res, &goodsPo); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return
	}
	return
}
