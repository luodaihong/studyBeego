package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"studyBeego/models"
)

type JsonUserController struct {
	beego.Controller
	orm orm.Ormer
}

//func (this *JsonUserController) Init(ct *context.Context, controllerName, actionName string, app interface{}) {
//	fmt.Println("JsonUserController init...")
//	beego.Controller(this).Init(ct, controllerName, actionName, app)
//}

func (this *JsonUserController) Prepare() {
	this.EnableXSRF = false
	this.orm = orm.NewOrm()
	this.orm.Using("default")
}

func (this *JsonUserController) Get() {
	id, _ := this.GetInt("id", 0)
	u := models.User{Id: id}

	err := this.orm.Read(&u)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": "error", "error": err.Error()}
	} else {
		this.Data["json"] = map[string]interface{}{"status": "success", "user": u}
	}
	this.ServeJSON()
}

func (this *JsonUserController) Post() {
	u := models.User{}
	err := this.ParseForm(&u)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": "error", "ParseForm Error": err.Error()}
	} else {
		u.Id, _ = this.GetInt("id", 0)
		valid := validation.Validation{}
		b, _ := valid.Valid(&u)
		if b {
			_, dbError := this.orm.Update(&u)
			if dbError == nil {
				this.Data["json"] = map[string]interface{}{"status": "success", "user": u}
			} else {
				this.Data["json"] = map[string]interface{}{"status": "error", "DB Error": dbError.Error()}
			}
		} else {
			this.Data["json"] = map[string]interface{}{"status": "error", "Valid Error": valid.Errors}
		}
	}
	this.ServeJSON()
}

func (this *JsonUserController) Put() {
	u := models.User{}
	err := this.ParseForm(&u)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": "error", "ParseForm Error": err}
	} else {
		valid := validation.Validation{}
		b, _ := valid.Valid(&u)
		if b {
			_, dbError := this.orm.Insert(&u)
			if dbError == nil {
				this.Data["json"] = map[string]interface{}{"status": "success", "user": u}
			} else {
				this.Data["json"] = map[string]interface{}{"status": "error", "DB Error": dbError}
			}

		} else {
			this.Data["json"] = map[string]interface{}{"status": "error", "Valid Error": valid.Errors}
		}

		//麻烦的写法
		//		valid.MaxSize(u.Name, 32, "name")
		//		valid.Range(u.Age, 18, 100, "age")
		//		valid.Email(u.Email, "email")
		//		valid.MaxSize(u.Email, 32, "email")
		//		if valid.HasErrors() {
		//			this.Data["json"] = map[string]interface{}{"status": "error", "orm info": valid.Errors}
		//		} else {
		//			this.Data["json"] = map[string]interface{}{"status": "success", "user": u}
		//		}
	}
	this.ServeJSON()
}

func (this *JsonUserController) Delete() {
	u := models.User{}
	u.Id, _ = this.GetInt("id", 0)
	_, dbError := this.orm.Delete(&u)
	if dbError != nil {
		this.Data["json"] = map[string]interface{}{"status": "error", "info": dbError.Error()}
	} else {
		this.Data["json"] = map[string]interface{}{"status": "success", "info": "deleted"}
	}
	this.ServeJSON()
}
