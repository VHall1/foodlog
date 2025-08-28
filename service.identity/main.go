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

	"github.com/vhall1/foodlog/service.identity/handler"
)

func main() {
	router := handler.NewRouter()
	router.SetupRoutes()

	fmt.Println("Starting server on :80")

	go func() {
		if err := router.Start(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := router.Shutdown(ctx); err != nil {
		fmt.Printf("Shutdown error: %v\n", err)
	}

	fmt.Println("Server gracefully stopped")
}
