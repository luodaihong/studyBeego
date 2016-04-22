package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/memcache"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"studyBeego/models"
	_ "studyBeego/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "test:test@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(models.User))
}

func main() {
	beego.Run()
}
