package goods

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"xframe/consts"
	"xframe/pkg"
	"xframe/service/goods/entity"
)

type AddReq struct {
	Name    string `json:"name" binding:"required"` //
	GoodsNo string `json:"goods_no"`                //
}

func (h *HandlerGoods) Add(c *gin.Context) {
	var err error
	ctx, span := otel.GetTracerProvider().Tracer(consts.Name).
		Start(c, "handler-Add")
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	var req AddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.ResponseErr(c, 100000, err.Error())
		return
	}
	span.SetAttributes(attribute.String("request.name", req.Name))
	goodsDo := entity.Goods{}
	if err := copier.Copy(&goodsDo, &req); err != nil {
		pkg.ResponseErr(c, 100001, err.Error())
		return
	}
	if err := h.ServiceGoods.Add(ctx, goodsDo); err != nil {
		pkg.ResponseErr(c, 100003, err.Error())
		return
	}
	pkg.ResponseSuc(c, true)
}