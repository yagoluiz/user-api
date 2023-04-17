package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yagoluiz/user-api/internal/usecases"
	"github.com/yagoluiz/user-api/pkg/logger"
	"github.com/yagoluiz/user-api/pkg/mocks"
)

func TestSearchHandler(t *testing.T) {
	usersMock := mocks.GetUsersMock()
	tests := []struct {
		name        string
		logger      logger.Logger
		usecase     usecases.UserSearchUseCaseInterface
		url         string
		retHttp     int
		validateMsg bool
		retMsg      string
	}{
		{
			name: "search handler when query param is empty",
			logger: func() logger.Logger {
				logger := mocks.NewLogger(t)
				logger.On("Warnf", mock.Anything, mock.Anything)
				return logger
			}(),
			usecase: func() usecases.UserSearchUseCaseInterface {
				u := mocks.NewUserSearchUseCaseInterface(t)
				return u
			}(),
			url:         "/v1/users/search",
			retHttp:     400,
			validateMsg: true,
			retMsg: `
				{"message":"Field query required or invalid"}
			`,
		},
		{
			name: "search handler when from param is empty",
			logger: func() logger.Logger {
				logger := mocks.NewLogger(t)
				logger.On("Warnf", mock.Anything, mock.Anything)
				return logger
			}(),
			usecase: func() usecases.UserSearchUseCaseInterface {
				u := mocks.NewUserSearchUseCaseInterface(t)
				return u
			}(),
			url:         "/v1/users/search?query=anything",
			retHttp:     400,
			validateMsg: true,
			retMsg: `
				{"message":"Field from required or invalid"}
			`,
		},
		{
			name: "search handler when from param is less than 1",
			logger: func() logger.Logger {
				logger := mocks.NewLogger(t)
				logger.On("Warnf", mock.Anything, mock.Anything)
				return logger
			}(),
			usecase: func() usecases.UserSearchUseCaseInterface {
				u := mocks.NewUserSearchUseCaseInterface(t)
				return u
			}(),
			url:         "/v1/users/search?query=anything&from=0",
			retHttp:     400,
			validateMsg: true,
			retMsg: `
				{"message":"Field from required or invalid"}
			`,
		},
		{
			name: "search handler when size param is empty",
			logger: func() logger.Logger {
				logger := mocks.NewLogger(t)
				logger.On("Warnf", mock.Anything, mock.Anything)
				return logger
			}(),
			usecase: func() usecases.UserSearchUseCaseInterface {
				u := mocks.NewUserSearchUseCaseInterface(t)
				return u
			}(),
			url:         "/v1/users/search?query=anything&from=1",
			retHttp:     400,
			validateMsg: true,
			retMsg: `
				{"message":"Field size required or invalid"}
			`,
		},
		{
			name: "search handler when size param is less than 1",
			logger: func() logger.Logger {
				logger := mocks.NewLogger(t)
				logger.On("Warnf", mock.Anything, mock.Anything)
				return logger
			}(),
			usecase: func() usecases.UserSearchUseCaseInterface {
				u := mocks.NewUserSearchUseCaseInterface(t)
				return u
			}(),
			url:         "/v1/users/search?query=anything&from=1&size=0",
			retHttp:     400,
			validateMsg: true,
			retMsg: `
				{"message":"Field size required or invalid"}
			`,
		},
		{
			name: "search handler when use case return error",
			logger: func() logger.Logger {
				logger := mocks.NewLogger(t)
				logger.On("Errorf", mock.Anything, mock.Anything)
				return logger
			}(),
			usecase: func() usecases.UserSearchUseCaseInterface {
				u := mocks.NewUserSearchUseCaseInterface(t)
				u.On("FindUser", mock.Anything, mock.Anything, mock.Anything).Once().Return(nil, errors.New(mock.Anything))
				return u
			}(),
			url:         "/v1/users/search?query=anything&from=1&size=1",
			retHttp:     500,
			validateMsg: true,
			retMsg: `
				{"message":"mock.Anything"}
			`,
		},
		{
			name: "search handler when use case return users",
			logger: func() logger.Logger {
				logger := mocks.NewLogger(t)
				logger.On("Infof", mock.Anything, mock.Anything)
				return logger
			}(),
			usecase: func() usecases.UserSearchUseCaseInterface {
				u := mocks.NewUserSearchUseCaseInterface(t)
				u.On("FindUser", mock.Anything, mock.Anything, mock.Anything).Once().Return(usersMock, nil)
				return u
			}(),
			url:     "/v1/users/search?query=anything&from=1&size=1",
			retHttp: 200,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := UserSearchHandler{
				logger:             test.logger,
				userSearchUserCase: test.usecase,
			}
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, test.url, nil)
			h.Search(res, req)

			expectedCode := res.Result().StatusCode
			expectedMsg := res.Body.String()

			assert.Equal(t, test.retHttp, expectedCode, fmt.Sprintf("want: %v, got: %v", test.retHttp, expectedCode))

			if test.validateMsg {
				assert.Contains(t, test.retMsg, expectedMsg, fmt.Sprintf("want: %v, got: %s", test.retMsg, expectedMsg))
			}
		})
	}
}
