package internal

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/sapiderman/seed-go/api"
	"github.com/sapiderman/seed-go/internal/config"
	log "github.com/sirupsen/logrus"
)

// Server struct is your server definitions, put your configs here
type Server struct {
	Host string
	Port int

	StartUpTime   time.Time
	ServerVersion string

	HTTPServer   *http.Server
	Router       *Router
	StaticFilter *api.StaticFilter

	// add aditional components here
	// Monitor	*Monitor
	// Database	*Database
	// MessageQ *MessageQ
}

// NewServer initializes server object
func NewServer(ctx context.Context) *Server {

	cfg := ctx.Value(ContextKey(ConfigKey)).(*config.Configuration)
	server := &Server{
		Router: NewRouter(),
	}

	server.HTTPServer = &http.Server{
		Addr: fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      server.Router.MuxRouter, // Pass our instance of gorilla/mux in.
	}

	// set our handler for static files
	server.StaticFilter = api.NewStaticFilter()

	server.StartUpTime = time.Now()
	server.ServerVersion = strings.Join([]string{VersionBuild, VersionMinor, VersionPatch}, ".")

	return server
}

// StartServer starts listening at given port
func (ws *Server) StartServer(ctx context.Context) {

	var wait time.Duration

	log.Info("initiaing server...")

	serverCtxKey := ContextKey(ServerKey)
	serverCtx := context.WithValue(ctx, serverCtxKey, ws)
	ws.Router.InitRoutes(serverCtx)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := ws.HTTPServer.ListenAndServe(); err != nil {
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
	ws.HTTPServer.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down........ byee")
	os.Exit(0)
}
