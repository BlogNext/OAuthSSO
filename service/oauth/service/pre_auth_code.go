package service

//预授权码
type PreAuthCode struct {
	//客户id
	clientId string
	//用户授权客户之后，重定向地址
	redirectUrl string
	//预授权码
	preAuthCode string
}

func NewPreAuthCode(clientId, redirectUrl string) *PreAuthCode {
	//这样写，new是分配在堆上的，不会逃逸
	preAuthCode := new(PreAuthCode)
	preAuthCode.clientId = clientId
	preAuthCode.redirectUrl = redirectUrl
	//只拷贝，只会拷贝地址作为返回值
	return preAuthCode
}

//设置预授权码
func (pac *PreAuthCode) SetPreAuthCode(preAuthCode string) {
	pac.preAuthCode = preAuthCode
}

//获取预授权码
func (pac *PreAuthCode) GetPreAuthCode() string {
	return pac.preAuthCode
}
