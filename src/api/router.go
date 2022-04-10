package api

import (
    "github.com/gogf/gf/v2/net/ghttp"
    "github.com/gogf/gf/v2/frame/g"
)

func InitRouters() {
	s := g.Server()
    s.BindHandler("/:name", func(r *ghttp.Request){
       r.Response.Writeln(r.Router.Uri)
    })
	s.SetPort(8998)
    s.Run()
}