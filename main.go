package main

import (
	"log"

	"github.com/XXena/shorter/internal/app"

	"github.com/XXena/shorter/config"
	_ "github.com/jackc/pgx"
)

// todo переместить в cmd
func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
