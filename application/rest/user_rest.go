package rest

import (
	"TestProject/domain/service"
	"TestProject/application/dto/builder"
	"encoding/json"
	"github.com/astaxie/beego"
	"TestProject/application/dto"
	"TestProject/domain/entities"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

var (
	userBuilder = builder.UserBuilder{}
	userService = service.UserService{}
)

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for static block"
// @Success 200 {object} dto.User
// @Failure 403 :uid is empty
// router [post]
func (u *UserController) Post() {
	var uDto dto.UserDto
	json.Unmarshal(u.Ctx.Input.RequestBody, &uDto)
	var uEntity entities.User
	var err error

	uEntity, err = userBuilder.FromDto(uDto)
	if err == nil {
		userService.AddUser(uEntity)
		u.Data["json"] = uDto
	}

	u.Data["json"] = uDto

	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// router /login [get]
//func (u *UserController) Login() {
//	username := u.GetString("username")
//	password := u.GetString("password")
//	if entities.Login(username, password) {
//		u.Data["json"] = "login success"
//	} else {
//		u.Data["json"] = "user not exist"
//	}
//	u.ServeJSON()
//}
