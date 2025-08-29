package dao

import (
	"database/sql"
	"time"
)

type PostgresUserDAO struct {
	db *sql.DB
}

func NewPostgresUserDAO(db *sql.DB) *PostgresUserDAO {
	return &PostgresUserDAO{db: db}
}

func (d *PostgresUserDAO) CreateUser(user *User) error {
	query := `INSERT INTO "User" ("name", "updatedAt") VALUES ($1, $2) RETURNING "id", "createdAt", "updatedAt"`
	updatedAt := time.Now()
	return d.db.QueryRow(query, user.Name, updatedAt).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (d *PostgresUserDAO) GetUserByID(id string) (*User, error) {
	query := `SELECT "id", "name", "createdAt", "updatedAt" FROM "User" WHERE "id" = $1`
	user := &User{}
	err := d.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *PostgresUserDAO) UpdateUser(user *User) error {
	query := `UPDATE "User" SET "name" = $1, "updatedAt" = $2 WHERE "id" = $3 RETURNING "name", "updatedAt"`
	updatedAt := time.Now()
	return d.db.QueryRow(query, user.Name, updatedAt, user.ID).Scan(&user.Name, &user.UpdatedAt)
}

func (d *PostgresUserDAO) DeleteUser(id string) error {
	query := `DELETE FROM "User" WHERE "id" = $1`
	_, err := d.db.Exec(query, id)
	return err
}
