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

	return conn, nil
}
