package models

type User struct {
	Id    int    `form:"-"`
	Name  string `form:"name" valid:"MinSize(3);MaxSize(32)";`
	Age   int    `form:"age" valid:"Range(18, 100)"`
	Email string `form:"email" valid:"Email; MaxSize(32)"`
}
