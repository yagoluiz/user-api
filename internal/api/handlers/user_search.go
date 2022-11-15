package handlers

import (
	"net/http"
	"strconv"

	"github.com/yagoluiz/user-api/internal/api/responses"
	"github.com/yagoluiz/user-api/internal/usercase"
)

type UserSearchHandler struct {
	userSearchUserCase *usercase.UserSearchUserCase
}

func NewUserSearchHandler(u *usercase.UserSearchUserCase) *UserSearchHandler {
	return &UserSearchHandler{userSearchUserCase: u}
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
// @Success     200   {object} entity.User
// @Failure     400
// @Failure     500
// @Router      /v1/users/search [get]
func (h *UserSearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		responses.ResponseJSON(w, http.StatusBadRequest, "Filed query required or invalid")
		return
	}
	from, err := strconv.Atoi(r.URL.Query().Get("from"))
	if err != nil || from < 0 {
		responses.ResponseJSON(w, http.StatusBadRequest, "Field from required or invalid")
		return
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size < 1 {
		responses.ResponseJSON(w, http.StatusBadRequest, "Field size required or invalid")
		return
	}

	users, err := h.userSearchUserCase.Search(query, from, size)
	if err != nil {
		responses.ResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.ResponseJSON(w, http.StatusOK, users)
}
