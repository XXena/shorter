package entities

import "time"

type Record struct {
	LongURL    string
	Token      string
	CreatedAt  time.Time
	ExpiryDate time.Time
}
