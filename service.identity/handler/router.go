package handler

import (
	"database/sql"
	"net/http"

	"github.com/vhall1/foodlog/service.identity/store"
)

type Router struct {
	Database *sql.DB
}

func SetupRoutes(h *http.ServeMux, r *Router) {
	userStore := store.NewUserStore(r.Database)
	h.Handle("POST /login", r.PostLogin(userStore))
	h.Handle("POST /register", r.PostRegister(userStore))
}
