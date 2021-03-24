package entity

import "github.com/go-playground/validator/v10"

//github.com/go-playground/validator/v10自定义错误消息接口
type ValidatorCustomMsg interface {
	//获取自定义的错误消息,统一定义一下
	GetError(err validator.ValidationErrors) string
}
