package webgo

import (
	ctx "context"
	"encoding/json"
	"net/http"

	"github.com/dwivedisshyam/webgo/pkg/log"
)

type Context struct {
	ctx.Context

	*WebGo

	Logger log.Logger
	resp   http.ResponseWriter
	req    *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request, webgo *WebGo) *Context {
	c := &Context{
		WebGo:  webgo,
		Logger: webgo.Logger,
		resp:   w,
		req:    r,
	}

	return c
}

func (c *Context) reset(w http.ResponseWriter, r *http.Request) {
	c.req = r
	c.resp = w
	c.Context = nil
}

func (c *Context) Bind(v interface{}) error {
	return json.NewDecoder(c.req.Body).Decode(v)
}
