package services

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/XXena/shorter/pkg/logger"

	"github.com/itchyny/base58-go"

	"github.com/XXena/shorter/internal/entities"
	"github.com/XXena/shorter/internal/helpers"
	"github.com/XXena/shorter/internal/repository"
)

type RecordService struct {
	Repo   repository.Record
	logger logger.Interface
}

func NewRecordService(r repository.Record, l logger.Interface) *RecordService {
	return &RecordService{
		Repo:   r,
		logger: l,
	}
}

func (s *RecordService) Create(record entities.Record) (string, error) {

	record.Token = GenerateShortLink(record.LongURL, s.logger)
	_, err := s.Repo.Create(record)

	return record.Token, err

}
func (s *RecordService) GetByURL(longURL string) (string, error) {
	record, err := s.Repo.GetByURL(longURL)

	if !(helpers.InTime(record.ExpiryDate, time.Now())) {
		//todo если срок действия истек, генерируется новый хэш и возвращается наружу
	}

	return record.Token, err
}

func (s *RecordService) Redirect(shortURL string) (string, error) {
	token := getToken(shortURL)
	record, err := s.Repo.GetByToken(token)
	if !(helpers.InTime(record.ExpiryDate, time.Now())) {
		//todo если срок действия истек, вернуть ошибку о сроке действия
	}

	return record.LongURL, err
}

func (s *RecordService) Update(recordID int, record entities.Record) error {
	return s.Repo.Update(recordID, record)
}

func (s *RecordService) Delete(recordID int) error {
	return s.Repo.Delete(recordID)
}

func GenerateShortLink(longURL string, l logger.Interface) (shortURL string) {
	urlHashBytes := sha256gen(longURL)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Gen([]byte(fmt.Sprintf("%d", generatedNumber)), l)
	return finalString[:8]
}

func base58Gen(bytes []byte, l logger.Interface) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		l.Fatal(fmt.Errorf("encoding error: %w", err))
	}
	return string(encoded)
}

func sha256gen(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hash.Sum(nil)
}

func getToken(shortURL string) string {
	return strings.TrimLeft(shortURL, "localhost/")
}
