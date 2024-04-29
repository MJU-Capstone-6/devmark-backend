package customerror

type CustomError struct {
	Data       interface{} `json:"data"`
	ErrorCode  string      `json:"error_code"`
	Msg        string      `json:"msg"`
	StatusCode int         `json:"status_code"`
}
