package entity

type ClientEntity struct {
	//客户id
	ClientId string `form:"client_id" json:"client_id" binding:"required"`
	//用户授权客户之后，重定向地址
	RedirectUrl string `form:"redirect_url" json:"redirect_url"`
}
