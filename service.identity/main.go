package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vhall1/foodlog/lib/bootstrap"
	"github.com/vhall1/foodlog/lib/util"
	"github.com/vhall1/foodlog/service.identity/handler"
)

func main() {
	svc := bootstrap.NewService("service.identity")

	mux := http.NewServeMux()
	handler.SetupRoutes(mux, &handler.Router{
		Database: svc.Postgres(),
	})

	s := svc.NewHttpServer(util.LoggerMiddleware(mux))

	fmt.Println("Starting server on :80")

	go func() {
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		fmt.Printf("Shutdown error: %v\n", err)
	}

	fmt.Println("Server gracefully stopped")
}
