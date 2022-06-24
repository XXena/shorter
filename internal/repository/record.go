package repository

import (
	"fmt"

	"github.com/XXena/shorter/internal/entities"
	"github.com/jackc/pgx"
)

type RecordPostgres struct {
	db *pgx.Conn
	//db *sqlx.DB
}

func (r RecordPostgres) Create(record entities.Record) (recordID int, err error) {
	query := `INSERT INTO records (long_url, token, created_at, expiry_date)
				VALUES ($1, $2, $3, $4)
				RETURNING id`

	//todo обработка дат
	err = r.db.QueryRow(query, record.LongURL, record.Token, record.CreatedAt, record.ExpiryDate).Scan(&recordID)
	if err != nil {
		fmt.Printf("QueryRow Create failed: %v\n", err)
	}
	fmt.Println("New record ID is:", recordID)

	return recordID, err
}

func (r RecordPostgres) GetByURL(longURL string) (record entities.Record, err error) {
	query := fmt.Sprintf(`SELECT r.id, r.long_url, r.token, r.created_at, r.expiry_date FROM %s r
								 WHERE r.long_url = $1`, recordsTable)

	err = r.db.QueryRow(query, longURL).Scan(
		&record.ID,
		&record.LongURL,
		&record.Token,
		&record.CreatedAt,
		&record.ExpiryDate)

	// todo struct scan https://www.randylough.com/projects/pgxscan  ?
	if err != nil {
		fmt.Printf("QueryRow GetByURL failed: %v\n", err)
	}

	return record, err
}

func (r RecordPostgres) Update(recordID int, record entities.Record) error {
	//TODO implement me
	panic("implement me")
}

func (r RecordPostgres) Delete(recordID int) error {
	//TODO implement me
	panic("implement me")
}

func NewRecordPostgres(db *pgx.Conn) *RecordPostgres {
	//func NewRecordPostgres(db *sqlx.DB) *RecordPostgres {
	return &RecordPostgres{db: db}
}
