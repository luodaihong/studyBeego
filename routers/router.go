package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
	"studyBeego/controllers"
)

var LoginFilter = func(this *context.Context) {
	loginUrl := beego.URLFor("controllers.AuthController.Login")
	if strings.HasPrefix(this.Input.URL(), loginUrl) || strings.Contains(this.Input.URL(), "/json/user") || strings.Contains(this.Input.URL(), "/regex") {
		return
	}

	_, logined := this.Input.Session("_username_logined").(string)
	if !logined {
		this.Redirect(302, loginUrl)
	}
}

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(LoginFilter),
		beego.NSRouter("/json/user", &controllers.JsonUserController{}),
		beego.NSAutoRouter(&controllers.HtmlUserController{}),
		beego.NSAutoRouter(&controllers.AuthController{}),
		beego.NSRouter("/regex/:username([\\w]+)", &controllers.RegexTestController{}, "*:Get"),
	)
	beego.AddNamespace(ns)
}
