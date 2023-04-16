package handlers

import (
	"net/http"
	"strconv"

	"github.com/yagoluiz/user-api/internal/api/dtos"
	"github.com/yagoluiz/user-api/internal/api/responses"
	"github.com/yagoluiz/user-api/internal/usecase"
	"github.com/yagoluiz/user-api/pkg/logger"
)

type UserSearchHandler struct {
	logger             logger.Logger
	userSearchUserCase usecase.UserSearchUseCaseInterface
}

func NewUserSearchHandler(l logger.Logger, u usecase.UserSearchUseCaseInterface) *UserSearchHandler {
	return &UserSearchHandler{logger: l, userSearchUserCase: u}
}

// User godoc
// @Summary     Search users by term
// @Description Search users by term
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       query query    string true "Term search"
// @Param       from  query    string true "From page search"
// @Param       size  query    string true "Size page search"
// @Success     200   {object} dtos.UserDto
// @Failure     400
// @Failure     500
// @Router      /v1/users/search [get]
func (h *UserSearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		h.logger.Warnf("Query param: %s", query)
		responses.ResponseJSON(w, http.StatusBadRequest, "Filed query required or invalid")
		return
	}

	from, err := strconv.Atoi(r.URL.Query().Get("from"))
	if err != nil || from < 1 {
		h.logger.Warnf("From param: %v", from)
		responses.ResponseJSON(w, http.StatusBadRequest, "Field from required or invalid")
		return
	}

	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size < 1 {
		h.logger.Warnf("Size param: %v", size)
		responses.ResponseJSON(w, http.StatusBadRequest, "Field size required or invalid")
		return
	}

	users, err := h.userSearchUserCase.FindUser(query, from-1, size)
	if err != nil {
		responses.ResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	h.logger.Infof("Users count: %d", len(users))

	usersDtos := dtos.DomainToUsersDto(users)

	responses.ResponseJSON(w, http.StatusOK, usersDtos)
}
