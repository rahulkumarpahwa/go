package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// config

	// setting up the router
	router := http.NewServeMux()

	// routes

	// Listener and Serve
	server := http.Server{
		Addr:        "",
		Handler:     router,
		ReadTimeout: time.Duration(time.Second * 20),
		IdleTimeout: time.Duration(time.Second * 20),
	}

	// Shutdown
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Server didn't start, %v", err)
		}
	}()

	<-signalChannel

	slog.Info("~Shutdown of the server started~")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	slog.Info("~Server Shutdown Successfully~")

}
