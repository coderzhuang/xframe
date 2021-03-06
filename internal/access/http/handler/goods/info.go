package goods

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"time"
	"xframe/pkg/common"
	"xframe/pkg/util"
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
// @Response     200  {object}  common.Response{data=InfoRes}
// @Router       /goods [get]
func (h *HandlerGoods) Info(c *gin.Context) {
	var req InfoReq
	if err := c.ShouldBindQuery(&req); err != nil {
		common.ResponseErr(c, 100000, err.Error())
		return
	}
	data, err := h.ServiceGoods.Info(c.Request.Context(), req.Id)
	if err != nil {
		common.ResponseErr(c, 100005, err.Error())
		return
	}
	util.Get(c.Request.Context()) // 测试trace
	result := InfoRes{}
	if err := copier.Copy(&result, data); err != nil {
		common.ResponseErr(c, 100006, err.Error())
		return
	}
	common.ResponseSuc(c, result)
}
