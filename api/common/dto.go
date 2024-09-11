package common

type LoginDto struct {
	Username string `form:"username" json:"username"` // 账号
	Password string `form:"password" json:"password"` // 密码
	Code     string `form:"code" json:"code"`         // 验证码
}

type RegisterDto struct {
	Username string `form:"username" json:"username"` // 账号
	Password string `form:"password" json:"password"` // 密码
	Code     string `form:"code" json:"code"`         // 验证码
}
