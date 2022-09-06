package test

import (
	"time"

	"github.com/XXena/shorter/internal/entities"
)

const (
	ID      = 9
	LongURL = "https://engineering.atspotify.com/2020/04/when-should-i-write-an-architecture-decision-record/"
	Token   = "XAJjKHF4"
)

func NewFakeRecord() entities.Record {
	now := time.Now()

	return entities.Record{
		ID:         ID,
		LongURL:    LongURL,
		Token:      Token,
		CreatedAt:  now,                  // 0001-01-01T00:00:01Z
		ExpiryDate: now.AddDate(1, 0, 0), // 26132-08-16T01:41:32Z
	}
}

//func NewRandomFakeRecord() entities.Record {
//	now := time.Now()
//	rand.Seed(time.Now().UnixNano())
//	str := RandSeq(10)
//
//	return entities.Record{
//		ID:         rand.Int(),
//		LongURL:    "https://" + str + "/",
//		Token:      Token, // todo нужен рандомный
//		CreatedAt:  now,
//		ExpiryDate: now.AddDate(1, 0, 0),
//	}
//}
