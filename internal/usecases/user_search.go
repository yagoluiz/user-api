package usecases

import (
	"github.com/yagoluiz/user-api/internal/domain"
	"github.com/yagoluiz/user-api/internal/repositories"
	"github.com/yagoluiz/user-api/pkg/logger"
)

type UserSearchUseCaseInterface interface {
	FindUser(term string, limit, page int) ([]*domain.User, error)
}

type UserSearchUseCase struct {
	logger     logger.Logger
	repository repositories.UserRepositoryInterface
}

func NewUserSearchUseCase(l logger.Logger, r repositories.UserRepositoryInterface) *UserSearchUseCase {
	return &UserSearchUseCase{logger: l, repository: r}
}

func (u *UserSearchUseCase) FindUser(term string, limit, page int) ([]*domain.User, error) {
	users, err := u.repository.Search(term, limit, page)
	if err != nil {
		u.logger.Errorf("Find user error: %s", err.Error())
		return nil, err
	}

	return users, nil
}
