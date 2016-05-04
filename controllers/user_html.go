package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"html/template"
	"studyBeego/models"
)

type HtmlUserController struct {
	beego.Controller
	orm orm.Ormer
}

func (this *HtmlUserController) Prepare() {
	this.orm = orm.NewOrm()
	this.orm.Using("default")
	this.Data["loginUser"] = this.GetSession("_username_logined")
}

func (this *HtmlUserController) Detail() {
	id, _ := this.GetInt("id", 0)
	u := models.User{Id: id}
	err := this.orm.Read(&u)

	if err != nil {
		this.Data["error"] = err
		this.TplName = "user/error.html"
	} else {
		this.Data["user"] = u
		this.TplName = "user/detail.html"
	}
}

func (this *HtmlUserController) List() {
	var users []*models.User
	querySeter := this.orm.QueryTable("user")
	querySeter.All(&users)
	this.Data["users"] = users
	this.TplName = "user/list.html"
}

func (this *HtmlUserController) Add() {
	this.TplName = "user/add.html"
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	if this.Ctx.Request.Method == "POST" {
		u := models.User{}
		err := this.ParseForm(&u)
		if err != nil {
			this.Data["error"] = err
			return
		}

		valid := validation.Validation{}
		b, _ := valid.Valid(&u)
		if !b {
			var buffer bytes.Buffer
			for _, one := range valid.Errors {
				buffer.WriteString(one.Field + " " + one.Message + ", ")
			}
			this.Data["error"] = buffer.String()
			return
		}

		_, dbError := this.orm.Insert(&u)
		if dbError == nil {
			this.Redirect(this.URLFor("HtmlUserController.List"), 302)
		} else {
			this.Data["error"] = dbError.Error()
		}
	}
}

func (this *HtmlUserController) Edit() {
	this.TplName = "user/edit.html"
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	id, _ := this.GetInt("id", 0)
	if id <= 0 {
		this.Data["error"] = "bad uer id"
		this.TplName = "user/error.html"
		return
	}

	u := models.User{Id: id}
	err := this.orm.Read(&u)
	if err != nil {
		this.Data["error"] = err
		this.TplName = "user/error.html"
		return
	}

	if this.Ctx.Request.Method == "GET" {
		this.Data["user"] = u
		return
	}

	u.Name = ""
	u.Age = 0
	u.Email = ""
	err = this.ParseForm(&u)
	if err != nil {
		this.Data["error"] = err
		this.Data["user"] = u
		return
	}

	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		var buffer bytes.Buffer
		for _, one := range valid.Errors {
			buffer.WriteString(one.Field + " " + one.Message + ", ")
		}
		this.Data["error"] = buffer.String()
		this.Data["user"] = u
		return
	}

	_, dbError := this.orm.Update(&u)
	if dbError == nil {
		this.Redirect(this.URLFor("HtmlUserController.List"), 302)
	} else {
		this.Data["error"] = dbError.Error()
		this.Data["user"] = u
	}
}

func (this *HtmlUserController) Delete() {
	u := models.User{}
	u.Id, _ = this.GetInt("id", 0)
	this.orm.Delete(&u)
	this.Redirect(this.URLFor("HtmlUserController.List"), 302)
}
