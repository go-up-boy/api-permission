package errorsx

//定义错误类型
const defaultCode = 1001

type CodeError struct {
	Code    int `json:"code"`
	Message string  `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Message: msg}
}

func NewErrorData(code int, msg string,data interface{}) error {
	return &CodeError{Code: code, Message: msg, Data: data}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Message
}

//------------------Response响应值----------------------

//func (e *CodeError) Info()响应返回值

type CodeErrorResponse  struct {
	Code    int `json:"code"`
	Message string  `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

// Info 返回相关参数
func (e *CodeError) Info() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
		Data:   e.Data,
	}
}