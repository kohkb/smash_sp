package domain

import (
	"time"
)

type Favorite struct {
	ID        uint
	VideoId   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
