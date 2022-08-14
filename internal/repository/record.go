package repository

import (
	"fmt"

	"github.com/XXena/shorter/pkg/logger"

	"github.com/XXena/shorter/internal/entities"
	"github.com/jackc/pgx"
)

type recordPostgres struct {
	db     *pgx.Conn
	logger logger.Interface
}

func (r recordPostgres) Create(record entities.Record) (recordID int, err error) {
	query := `INSERT INTO records (long_url, token, created_at, expiry_date)
				VALUES ($1, $2, $3, $4)
				RETURNING id`

	err = r.db.QueryRow(query, record.LongURL, record.Token, record.CreatedAt, record.ExpiryDate).Scan(&recordID)
	if err != nil {
		r.logger.Error(fmt.Errorf("db query Create failed: %w", err))

		return recordID, err
	}

	r.logger.Info(fmt.Sprintf("New record ID is: %d", recordID))

	return recordID, err
}

func (r recordPostgres) GetByURL(longURL string) (record entities.Record, err error) {
	query := `SELECT r.id, r.long_url, r.token, r.created_at, r.expiry_date FROM records r
								 WHERE r.long_url = $1`

	err = r.db.QueryRow(query, longURL).Scan(
		&record.ID,
		&record.LongURL,
		&record.Token,
		&record.CreatedAt,
		&record.ExpiryDate)

	if err != nil {
		r.logger.Error(fmt.Errorf("db query GetByURL failed: %w", err))

		return record, err

	}

	return record, err
}

func (r recordPostgres) GetByToken(token string) (record entities.Record, err error) {
	query := `SELECT r.id, r.long_url, r.token, r.created_at, r.expiry_date FROM records r
								 WHERE r.token = $1`
	err = r.db.QueryRow(query, token).Scan(
		&record.ID,
		&record.LongURL,
		&record.Token,
		&record.CreatedAt,
		&record.ExpiryDate)

	if err != nil {
		r.logger.Error(fmt.Errorf("db query GetByToken failed: %w", err))

		return record, err
	}

	return record, err
}

func (r recordPostgres) Update(recordID int, record entities.Record) error {
	//TODO implement me
	panic("implement me")
}

func (r recordPostgres) Delete(recordID int) error {
	//TODO implement me
	panic("implement me")
}

func NewRecordPostgres(db *pgx.Conn, l logger.Interface) *recordPostgres {
	return &recordPostgres{
		db:     db,
		logger: l,
	}
}
