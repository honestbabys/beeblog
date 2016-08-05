package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	IsExit := l.Input().Get("exit") == "true"
	if IsExit {
		l.Ctx.SetCookie("uname", "", -1, "/")
		l.Ctx.SetCookie("pwd", "", -1, "/")
		l.Redirect("/", 301)
		return
	}
	l.TplName = "login.html"

}

//表单提交数据接收
func (l *LoginController) Post() {
	uname := l.Input().Get("username")
	pwd := l.Input().Get("pwd")
	autologin := l.Input().Get("autologin") == "on"

	if beego.AppConfig.String("adminName") == uname &&
		beego.AppConfig.String("adminPass") == pwd {
		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1
		}
		l.Ctx.SetCookie("uname", uname, maxAge, "/")
		l.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	l.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	return uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass")
}
