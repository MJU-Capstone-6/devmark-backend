package customerror

func (e *CustomError) Error() string {
	return e.Msg
}

func NewError(statusCode int, msg string, errorCode string) *CustomError {
	return &CustomError{
		Msg:        msg,
		StatusCode: statusCode,
		ErrorCode:  errorCode,
	}
}

func GetError(e *CustomError, data interface{}) *CustomError {
	return &CustomError{
		Msg:        e.Msg,
		StatusCode: e.StatusCode,
		ErrorCode:  e.ErrorCode,
		Data:       e.Data,
	}
}
