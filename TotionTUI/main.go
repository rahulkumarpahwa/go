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
	bubble "github.com/rahulkumarpahwa/go/TotionTUI/internal/TUI"
	"github.com/rahulkumarpahwa/go/TotionTUI/internal/routes/notes"
	"github.com/rahulkumarpahwa/go/TotionTUI/internal/storage"
)

func main() {
	// A. First Starting the Server
	
	// config


	//setting the DB
	DB, err := storage.Open()
	if err != nil {
		log.Fatalf("DB didn't open, %v", err)
	}

	defer DB.Close()

	// setting up the storage
	notesStorage := storage.NotesStorage{DB: DB}

	// setting up the router
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is available!"))
	})

	// routes
	notesHandler := notes.NotesHandler{Storage: notesStorage}
	router.HandleFunc("POST /api/notes", notesHandler.CreateNote)

	// Listener and Serve
	server := http.Server{
		Addr:        "localhost:8080",
		Handler:     router,
		ReadTimeout: time.Duration(time.Second * 20),
		IdleTimeout: time.Duration(time.Second * 20),
	}

	// Start & Shutdown
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			// http.ErrServerClosed is the general error which ListenAndServe gives on closing just to show it is closing.
			log.Fatalf("Server didn't start, %v", err)
		}
	}()

	go func() {
		// B. Second Showing : Bubble Tea TUI
		p := tea.NewProgram(bubble.Initialization())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	}()

	<-signalChannel

	slog.Info("~ Server Shutdown Begin ~")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	slog.Info("~ Server Shutdown Successfully ~")

}
