package user

import "time"

type User struct {
	ID            int       `gorm:"primaryKey;column:id" json:"id"`                // 自增ID
	UserID        int64     `gorm:"column:user_id" json:"user_id"`                 // 用户ID
	UserName      string    `gorm:"column:user_name" json:"user_name"`             // 用户真实姓名
	NickName      string    `gorm:"column:nick_name" json:"nick_name"`             // 用户昵称
	Sex           int       `gorm:"column:sex" json:"sex"`                         // 会员性别 0男 1女
	AvatarURL     string    `gorm:"column:avatar_url" json:"avatar_url"`           // 用户头像
	Mobile        string    `gorm:"column:mobile" json:"mobile"`                   // 手机号
	Pwd           string    `gorm:"column:pwd" json:"pwd"`                         // 密码
	LastLoginTime time.Time `gorm:"column:last_login_time" json:"last_login_time"` // 最后一次登录时间
	IsDeleted     int       `gorm:"column:is_deleted" json:"is_deleted"`           // 逻辑删除
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`         // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`         // 更新时间
}

func (User) TableName() string {
	return "tb_user"
}
