package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "studyBeego/models"
)

type AuthController struct {
	beego.Controller
	orm orm.Ormer
}

func (this *AuthController) Prepare() {
	this.EnableXSRF = false
	this.orm = orm.NewOrm()
	this.orm.Using("default")

}

func (this *AuthController) Login() {
	this.TplName = "login.html"

	if this.Ctx.Request.Method == "POST" {
		name := this.GetString("name", "")
		password := this.GetString("password", "")
		if name != "" && name == password {
			exist := this.orm.QueryTable("user").Filter("name", name).Exist()
			if !exist {
				this.Data["error"] = "user not exist."
			} else {
				this.SetSession("_username_logined", name)
				this.Redirect("/htmluser/list", 302)
			}
		} else {
			this.Data["error"] = "name must equals password, and not null"
		}
	}
}

func (this *AuthController) Logout() {
	this.DelSession("_username_logined")
	this.Redirect(this.URLFor("AuthController.Login"), 302)
}
