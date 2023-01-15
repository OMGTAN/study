package models

import (
	"blog/common"
	"encoding/json"
	"net/http"
)

var Res Response

type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
}

func (*Response) Success(w http.ResponseWriter, data interface{}) {

	res := new(Response)
	res.Code = 200
	res.Data = data

	result, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(result)
	common.PrintErr(err)
}

func (*Response) Fail(w http.ResponseWriter, e error) {

	res := new(Response)
	res.Code = 999
	res.Error = e.Error()

	result, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(result)
	common.PrintErr(err)
}
