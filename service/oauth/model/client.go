package model

//授信的客户端
type ClientModel struct {
	//客户id
	clientId string
	//用户授权客户之后，重定向地址
	redirectUrl string
}

//设置clientId
func (c *ClientModel) SetClientId(clientId string) {
	c.clientId = clientId
}

func (c *ClientModel) SetRedirectUrl(redirectUrl string) {
	c.redirectUrl = redirectUrl
}
