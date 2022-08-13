package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/XXena/shorter/pkg/postgres"

	"github.com/XXena/shorter/pkg/httpserver"

	"github.com/XXena/shorter/internal/repository"
	"github.com/XXena/shorter/internal/services"

	"github.com/XXena/shorter/internal/handlers"

	"github.com/XXena/shorter/pkg/logger"

	"github.com/XXena/shorter/config"
	_ "github.com/jackc/pgx"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	db, err := postgres.NewPostgresDB(cfg.PG, l)
	if err != nil {
		l.Fatal(fmt.Errorf("error occurred while running app: %w", err))
	}

	Repository := repository.NewRepository(db, l)
	// todo defer pg.Close()

	Service := services.NewService(Repository, l)

	Handler := handlers.NewHandler(Service, l)

	srv := new(httpserver.Server)
	go func() {
		fmt.Printf("Listening to port %s \n", cfg.HTTP.Port)
		if err := srv.Run(cfg.HTTP.Port, Handler.InitRoutes()); err != nil {
			l.Fatal(fmt.Errorf("error occurred while running http server: %w", err))
		}
	}()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info(fmt.Sprintf("app - Run - signal: %s", s))
	case err = <-srv.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

}
