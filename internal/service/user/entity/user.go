package entity

import "time"

type User struct {
	UserID        int64     `json:"user_id"`         // 用户ID
	UserName      string    `json:"user_name"`       // 用户真实姓名
	NickName      string    `json:"nick_name"`       // 用户昵称
	Sex           int       `json:"sex"`             // 会员性别 0男 1女
	AvatarURL     string    `json:"avatar_url"`      // 用户头像
	Mobile        string    `json:"mobile"`          // 手机号
	Pwd           string    `json:"pwd"`             // 密码
	LastLoginTime time.Time `json:"last_login_time"` // 最后一次登录时间
}
