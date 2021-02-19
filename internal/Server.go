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
	srvLog = log.WithField("module", "server")

	// StartUpTime records first ime up
	StartUpTime time.Time
	// ServerVersion is a semver versioning
	ServerVersion string
	// HTTPServer object
	HTTPServer *http.Server
	// Router boject
	Router *router.Router
	// Repo is the database object
	Repo *connector.DbPool
	// Handler is all the handlera
	Handler *handlers.MyHandlers

	address string

	// add aditional components here
	// Monitor	*Monitor
	// MessageQ *MessageQ
)

// InitializeServer initializes all server connections
func InitializeServer() error {
	logf := srvLog.WithField("func", "InitializeServer")

	StartUpTime = time.Now()
	ServerVersion = config.Get("app.version")

	Router = router.NewRouter()
	Router.Router = mux.NewRouter()

	logf.Info("connecting database...")
	db, err := connector.NewDbInstance()
	if err != nil {
		return err
	}

	Repo = db
	Handler = handlers.NewHandlers(db)
	Router.Handler = Handler

	logf.Info("initializing routes...")
	router.InitRoutes(Router)

	address := fmt.Sprintf("%s:%s", config.Get("server.host"), config.Get("server.port"))
	HTTPServer = &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15, // Good practice to set timeouts to avoid Slowloris attacks.
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      Router.Router, // Pass our instance of gorilla/mux in.
	}

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

	logf.Info("starting server...")
	logf.Info("listening at: ", address)
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := HTTPServer.ListenAndServe(); err != nil {
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
	HTTPServer.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logf.Info("shutting down........ byee")
	os.Exit(0)
}
