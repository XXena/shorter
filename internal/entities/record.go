package entities

import "time"

type Record struct {
	ID         int
	LongURL    string
	Token      string
	CreatedAt  time.Time
	ExpiryDate time.Time
}
