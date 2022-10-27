package usercase

import "github.com/yagoluiz/user-api/internal/repositories"

type UserSearchUserCase struct {
	Repository *repositories.UserRepository
}

func NewUserSearchUserCase(r *repositories.UserRepository) *UserSearchUserCase {
	return &UserSearchUserCase{Repository: r}
}

func (u *UserSearchUserCase) Search() error {
	return nil
}
