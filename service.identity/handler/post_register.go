package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vhall1/foodlog/service.identity/store"
)

func (_ *Router) PostRegister(s *store.UserStore) http.Handler {
	type request struct {
		Name string
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}

		user, err := s.Create(req.Name)
		if err != nil {
			log.Default().Printf("Error creating user: %v", err)
			http.Error(w, "could not create user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	})
}
