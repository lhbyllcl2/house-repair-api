package response

import (
	"net/http"
	"strings"

	"github.com/tal-tech/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

//成功返回
func Success(w http.ResponseWriter, resp interface{}) {
	var body Body
	body.Msg = "success"
	if resp == nil {
		resp = struct{}{}
	}
	body.Data = resp
	httpx.OkJson(w, body)
}
func Failed(w http.ResponseWriter, code int, err error) {
	errMsg := err.Error()
	if strings.Contains(errMsg, "rpc error") {
		index := strings.Index(errMsg, "desc = ")
		if index > 0 {
			errMsg = errMsg[index+len("desc = "):]
		}
	}
	body := Body{
		Code: code,
		Msg:  errMsg,
		Data: struct{}{},
	}
	httpx.OkJson(w, body)
}

//返回失败
func FailedWithCode(w http.ResponseWriter, code int) {
	body := Body{
		Code: code,
		Msg:  zhCNText[code],
		Data: struct{}{},
	}
	httpx.OkJson(w, body)
}
