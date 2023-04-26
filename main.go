package main

import (
	"log"
	_ "loveHome/controllers"
	_ "loveHome/routers"
	"net/http"
	"strings"

	"github.com/astaxie/beego/context"

	_ "loveHome/models"

	"github.com/astaxie/beego"
)

func ignoreStaticPath() {
	//过滤器
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	log.Println(orpath)
	//如果url中有api，则取消静态路由重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+orpath)
}

func main() {
	ignoreStaticPath()
	//beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run(":8899")
}
