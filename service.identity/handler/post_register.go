package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vhall1/foodlog/service.identity/dao"
)

type registerRequest struct {
	Name string `json:"name"`
}

func (r *Router) PostRegister() http.Handler {
	userDAO := dao.NewUserDAO(r.Database)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req registerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}

		user := dao.User{Name: req.Name}
		if err := userDAO.CreateUser(&user); err != nil {
			log.Default().Printf("Error creating user: %v", err)
			http.Error(w, "could not create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	})
}
