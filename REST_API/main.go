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

	"github.com/joho/godotenv"
	"github.com/rahulkumarpahwa/go/REST_API/internal/config"
	"github.com/rahulkumarpahwa/go/REST_API/internal/http/handlers/student"
	"github.com/rahulkumarpahwa/go/REST_API/internal/storage/sqlite"
)

func main() {

	// Loading .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load Config
	cfg := config.MustLoad()

	// logger

	// database setup
	db, err := sqlite.New(cfg)
	if err != nil {
		log.Fatalf("Failed to Connect DB: %v", err)
	}
	defer db.DB.Close()

	// Student Handler
	studentHandler := student.StudentHandler{Storage: db}

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", Health)
	router.HandleFunc("POST /api/students", studentHandler.CreateStudent)
	router.HandleFunc("GET /api/students/{id}", studentHandler.GetStudentById)
	router.HandleFunc("GET /api/students", studentHandler.GetStudentsList)

	// setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	// using the goroutine to make the graceful shutting down as in the prod level apps.
	// also, now this main will not be waited by one any so we will use 'done' channel to make this stop.

	doneChan := make(chan os.Signal, 1)

	signal.Notify(doneChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server : %v", err)
		}
	}()

	<-doneChan

	slog.Info("Shutting Down the server!")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // closing context to free resources

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server!", slog.String("error", err.Error()))
	}

	slog.Info("server shut down successfully!")

}

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status Available!")
}

/*
Timeline (super important)
Server running
   ↓
User presses Ctrl+C
   ↓
Signal goes into channel (NOT killing program)
   ↓
<-doneChan unblocks
   ↓
Server still running here ⚠️
   ↓
You call server.Shutdown(ctx)
   ↓
Graceful shutdown starts
   ↓
Max wait = 5 seconds
   ↓
Server stops
*/

/*
What is context.Background()?
ctx := context.Background()
It is the root (starting) context in Go.
Think of it as:
“An empty, never-ending context from which all other contexts are derived.”
*/

/*
what we have doing from line no : 47?
We create a buffered channel to receive OS signals. Using signal.Notify, we tell Go to send interrupt signals (like Ctrl+C or SIGTERM) into this channel instead of terminating the program.

The server is started in a goroutine because ListenAndServe() is blocking. Meanwhile, the main goroutine waits on <-doneChan, which blocks until a signal is received.

When Ctrl+C is pressed, the signal is sent to the channel (not killing the program), and the main function resumes execution. At this point, the server is still running.

Then we initiate graceful shutdown using server.Shutdown(ctx), where the context ensures that the shutdown process completes within 5 seconds. Finally, logs confirm the shutdown.
*/

/*
NOTE :
The subtle but VERY important point
When you call:
server.Shutdown(ctx)
👉 It intentionally forces:
server.ListenAndServe()
to return with this error:
http: Server closed (so we need to handle the case :
err != http.ErrServerClosed
)
⚠️ This is NOT a real error
It’s Go saying:
“Hey, I stopped the server because YOU asked me to.”
*/
