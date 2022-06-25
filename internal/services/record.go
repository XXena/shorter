package services

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/XXena/shorter/config"

	"github.com/itchyny/base58-go"

	"github.com/XXena/shorter/internal/entities"
	"github.com/XXena/shorter/internal/helpers"
	"github.com/XXena/shorter/internal/repository"
)

type RecordService struct {
	r repository.Record
}

func NewRecordService(r repository.Record) *RecordService {
	return &RecordService{r: r}
}

func (s *RecordService) Create(record entities.Record) (string, error) {

	record.Token = GenerateShortLink(record.LongURL)
	_, err := s.r.Create(record)

	return config.ShortAddr + record.Token, err

}
func (s *RecordService) GetByURL(longURL string) (string, error) {
	record, err := s.r.GetByURL(longURL)

	if !(helpers.InTime(record.ExpiryDate, time.Now())) {
		//todo если срок действия истек, генеринруется новый хэш и возвращается наружу
	}

	return config.ShortAddr + record.Token, err
}

func (s *RecordService) Update(recordID int, record entities.Record) error {
	return s.r.Update(recordID, record)
}

func (s *RecordService) Delete(recordID int) error {
	return s.r.Delete(recordID)
}

func GenerateShortLink(longURL string) (shortURL string) {
	urlHashBytes := sha256gen(longURL)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Gen([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

func base58Gen(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func sha256gen(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hash.Sum(nil)
}
