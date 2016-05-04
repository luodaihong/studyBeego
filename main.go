package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/memcache"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"studyBeego/models"
	_ "studyBeego/routers"
)

var LoginFilter = func(this *context.Context) {
	if strings.HasPrefix(this.Input.URL(), "/auth/login") {
		return
	}

	_, logined := ctx.Input.Session("_username_logined").(string)
	if !logined {
		this.Redirect(302, "/auth/login")
	}
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "test:test@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(models.User))
	beego.InsertFilter("/*", beego.BeforeRouter, LoginFilter)
}

func main() {
	beego.Run()
}
