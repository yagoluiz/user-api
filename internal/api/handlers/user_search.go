package handlers

import (
	"net/http"
	"strconv"

	"github.com/yagoluiz/user-api/internal/api/dtos"
	"github.com/yagoluiz/user-api/internal/api/responses"
	"github.com/yagoluiz/user-api/internal/usecases"
	"github.com/yagoluiz/user-api/pkg/logger"
)

type UserSearchHandler struct {
	logger             logger.Logger
	userSearchUserCase usecases.UserSearchUseCaseInterface
}

func NewUserSearchHandler(l logger.Logger, u usecases.UserSearchUseCaseInterface) *UserSearchHandler {
	return &UserSearchHandler{logger: l, userSearchUserCase: u}
}

// User godoc
// @Summary		Search users by term
// @Description	Search users by term
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		query	query	string true "Term search"
// @Param		from	query	string true "From page search"
// @Param		size	query	string true "Size page search"
// @Success		200		{array}		dtos.UserDto
// @Failure		400		{object}	dtos.Error
// @Failure		500		{object}	dtos.Error
// @Router		/v1/users/search [get]
func (h *UserSearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		h.logger.Warnf("Query param: %s", query)
		responses.ResponseJSON(w, http.StatusBadRequest, dtos.Error{Message: "Field query required or invalid"})
		return
	}

	from, err := strconv.Atoi(r.URL.Query().Get("from"))
	if err != nil || from < 1 {
		h.logger.Warnf("From param: %v", from)
		responses.ResponseJSON(w, http.StatusBadRequest, dtos.Error{Message: "Field from required or invalid"})
		return
	}

	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size < 1 {
		h.logger.Warnf("Size param: %v", size)
		responses.ResponseJSON(w, http.StatusBadRequest, dtos.Error{Message: "Field size required or invalid"})
		return
	}

	users, err := h.userSearchUserCase.FindUser(query, from-1, size)
	if err != nil {
		h.logger.Errorf("User search error: %s", err.Error())
		responses.ResponseJSON(w, http.StatusInternalServerError, dtos.Error{Message: err.Error()})
		return
	}

	h.logger.Infof("Users count: %d", len(users))

	usersDtos := dtos.DomainToUsersDto(users)

	responses.ResponseJSON(w, http.StatusOK, usersDtos)
}
