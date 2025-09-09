package store

import (
	"database/sql"
	"time"

	"github.com/vhall1/foodlog/service.nutrition/domain"
)

type DayStore struct {
	db *sql.DB
}

func NewDayStore(db *sql.DB) *DayStore {
	return &DayStore{db: db}
}

func (s *DayStore) Create(date time.Time, calories int, userID uint32) (*domain.Day, error) {
	query := `INSERT INTO "Day" ("date", "calories", "userId", "updatedAt") VALUES ($1, $2, $3, $4) RETURNING "id", "date", "calories", "userId", "createdAt", "updatedAt"`
	updatedAt := time.Now()
	row := s.db.QueryRow(query, date, calories, userID, updatedAt)
	return scanDay(row)
}

func (s *DayStore) GetByID(id uint32) (*domain.Day, error) {
	query := `SELECT "id", "date", "calories", "userId", "createdAt", "updatedAt" FROM "Day" WHERE "id" = $1`
	row := s.db.QueryRow(query, id)
	return scanDay(row)
}

func scanDay(row *sql.Row) (*domain.Day, error) {
	day := &domain.Day{}
	err := row.Scan(&day.ID, &day.Date, &day.Calories, &day.UserID, &day.CreatedAt, &day.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return day, nil
}
