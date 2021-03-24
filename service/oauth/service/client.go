package service

type Client struct {
	//客户id
	clientId string
	//用户授权客户之后，重定向地址
	redirectUrl string
}

func NewClient(clientId, redirectUrl string) *Client {
	client := new(Client)
	client.clientId = clientId
	client.redirectUrl = redirectUrl
	return client
}

func (c *Client) GetClientId() string {
	return c.clientId
}

func (c *Client) GetRedirectUrl() string {
	return c.redirectUrl
}
