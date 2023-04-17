package usecases

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yagoluiz/user-api/internal/domain"
	"github.com/yagoluiz/user-api/internal/repositories"
	"github.com/yagoluiz/user-api/pkg/logger"
	"github.com/yagoluiz/user-api/pkg/mocks"
)

func TestFindUserUseCase(t *testing.T) {
	usersMock := mocks.GetUsersMock()
	tests := []struct {
		name       string
		logger     logger.Logger
		repository repositories.UserRepositoryInterface
		retResult  []*domain.User
		retErr     string
	}{
		{
			name: "find user when repository contains error",
			logger: func() logger.Logger {
				logger := mocks.NewLogger(t)
				logger.On("Errorf", mock.Anything, mock.Anything)
				return logger
			}(),
			repository: func() repositories.UserRepositoryInterface {
				repo := mocks.NewUserRepositoryInterface(t)
				repo.On("Search", mock.Anything, mock.Anything, mock.Anything).Once().Return(nil, errors.New(mock.Anything))
				return repo
			}(),
			retResult: nil,
			retErr:    mock.Anything,
		},
		{
			name: "find user when repository contains users",
			logger: func() logger.Logger {
				logger := mocks.NewLogger(t)
				return logger
			}(),
			repository: func() repositories.UserRepositoryInterface {
				repo := mocks.NewUserRepositoryInterface(t)
				repo.On("Search", mock.Anything, mock.Anything, mock.Anything).Once().Return(usersMock, nil)
				return repo
			}(),
			retResult: usersMock,
			retErr:    "<nil>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := UserSearchUseCase{
				logger:     test.logger,
				repository: test.repository,
			}
			exec, err := u.FindUser("anything", 1, 1)
			assert.Equal(t, test.retResult, exec, fmt.Sprintf("%v", err))
			assert.Equal(t, test.retErr, fmt.Sprintf("%v", err))
		})
	}
}
