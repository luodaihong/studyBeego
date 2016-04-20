package routers

import (
	"github.com/astaxie/beego"
	"studyBeego/controllers"
)

func init() {
	beego.Router("/json/user", &controllers.JsonUserController{})
	beego.AutoRouter(&controllers.HtmlUserController{})
}
