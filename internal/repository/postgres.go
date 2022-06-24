package repository

import (
	"log"
	"strconv"

	"github.com/jackc/pgx"

	"github.com/XXena/shorter/config"

	_ "github.com/jackc/pgx"
)

const (
	recordsTable = "records"
)

func NewPostgresDB(cfg config.PG) (*pgx.Conn, error) {
	//func NewPostgresDB(cfg config.PG) (*sqlx.DB, error) {
	//db, err := sqlx.Open(cfg.DBDriver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	//db, err := sqlx.Open(cfg.DBDriver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	//db, err := sqlx.Open(`postgres`, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	//db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	//if err != nil {
	//	log.Fatalln(err)
	//}

	port, err := strconv.ParseUint(cfg.Port, 10, 16)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     uint16(port),
		Database: cfg.DBName,
		User:     cfg.Username,
		Password: cfg.Password,
		//Logger:               nil, // todo подключить логгер
		//LogLevel:             0,
	})

	if err != nil {
		return nil, err
	}

	//err = db.Ping()
	//
	//if err != nil {
	//	return nil, err
	//}

	//return db, nil
	return conn, nil
}
