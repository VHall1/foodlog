package handler

import (
	"database/sql"
	"net/http"
)

type Router struct {
	Database *sql.DB
}

func SetupRoutes(h *http.ServeMux, r *Router) {
	h.Handle("POST /login", r.PostLogin())
	h.Handle("POST /register", r.PostRegister())
}
