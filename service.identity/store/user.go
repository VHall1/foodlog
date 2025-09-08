package store

import (
	"database/sql"
	"time"

	"github.com/vhall1/foodlog/service.identity/domain"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) CreateUser(user *domain.User) error {
	query := `INSERT INTO "User" ("name", "updatedAt") VALUES ($1, $2) RETURNING "id", "createdAt", "updatedAt"`
	updatedAt := time.Now()
	return s.db.QueryRow(query, user.Name, updatedAt).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (s *UserStore) GetUserByID(id uint32) (*domain.User, error) {
	query := `SELECT "id", "name", "createdAt", "updatedAt" FROM "User" WHERE "id" = $1`
	user := &domain.User{}
	err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserStore) UpdateUser(user *domain.User) error {
	query := `UPDATE "User" SET "name" = $1, "updatedAt" = $2 WHERE "id" = $3 RETURNING "name", "updatedAt"`
	updatedAt := time.Now()
	return s.db.QueryRow(query, user.Name, updatedAt, user.ID).Scan(&user.Name, &user.UpdatedAt)
}
