package goods

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"xframe/internal/service/goods/entity"
	"xframe/pkg/common"
)

type AddReq struct {
	Name    string `json:"name" binding:"required"` //
	GoodsNo string `json:"goods_no"`                //
}

// Add
// @Summary      add goods
// @Description
// @Accept       json
// @Produce      json
// @Param        a body AddReq false " "
// @Response     200  {object}  common.Response
// @Router       /goods [post]
func (h *HandlerGoods) Add(c *gin.Context) {
	var req AddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ResponseErr(c, 100000, err.Error())
		return
	}
	goodsDo := entity.Goods{}
	if err := copier.Copy(&goodsDo, &req); err != nil {
		common.ResponseErr(c, 100001, err.Error())
		return
	}
	if err := h.ServiceGoods.Add(c.Request.Context(), goodsDo); err != nil {
		common.ResponseErr(c, 100003, err.Error())
		return
	}
	common.ResponseSuc(c, true)
}
