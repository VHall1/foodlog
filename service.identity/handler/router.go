package handler

import (
	"database/sql"
	"net/http"
)

type Router struct {
	Database *sql.DB
}

type handler interface {
	GetLogin() http.Handler
}

func SetupRoutes(r *http.ServeMux, h handler) {
	r.Handle("POST /login", h.GetLogin())
}
