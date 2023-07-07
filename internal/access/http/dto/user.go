package dto

type RegisterReq struct {
	Username string `json:"username" validate:"required"` //
	Password string `json:"password" validate:"required"` //
}

type GetUserReq struct {
	UserId int64 `form:"user_id" validate:"required"` //
}

type GetUserResp struct {
	UserID    string `json:"user_id"`    // 用户ID
	NickName  string `json:"nick_name"`  // 用户昵称
	Sex       int    `json:"sex"`        // 会员性别 0男 1女
	AvatarURL string `json:"avatar_url"` // 用户头像
	Mobile    string `json:"mobile"`     // 手机号
}
