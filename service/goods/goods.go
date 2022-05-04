package goods

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"test/consts"
	"test/infrastructure/repository"
	"test/service/goods/entity"
	"time"
)

type Goods struct {
	Rdb       *redis.Client
	repoGoods repository.IGoodsRepository
}

func New(rdb *redis.Client, repoGoods repository.IGoodsRepository) *Goods {
	return &Goods{Rdb: rdb, repoGoods: repoGoods}
}

func (s *Goods) Add(ctx context.Context, data entity.Goods) error {
	var err error
	ctx, span := otel.GetTracerProvider().Tracer(consts.Name).
		Start(ctx, "Service-Add")
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()
	err = s.repoGoods.Add(ctx, data)
	return err
}

func (s *Goods) Info(ctx context.Context, id int) (res *entity.Goods, err error) {
	ctx, span := otel.GetTracerProvider().Tracer(consts.Name).
		Start(ctx, "Service-Info")
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	rKey := fmt.Sprintf("goods:%d", id)
	var cache string
	cache, err = s.Rdb.Get(ctx, rKey).Result()
	if err != nil && err != redis.Nil {
		return
	}
	res = &entity.Goods{}
	if err != nil { // 不存在缓存
		var data *entity.Goods
		data, err = s.repoGoods.Info(ctx, id)
		if err != nil {
			return
		}
		if err = copier.Copy(res, &data); err != nil {
			return
		}
		var b []byte
		b, err = json.Marshal(res)
		if err != nil {
			return
		}
		err = s.Rdb.Set(ctx, rKey, b, 600*time.Second).Err()
		if err != nil {
			return
		}
	} else {
		err = json.Unmarshal([]byte(cache), res)
		if err != nil {
			return
		}
	}
	return
}
