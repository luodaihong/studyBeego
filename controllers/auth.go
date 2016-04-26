package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"html/template"
	_ "studyBeego/models"
)

type AuthController struct {
	beego.Controller
	orm orm.Ormer
}

func (this *AuthController) Prepare() {
	this.orm = orm.NewOrm()
	this.orm.Using("default")

}

func (this *AuthController) Login() {
	session_user := this.GetSession("_username_logined")
	if session_user != nil {
		beego.Debug(session_user, " has logined.")
		this.Redirect(this.URLFor("HtmlUserController.List"), 302)
	}
	this.TplName = "login.html"
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

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
