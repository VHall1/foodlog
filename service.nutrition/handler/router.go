package handler

import (
	"database/sql"
	"net/http"

	"github.com/vhall1/foodlog/service.nutrition/store"
)

type Router struct {
	Database *sql.DB
}

func SetupRoutes(h *http.ServeMux, r *Router) {
	dayStore := store.NewDayStore(r.Database)
	h.Handle("GET /days/{id}", r.GetDay(dayStore))
	h.Handle("POST /days", r.PostDay(dayStore))
}
