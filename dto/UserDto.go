package dto

import "GoShortLink/model"

// UserDto Dto
type UserDto struct {
	Name      string `json:"name"`
	TelePhone string `json:"telephone"`
}

func ToUserDto(user *model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		TelePhone: user.Telephone,
	}
}
