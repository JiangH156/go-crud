package dto

import "john/gin-curd/models"

type UserDto struct {
	Name      string
	Telephone string
}

func ToUserDto(user models.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
