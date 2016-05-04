package models

import (
	"github.com/astaxie/beego/validation"
	"strings"
)

type User struct {
	Id    int    `form:"-"`
	Name  string `form:"name" valid:"MinSize(3);MaxSize(32)";`
	Age   int    `form:"age" valid:"Range(18, 100)"`
	Email string `form:"email" valid:"Email; MaxSize(32)"`
}

func (u *User) Valid(v *validation.Validation) {
	//after all tests in StructTag succeed, run this func for custom valid
	if strings.Index(u.Name, "xijinping") != -1 {
		v.SetError("Name", "Can't contain 'xijinping' in Name")
	}
}
