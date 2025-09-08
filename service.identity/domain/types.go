package domain

import "time"

type User struct {
	ID        uint32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
