package dto

type RegisterReq struct {
	Username string `json:"username" validate:"required"` //
	Password string `json:"password" validate:"required"` //
}
