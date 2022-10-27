package handlers

import (
	"net/http"

	"github.com/yagoluiz/user-api/internal/usercase"
)

type UserSearchHandler struct {
	UserSearchUserCase *usercase.UserSearchUserCase
}

func NewUserSearchHandler(u *usercase.UserSearchUserCase) *UserSearchHandler {
	return &UserSearchHandler{UserSearchUserCase: u}
}

func (h *UserSearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Search"))
}
