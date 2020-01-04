package controllers

import (
	"assets-api/errors"
	"github.com/astaxie/beego"
	"math/rand"
	"time"
)

type BaseController struct {
	beego.Controller
}

func NewResponse() *Response {
	p := &Response{}
	p.Error = &errors.Success

	return p
}

type Response struct {
	Error   *errors.ErrorCode
	Result  interface{} `json:"'[]'"`
	Header  map[string]string
	TraceId uint64
}

func (c *BaseController) ResponseDone(resp *Response) {
	result := make(map[string]interface{})
	result["errno"] = resp.Error.Errno
	if resp.Error.Errno != 0 {
		result["errmsg"] = resp.Error.ErrMsg
	}
	if resp.Error.Errno == 0 {
		result["data"] = resp.Result
	}
	result["request_id"] = getTraceID()
	c.Data["json"] = result
	c.ServeJSON()
}

func getTraceID() uint64 {
	rand.Seed(time.Now().UnixNano())
	return uint64(rand.Uint32())<<12 + uint64(rand.Uint32())
}
