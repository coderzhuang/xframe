package server

import (
	"context"
	"github.com/jinzhu/copier"
	"log"
	grpcMall "xframe/internal/access/grpc/proto/mall"
	"xframe/internal/service/goods"
	"xframe/internal/service/goods/entity"
)

type Mall struct {
	grpcMall.UnimplementedMallServer
	ServiceGoods *goods.Goods
}

func New(s *goods.Goods) *Mall {
	return &Mall{ServiceGoods: s}
}

func (s *Mall) AddGoods(ctx context.Context, in *grpcMall.AddGoodsRequest) (res *grpcMall.AddGoodsReply, err error) {
	res = &grpcMall.AddGoodsReply{}
	goodsDo := entity.Goods{}
	if err = copier.Copy(&goodsDo, in); err != nil {
		res.Code = 100001
		res.Msg = err.Error()
		return
	}
	if err = s.ServiceGoods.Add(ctx, goodsDo); err != nil {
		res.Code = 100003
		res.Msg = err.Error()
		return
	}
	res.Data = true
	return
}

func (s *Mall) GetGoods(ctx context.Context, in *grpcMall.GetGoodsRequest) (res *grpcMall.GetGoodsReply, err error) {
	res = &grpcMall.GetGoodsReply{}
	if in.Id <= 0 {
		res.Code = 100004
		res.Msg = "id 不规范"
		return
	}
	var data *entity.Goods
	data, err = s.ServiceGoods.Info(ctx, int(in.Id))
	if err != nil {
		res.Code = 100005
		res.Msg = err.Error()
		return
	}

	goodsDto := grpcMall.Goods{}
	if err = copier.Copy(&goodsDto, data); err != nil {
		log.Printf("could not greet: %s", err.Error())
		res.Code = 100006
		res.Msg = err.Error()
		return
	}
	res.Data = &goodsDto
	return
}
