package service

import (
	"TestProject/domain/entities"
	"TestProject/infrastructure/repository"
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
