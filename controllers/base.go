package controllers

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OutResponse(code int, msg string, data interface{}) Response {
	Rsp := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	return Rsp
}
