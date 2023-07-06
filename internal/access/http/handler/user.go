package handler

import (
	"github.com/coderzhuang/core"
	"github.com/gin-gonic/gin"
	"xframe/internal/access/http/dto"
	"xframe/internal/consts"
	"xframe/internal/service/user"
	"xframe/internal/service/user/entity"
)

type HandlerUser struct {
	ServiceUser *user.User
}

func New(s *user.User) *HandlerUser {
	return &HandlerUser{ServiceUser: s}
}

func (h *HandlerUser) Register(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.BindJSON(&req); err != nil {
		core.ResponseErr(c, consts.SERVER_ERROR, err.Error())
		return
	}

	err := h.ServiceUser.Register(c, &entity.User{
		UserName: req.Username,
		Pwd:      req.Password,
	})
	if err != nil {
		core.ResponseErr(c, consts.SERVER_ERROR, err.Error())
		return
	}
	core.ResponseSuc(c, nil)
}

func (h *HandlerUser) Login(c *gin.Context) {

}

func (h *HandlerUser) GetCaptcha(c *gin.Context) {

}
