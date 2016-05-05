package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.Data["error"] = "Page not found. You may jump to list page by click bellowing link"
	this.TplName = "user/error.html"
}

func (this *ErrorController) Error500() {
	//this.Abort("500") cause here called
	this.Data["error"] = "Server may have bug. You may jump to list page by click bellowing link"
	this.TplName = "user/error.html"
}

func (this *ErrorController) ErrorDB() {
	//this.Abort("DB") cause here called
	this.Data["error"] = "Db may have bug. You may jump to list page by click bellowing link"
	this.TplName = "user/error.html"
}
