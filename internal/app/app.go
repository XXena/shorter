package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	db, err := repository.NewPostgresDB(cfg.PG)
	if err != nil {
		log.Fatalf("error ocurred while running app: %s", err.Error())
	}

	Repository := repository.NewRepository(db)
	// todo defer pg.Close()

	Service := services.NewService(Repository)

	Handler := handlers.NewHandler(Service)

	srv := new(httpserver.Server)
	go func() {
		fmt.Printf("Listening to port %s \n", cfg.HTTP.Port)
		if err := srv.Run(cfg.HTTP.Port, Handler.InitRoutes()); err != nil {
			log.Fatalf("error ocurred while running http server: %s", err.Error())
		}
	}()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
		// todo добавить notify (для ловли ошибок без падения?)
		//  case err = <-httpServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

}
