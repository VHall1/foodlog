package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vhall1/foodlog/service.identity/store"
)

func (r *Router) PostLogin(s *store.UserStore) http.Handler {
	type request struct {
		Id uint32 `json:"id"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}

		user, err := s.GetUserByID(req.Id)
		if err != nil {
			log.Default().Printf("Error getting user by ID %d: %v", req.Id, err)
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

}
