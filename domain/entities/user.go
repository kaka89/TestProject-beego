package entities

import (
	//"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Name     string
	Password string
	Age      int16
}

func init() {
	// 需要在init中注册定义的model
	//orm.RegisterModel(new(User))
}
