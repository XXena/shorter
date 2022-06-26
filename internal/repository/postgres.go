package repository

import (
	"fmt"
	"strconv"

	"github.com/XXena/shorter/pkg/logger"

	"github.com/jackc/pgx"

	"github.com/XXena/shorter/config"

	_ "github.com/jackc/pgx"
)

const (
	recordsTable = "records"
)

func NewPostgresDB(cfg config.PG, l logger.Interface) (*pgx.Conn, error) {
	port, err := strconv.ParseUint(cfg.Port, 10, 16)
	if err != nil {
		l.Fatal(fmt.Errorf("unable to parse database port: %w", err))
	}

	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     uint16(port),
		Database: cfg.DBName,
		User:     cfg.Username,
		Password: cfg.Password,
		//Logger:               nil, // todo подключить логгер db?
		//LogLevel:             0,
	})

	if err != nil {
		return nil, err
	}

	return conn, nil
}
