package dtos

import (
	"time"

	"github.com/yagoluiz/user-api/internal/domain"
)

type UserDto struct {
	UserID    string    `json:"userId"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Priority  int       `json:"priority"`
	CreatedAt time.Time `json:"createdAt"`
}

func DomainToUsersDto(users []*domain.User) []*UserDto {
	var usersDtos []*UserDto
	for _, v := range users {
		userDto := UserDto{
			UserID:    v.UserID,
			Name:      v.Name,
			Username:  v.Username,
			Priority:  v.Priority,
			CreatedAt: v.CreatedAt,
		}

		usersDtos = append(usersDtos, &userDto)
	}

	return usersDtos
}
