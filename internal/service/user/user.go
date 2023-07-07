package user

import (
	"context"
	"errors"
	"git.juqitech.com/go/hms-components/crypto"
	"git.juqitech.com/go/hms-components/snowflake"
	"strconv"
	"xframe/internal/repository/user"
	"xframe/internal/service/user/entity"
)

type User struct {
	repoUser user.IUserRepository
}

func New(repoUser user.IUserRepository) *User {
	return &User{repoUser: repoUser}
}

func (s *User) createUser(ctx context.Context, data *entity.User) (err error) {
	userId, _ := snowflake.NewId()
	pwd, err := crypto.MD5(strconv.FormatInt(userId, 10) + data.Pwd)
	if err != nil {
		return
	}
	data.UserID = userId
	data.Pwd = pwd
	data.NickName = data.UserName
	err = s.repoUser.CreateUSer(ctx, data)
	return
}

func (s *User) GetUser(ctx context.Context, userId int64) (res *entity.User, err error) {
	return s.repoUser.GetUserById(ctx, userId)
}

func (s *User) Register(ctx context.Context, data *entity.User) (err error) {
	userInfo, err := s.repoUser.GetUserByName(ctx, data.UserName)
	if err != nil {
		return
	}
	if userInfo != nil {
		err = errors.New("用户已存在")
		return
	}

	err = s.createUser(ctx, data)
	if err != nil {
		return
	}

	//token, err := s.GenerateShopUserToken(ctx, userInfo, "")
	//if err != nil {
	//	nil, gins.NewAPIError("创建用户失败", apiconstant.RESPONSE_ERROR)
	//	return
	//}
	//
	//refreshToken, err := s.GenerateShopUserRefreshToken(ctx, userInfo, "")
	//if err != nil {
	//	nil, gins.NewAPIError("创建用户失败", apiconstant.RESPONSE_ERROR)
	//	return
	//}
	return
}
