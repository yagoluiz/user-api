package usercase

import (
	"github.com/yagoluiz/user-api/internal/domain"
)

type UserSearchUserCase struct {
	repository domain.UserRepository
}

func NewUserSearchUserCase(r domain.UserRepository) *UserSearchUserCase {
	return &UserSearchUserCase{repository: r}
}

func (u *UserSearchUserCase) FindUser(term string, limit, page int) ([]*domain.User, error) {
	users, err := u.repository.Search(term, limit, page)
	if err != nil {
		return nil, err
	}

	return users, nil
}
