package proxy

import (
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

type proxy struct {
	address string
	path    string
	ctx     *gin.Context
}

func NewProxy(ctx *gin.Context, path, address string) *proxy {
	return &proxy{address: address, path: path, ctx: ctx}
}

func (p *proxy) Proxy() error {
	r, err := url.Parse(p.address)
	if err != nil {
		return err
	}
	p.ctx.Request.URL.Path = p.path
	httputil.NewSingleHostReverseProxy(r).ServeHTTP(p.ctx.Writer, p.ctx.Request)
	return nil
}
