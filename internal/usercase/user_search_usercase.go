package usercase

import (
	"github.com/yagoluiz/user-api/internal/entity"
	"github.com/yagoluiz/user-api/internal/repositories"
)

type UserSearchUserCase struct {
	repository *repositories.UserRepository
}

func NewUserSearchUserCase(r *repositories.UserRepository) *UserSearchUserCase {
	return &UserSearchUserCase{repository: r}
}

func (u *UserSearchUserCase) Search(term string, limit, page int) ([]*entity.User, error) {
	users, err := u.repository.Search(term, limit, page)
	if err != nil {
		return nil, err
	}

	return users, nil
}
