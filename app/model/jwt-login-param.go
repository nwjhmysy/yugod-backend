package model

type LoginParam struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginClaims struct {
	Id       uint   `json:"id"`
	Password string `json:"password"`
	Auth     uint   `json:"auth"`
}
