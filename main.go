package main

import (
	"log"

	"github.com/XXena/shorter/internal/app"

	"github.com/XXena/shorter/config"
	_ "github.com/jackc/pgx"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	//conn, err := pgx.Connect(pgx.ConnConfig{
	//	Host:     "localhost",
	//	Port:     5436,
	//	Database: "postgres",
	//	User:     "postgres",
	//	Password: "password",
	//})
	//if err != nil {
	//	log.Fatalf("Unable to connect to database: %v\n", err)
	//}
	//defer func(conn *pgx.Conn) {
	//	err := conn.Close()
	//	if err != nil {
	//		log.Fatalf("conn close error: %s", err)
	//
	//	}
	//}(conn)
	//var longUrl string
	//var token string
	//err = conn.QueryRow("select long_url, token from records where id=$1", 1).Scan(&longUrl, &token)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(longUrl, token)
	// Run
	app.Run(cfg)
}
