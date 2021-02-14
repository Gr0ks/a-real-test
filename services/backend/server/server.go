package server

import (
	"backend/api"
	"backend/api/delivery"
	"backend/api/usecase"
	"backend/api/repository"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"os/signal"
	"log"
	"context"
	"time"
)

type App struct {
	httpServer *http.Server

	useCase api.UseCase
}

func NewApp() *App {
	repo := repository.NewObjectRepository()
	useCase := usecase.NewObjectCreator(repo)
	return &App{
		useCase: useCase,
	}
}

func (a *App) Run(port string) error {
	r := mux.NewRouter()
	delivery.RegisterHTTPEndpoint(r, a.useCase)
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, r))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}