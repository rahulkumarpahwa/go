package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/rahulkumarpahwa/go/TotionTUI/internal/bubble"
)

func main() {

	p := tea.NewProgram(bubble.Initialization())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	// config

	// setting up the router
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is available!"))
	})

	// routes

	// Listener and Serve
	server := http.Server{
		Addr:        "localhost:8080",
		Handler:     router,
		ReadTimeout: time.Duration(time.Second * 20),
		IdleTimeout: time.Duration(time.Second * 20),
	}
	// Shutdown
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed { 
			// http.ErrServerClosed is the general error which ListenAndServe gives on closing just to show it is closing.
			log.Fatalf("Server didn't start, %v", err)
		}
	}()

	<-signalChannel

	slog.Info("~ Shutdown of the server started ~")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	slog.Info("~ Server Shutdown Successfully ~")

}
