package goods

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"time"
	"xframe/internal/consts"
	"xframe/pkg/common"
)

type InfoReq struct {
	Id int `form:"id" validate:"required,gt=0"` // 商品ID
}

type InfoRes struct {
	Name       string    `json:"name"`                         // 商品名称
	GoodsNo    string    `json:"goods_no"`                     // 商品NO
	CreateTime time.Time `json:"create_at" copier:"CreatedAt"` // 商品创建时间
}

// Info
// @Summary      goods detail
// @Description
// @Produce      json
// @Param        id query string true "商品ID"
// @Response     200  {object}  pkg.Response{data=InfoRes}
// @Router       /goods [get]
func (h *HandlerGoods) Info(c *gin.Context) {
	var err error
	ctx, span := otel.GetTracerProvider().Tracer(consts.Name).
		Start(c, "handler-Info")
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	var req InfoReq
	if err := c.ShouldBindQuery(&req); err != nil {
		common.ResponseErr(c, 100000, err.Error())
		return
	}
	span.SetAttributes(attribute.Int("request.id", req.Id))
	data, err := h.ServiceGoods.Info(ctx, req.Id)
	if err != nil {
		common.ResponseErr(c, 100005, err.Error())
		return
	}
	result := InfoRes{}
	if err := copier.Copy(&result, data); err != nil {
		common.ResponseErr(c, 100006, err.Error())
		return
	}
	common.ResponseSuc(c, result)
}
