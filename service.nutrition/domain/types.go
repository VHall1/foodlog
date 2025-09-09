package domain

import "time"

type Day struct {
	ID        uint32
	Date      time.Time
	Calories  int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint32
}

type Foodlog struct {
	ID          uint32
	Date        time.Time
	Calories    int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uint32
}
