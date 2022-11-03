package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yagoluiz/user-api/internal/usercase"
)

type UserSearchHandler struct {
	userSearchUserCase *usercase.UserSearchUserCase
}

func NewUserSearchHandler(u *usercase.UserSearchUserCase) *UserSearchHandler {
	return &UserSearchHandler{userSearchUserCase: u}
}

func (h *UserSearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.userSearchUserCase.Search("maria")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
