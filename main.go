package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/memcache"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"studyBeego/controllers"
	"studyBeego/models"
	_ "studyBeego/routers"
)

var LoginFilter = func(this *context.Context) {
	if strings.HasPrefix(this.Input.URL(), "/auth/login") {
		return
	}

	_, logined := this.Input.Session("_username_logined").(string)
	if !logined {
		this.Redirect(302, "/auth/login")
	}
}

func init() {
	//beego.BConfig.Log.AccessLogs = true
	//debug 7, info 6, warning 4, error 3
	beego.BeeLogger.SetLogger("multifile", `{"filename":"studyBeego.log","level":3,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)

	toolbox.AddHealthCheck("mysql", &controllers.DatabaseCheck{})

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "test:test@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(models.User))
	beego.InsertFilter("/*", beego.BeforeRouter, LoginFilter)
	beego.ErrorController(&controllers.ErrorController{})
}

func main() {
	beego.Run()
}
