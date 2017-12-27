package builder

import (
	"TestProject/domain/entities"
	"TestProject/application/dto"
)

type UserBuilder struct{}

/**
 * 将entity转换成低头
 */
func (ub *UserBuilder) FromEntity(entity *entities.User) (dto *dto.UserDto, err error) {
	dto.Age = entity.Age
	dto.Id = entity.Id
	//dto.Password = entity.Password //password 不应输出
	dto.Username = entity.Password
	return
}

/**
 * 将dto 转换成 entity
 */
func (ub *UserBuilder) FromDto(dto dto.UserDto) (entity entities.User, err error) {
	entity.Age = dto.Age
	entity.Id = dto.Id
	entity.Password = dto.Password
	entity.Password = dto.Username
	return
}
