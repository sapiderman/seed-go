package internal

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sapiderman/seed-go/internal/config"
	"github.com/sapiderman/seed-go/internal/connector"
	"github.com/sapiderman/seed-go/internal/router"
	log "github.com/sirupsen/logrus"
)

var server Server

// Server struct is your server definitions, put your configs here
type Server struct {
	Host string
	Port int

	StartUpTime   time.Time
	ServerVersion string

	HTTPServer *http.Server
	Router     *mux.Router

	// add aditional components here
	// Monitor	*Monitor
	// Database	*Database
	// MessageQ *MessageQ
}

// StartServer starts listening at given port
func StartServer() {

	var wait time.Duration

	log.Info("initializing server...")
	server.Router = mux.NewRouter()

	log.Info("initializing routes...")
	router.InitRoutes(server.Router)

	log.Info("connecting database...")
	connector.InitializeDBInstance()

	address := fmt.Sprintf("%s:%s", config.Get("server.host"), config.Get("server.port"))
	server.HTTPServer = &http.Server{
		Addr: address,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      server.Router, // Pass our instance of gorilla/mux in.
	}

	log.Info("starting server...")
	log.Info("listening at %s", address)
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.HTTPServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	gracefulStop := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(gracefulStop, os.Interrupt)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	// Block until we receive our signal.
	<-gracefulStop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.HTTPServer.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("shutting down........ byee")
	os.Exit(0)
}
