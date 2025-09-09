package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vhall1/foodlog/service.nutrition/store"
)

func (_ *Router) GetDay(s *store.DayStore) http.Handler {
	type request struct {
		ID uint32
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		day, err := s.GetByID(req.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(day)
	})
}
