package goods

import (
	"github.com/coderzhuang/core"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"xframe/internal/consts"
	"xframe/internal/service/goods/entity"
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
		core.ResponseErr(c, consts.SERVER_ERROR, err.Error())
		return
	}
	goodsDo := entity.Goods{}
	if err := copier.Copy(&goodsDo, &req); err != nil {
		core.ResponseErr(c, consts.SERVER_ERROR, err.Error())
		return
	}
	if err := h.ServiceGoods.Add(c.Request.Context(), goodsDo); err != nil {
		core.ResponseErr(c, consts.SERVER_ERROR, err.Error())
		return
	}
	core.ResponseSuc(c, true)
}
