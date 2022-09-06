package services

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/XXena/shorter/pkg/logger"

	"github.com/itchyny/base58-go"

	"github.com/XXena/shorter/internal/entities"
	"github.com/XXena/shorter/internal/helpers"
	"github.com/XXena/shorter/internal/repository"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	saltLen     = 5
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

	link, err := s.GenerateShortLink(record.LongURL, s.logger)
	if err != nil {

		return "", err
	}

	record.Token = link
	_, err = s.Repo.Create(record)

	return record.Token, err

}

func (s *RecordService) ForwardToCreate(url string, expiry time.Time) ([]byte, error) {
	now := time.Now()

	if expiry.IsZero() {
		expiry = now.AddDate(1, 0, 0) // todo нужен настраиваемый срок действия
	}

	token, err := s.Create(entities.Record{
		LongURL:    url,
		CreatedAt:  now,
		ExpiryDate: expiry,
	})

	if err != nil {
		return []byte(token), err
	}

	return []byte(token), nil

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

func (s *RecordService) GenerateShortLink(longURL string, l logger.Interface) (shortURL string, err error) {
	salt := saltGen(saltLen)
	urlHashBytes := sha256gen(longURL + salt)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString, err := base58Gen([]byte(fmt.Sprintf("%d", generatedNumber)), l)

	if err != nil {
		return "", err
	}

	ok, err := s.checkHashCollision(finalString[:8], l)
	if (err != nil) || (!ok) {
		if strings.Contains(err.Error(), "hash collision found") {
			finalString, err := s.GenerateShortLink(longURL, l) // todo норм ли вызывать рекурсивно?
			if err != nil {

				return "", err
			}

			return finalString[:8], err
		}

		return "", err
	}

	return finalString[:8], err
}

func (s *RecordService) checkHashCollision(hash string, l logger.Interface) (bool, error) {
	record, err := s.Repo.GetByToken(hash)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") { // todo типизация ошибок

			return true, nil
		}

		l.Error(fmt.Errorf("checking hash collision error: %w", err))

		return false, err
	}

	if record.Token == hash {
		l.Error(fmt.Errorf("hash collision found: %w", err))

		return false, err
	}

	return false, err
}

func base58Gen(bytes []byte, l logger.Interface) (string, error) {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		l.Error(fmt.Errorf("encoding error: %w", err))
	}

	return string(encoded), err
}

func sha256gen(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))

	return hash.Sum(nil)
}

func saltGen(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func getToken(shortURL string) string {
	return strings.TrimLeft(shortURL, "localhost/")
}
