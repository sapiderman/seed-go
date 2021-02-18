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
	"github.com/sapiderman/seed-go/internal/handlers"
	"github.com/sapiderman/seed-go/internal/router"

	log "github.com/sirupsen/logrus"
)

var (
	server Server
	srvLog = log.WithField("module", "sqlx")
)

// Server struct is your server definitions, put your configs here
type Server struct {
	StartUpTime   time.Time
	ServerVersion string
	HTTPServer    *http.Server
	Router        *router.Router
	Repo          *connector.DbPool
	Handler       *handlers.MyHandlers

	// add aditional components here
	// Monitor	*Monitor
	// MessageQ *MessageQ
}

// InitializeServer initializes all server connections
func InitializeServer() error {
	logf := srvLog.WithField("func", "newServer")

	server.StartUpTime = time.Now()
	server.ServerVersion = config.Get("app.version")

	server.Router.Router = mux.NewRouter()

	logf.Info("connecting database...")
	db, err := connector.NewDbInstance()
	if err != nil {
		return err
	}

	server.Repo = db
	server.Handler, err = handlers.NewHandlers(db)
	if err != nil {
		return err
	}
	server.Router.Handler = server.Handler

	return nil
}

// StartServer starts listening at given port
func StartServer() {

	var wait time.Duration
	logf := srvLog.WithField("func", "StartServer")

	logf.Info("initializing server...")
	err := InitializeServer()
	if err != nil {
		logf.Error(err)
	}

	logf.Info("initializing routes...")
	router.InitRoutes(server.Router)

	address := fmt.Sprintf("%s:%s", config.Get("server.host"), config.Get("server.port"))
	server.HTTPServer = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15, // Good practice to set timeouts to avoid Slowloris attacks.
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      server.Router.Router, // Pass our instance of gorilla/mux in.
	}

	logf.Info("starting server...")
	logf.Info("listening at: ", address)
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.HTTPServer.ListenAndServe(); err != nil {
			logf.Error(err)
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
	logf.Info("shutting down........ byee")
	os.Exit(0)
}
