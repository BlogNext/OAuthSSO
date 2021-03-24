package service

//预授权码
type PreAuthCode struct {
	//客户
	Client *Client
	//用户
	user *User
	//预授权码
	preAuthCode string
}

//创建一个预授权码
func NewPreAuthCode(preAuthCode string, client *Client, user *User) *PreAuthCode {
	PAC := new(PreAuthCode)
	PAC.preAuthCode = preAuthCode
	PAC.Client = client
	PAC.user = user
	return PAC
}

//获取预授权码
func (pac *PreAuthCode) GetPreAuthCode() string {
	return pac.preAuthCode
}
