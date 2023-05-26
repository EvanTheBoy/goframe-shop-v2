package backend

import "time"

type LoginDoReq struct {
	//g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}

// LoginDoRes for jwt
type LoginDoRes struct {
	Info   interface{} `json:"info"`
	Token  string      `json:"token"`
	Expire time.Time   `json:"expire"`
}
