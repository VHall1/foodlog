package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vhall1/foodlog/service.identity/dao"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	// Dummy: expects id in query param, returns User if exists
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	// Get DB and DAO
	db, err := getDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userDAO := dao.NewPostgresUserDAO(db)

	user, err := userDAO.GetUserByID(id)
	if err != nil {
		log.Default().Printf("Error getting user by ID %s: %v", id, err)
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
