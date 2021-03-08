package error2

// Error 错误处理封装
type Error struct {
	Code int
}

// NewError 创建一个错误
func NewError(code int) Error {
	return Error{
		Code: code,
	}
}

func (e Error) Error() string {
	return translation(e.Code)
}
