package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/yagoluiz/user-api/internal/handlers"
)

func UserRouters(r chi.Router, h *handlers.UserSearchHandler) {
	r.Mount("/api/v1/", r)
	r.Get("/users/search", h.Search)
}
