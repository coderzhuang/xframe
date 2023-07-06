package user

import (
	"context"
	"errors"
	"fmt"
	"git.juqitech.com/go/hms-components/crypto"
	"git.juqitech.com/go/hms-components/snowflake"
	"regexp"
	"xframe/internal/repository/user"
	"xframe/internal/service/user/entity"
)

type User struct {
	repoUser user.IUserRepository
}

func New(repoUser user.IUserRepository) *User {
	return &User{repoUser: repoUser}
}

func (s *User) checkPasswordLever(pwd string) error {
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

func (s *User) createUser(ctx context.Context, data *entity.User) (err error) {
	userId, _ := snowflake.NewIdString()
	pwd, err := crypto.MD5(userId + data.Pwd)
	if err != nil {
		return
	}
	data.UserID = userId
	data.Pwd = pwd
	data.NickName = data.UserName
	err = s.repoUser.CreateUSer(ctx, data)
	return
}

func (s *User) Register(ctx context.Context, data *entity.User) (err error) {
	if err = s.checkPasswordLever(data.Pwd); err != nil {
		err = errors.New("密码长度8~16位，由数字+字母组成,不能包含其他字符哦")
		return
	}

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
