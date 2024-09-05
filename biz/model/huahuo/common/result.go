package result

import "github.com/cloudwego/hertz/pkg/protocol/consts"

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func New() *Result {
	return &Result{}
}
func (r *Result) Success(data interface{}) *Result {
	r.Code = consts.StatusOK
	r.Msg = "Success"
	r.Data = data
	return r
}
func (r *Result) SuccessWithMsg(msg string, data interface{}) *Result {
	r.Code = consts.StatusOK
	r.Msg = msg
	r.Data = data
	return r
}
func (r *Result) Error(code int, msg string) *Result {
	r.Code = code
	r.Msg = msg
	r.Data = nil
	return r
}
