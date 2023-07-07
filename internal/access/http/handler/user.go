package handler

import (
	"fmt"
	"github.com/coderzhuang/core"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"regexp"
	"xframe/internal/access/http/dto"
	"xframe/internal/consts"
	"xframe/internal/service/user"
	"xframe/internal/service/user/entity"
	"xframe/pkg/log"
)

type HandlerUser struct {
	ServiceUser *user.User
}

func New(s *user.User) *HandlerUser {
	return &HandlerUser{ServiceUser: s}
}

func (s *HandlerUser) checkPasswordLever(pwd string) error {
	if len(pwd) < 8 || len(pwd) > 16 {
		return fmt.Errorf("密码强度必须为数字/字母/字符2中, 8~16位哦")
	}

	var num = `[0-9]{1}`
	var aZ = `[a-zA-Z]{1}`
	var specialChar = `(?m)[!@#\$%\^&\*\(\)_\+\{\}:"<>\?/\.,;'\\]`
	var matched = 0

	if b, _ := regexp.MatchString(num, pwd); b {
		matched += 1
	}

	if b, _ := regexp.MatchString(aZ, pwd); b {
		matched += 1
	}

	if b, _ := regexp.MatchString(specialChar, pwd); b {
		matched += 1
	}

	if matched < 2 {
		return fmt.Errorf("密码强度必须为数字/字母/字符2中, 8~16位哦")
	}

	return nil
}

func (h *HandlerUser) GetUser(c *gin.Context) {
	var req dto.GetUserReq
	if err := c.BindQuery(&req); err != nil {
		core.ResponseErr(c, consts.SERVER_ERROR, err.Error())
		return
	}

	userInfo, err := h.ServiceUser.GetUser(c, req.UserId)
	if err != nil {
		core.ResponseErr(c, consts.SERVER_ERROR, err.Error())
		return
	}
	res := &dto.GetUserResp{}
	if err = copier.Copy(res, userInfo); err != nil {
		core.ResponseErr(c, consts.SERVER_ERROR, err.Error())
		return
	}
	core.ResponseSuc(c, res)
}

func (h *HandlerUser) Register(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.BindJSON(&req); err != nil {
		core.ResponseErr(c, consts.SERVER_ERROR, err.Error())
		return
	}
	if err := h.checkPasswordLever(req.Password); err != nil {
		log.Errorf("Register checkPasswordLever username: %s err: %+v", req.Username, err)
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
