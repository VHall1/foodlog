package store

import (
	"database/sql"
	"errors"
	"time"

	"github.com/vhall1/foodlog/service.identity/domain"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) Create(name string) (*domain.User, error) {
	query := `INSERT INTO "User" ("name", "updatedAt") VALUES ($1, $2) RETURNING "id", "name", "createdAt", "updatedAt"`
	updatedAt := time.Now()
	return scanUser(s.db.QueryRow(query, name, updatedAt))
}

func (s *UserStore) FindByID(id uint32) (*domain.User, error) {
	query := `SELECT "id", "name", "createdAt", "updatedAt" FROM "User" WHERE "id" = $1`
	return scanUser(s.db.QueryRow(query, id))
}

func scanUser(row *sql.Row) (*domain.User, error) {
	user := &domain.User{}
	if err := row.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt); err != nil {
		// return nil if no rows found
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
