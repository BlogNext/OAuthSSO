package exception

//自定义异常
type MyException interface {
	error
	SetErrorCode(errorCode int)
	GetErrorCode() int
	SetErrorMsg(errMsg string)
}


const(
	ParamErr = iota + 1 //参数错误
)

type BaseException struct {
	errorMsg  string
	errorCode int
}

func NewException(errorCode int, errorMsg string) *BaseException {
	baseException := new(BaseException)
	baseException.errorCode = errorCode
	baseException.errorMsg = errorMsg
	return baseException
}

func (b *BaseException) SetErrorCode(errorCode int) {
	b.errorCode = errorCode
}

func (b *BaseException) SetErrorMsg(errorMsg string) {
	b.errorMsg = errorMsg
}

func (b *BaseException) GetErrorCode() int {
	return b.errorCode
}

func (b *BaseException) Error() string {
	return b.errorMsg
}
