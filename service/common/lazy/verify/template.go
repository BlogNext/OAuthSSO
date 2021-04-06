package verify

import "errors"

//资源
type Resource interface{}

//资源处理
//返回资源
type ResourceHandler func() Resource

//模板
type Template struct {
	verify Verify
}

func NewTemplate(url string) *Template {
	return &Template{
		verify: &VerifyAccessToken{
			url: url,
		},
	}
}

//获取资源
//返回资源,和错误信息
func (t *Template) GetResource(token string, handler ResourceHandler) (resource interface{}, err error) {

	defer func() {
		if exception := recover(); exception != nil {
			resource = nil
			err = exception.(error)
		}
	}()

	if t.verify.Verify(token) == false {
		return nil, errors.New("验证token失败")
	}

	//回调资源
	return handler(), nil
}

//获取验证
func(t *Template) GetVerify() Verify{
	return t.verify
}
