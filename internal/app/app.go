package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"refactoring/internal/config"
	"refactoring/internal/delivery"
	"refactoring/internal/repository"
	"refactoring/internal/server"
	"refactoring/internal/service"
	"syscall"
	"time"
)

const timeout = 5 * time.Second

// @title TrueConf Backend Test
// @version 2.0
// @description Тех. Задание TrueConf

// @host localhost:3333
// @BasePath /
func Run(configDir string) {
	// при добавлении контейнеризации, очевидно, надо убрать
	os.Setenv("HTTP_HOST", "localhost")

	cfg, err := config.InitConfig(configDir)
	if err != nil {
		log.Fatal("Error occurred while loading config: ", err.Error())
	}

	repos := repository.NewRepository(
		repository.NewUserRepo(cfg.DB.UsersStore),
	)

	services := service.NewServices(
		service.NewUserService(repos),
	)

	h := delivery.NewHandler(services)
	mux := h.Init()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	srv := server.NewServer(cfg, mux)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic occurred: %s\n", err)
			}
		}()
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Println("error happened: ", err.Error())
		}
	}()

	log.Println("Application is running")

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	log.Println("Application is shutting down")

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("error occurred on server shutting down: %s", err.Error())
	}
}
