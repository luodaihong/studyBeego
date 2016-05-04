package routers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	"studyBeego/controllers"
)

func init() {
	beego.Router("/json/user", &controllers.JsonUserController{})
	beego.AutoRouter(&controllers.HtmlUserController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.Router("/regex/:username([\\w]+)", &controllers.RegexTestController{}, "*:Get")
	//	beego.Get("/", func(ctx *context.Context) {
	//		ctx.Redirect(302, "/auth/login")
	//	})
}
