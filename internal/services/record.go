package services

import (
	"crypto/sha1"
	"fmt"

	"github.com/XXena/shorter/internal/entities"
	"github.com/XXena/shorter/internal/repository"
)

const salt = "sdlfghdfjhgpsfdhgpadifbvpia"

type RecordService struct {
	r repository.Record
}

func NewRecordService(r repository.Record) *RecordService {
	return &RecordService{r: r}
}

func (s *RecordService) Create(entities.Record) (shortURL string, err error) {

	// todo generate hash from url,
	// pack and send all data to db
	return "", nil

}
func (s *RecordService) GetByURL(longURL string) (shortURL string, err error) {
	return s.r.GetByURL(longURL)
}

func (s *RecordService) Update(recordID int, record entities.Record) error {
	return s.r.Update(recordID, record)
}

func (s *RecordService) Delete(recordID int) error {
	return s.r.Delete(recordID)
}

// todo generate normal short hash
func generateURLHash(url string) string {
	hash := sha1.New()
	hash.Write([]byte(url))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
