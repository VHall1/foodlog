package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/vhall1/foodlog/service.nutrition/store"
)

func (_ *Router) PostDay(s *store.DayStore) http.Handler {
	type request struct {
		Date     time.Time
		Calories int
		UserID   uint32
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		day, err := s.Create(req.Date, req.Calories, req.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(day)
	})
}
