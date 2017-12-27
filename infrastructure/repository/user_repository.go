package repository

import (
	"github.com/astaxie/beego/orm"
	"TestProject-beego/domain/entities"

)
//todo 放在父类里面更合适


func init()  {

}

type UserRepository struct {
	//func Save(user *entities.User);

}
/**
 * 存储至数据库
 */
func (u *UserRepository) Save(user entities.User){
	//SQL
	var o = orm.NewOrm()//应该放到外面去
	o.Using("default")
	//	 ormer.Raw("UPDATE `user` SET `user_name` = ? WHERE `user_name` = ?", "slene", "testing").Exec()
	o.Insert(&user)
}
