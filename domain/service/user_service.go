package service

import (
	"TestProject-beego/domain/entities"
	"TestProject-beego/infrastructure/repository"
)

var (
	up = repository.UserRepository{}
)

type UserService struct{}

func init() {

}

func (us *UserService) AddUser(u entities.User) int {
	user := entities.User{u.Id, "astaxie", "11111", 88}

	up.Save(user)
	return u.Id
}
