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

	"github.com/gorilla/mux"
	"github.com/sapiderman/seed-go/api"
	"github.com/sapiderman/seed-go/internal/config"
	"github.com/sapiderman/seed-go/internal/handlers"
	"github.com/sapiderman/seed-go/internal/logger"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgorilla"
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

// Router stores the Mux instance.
type Router struct {
	MuxRouter *mux.Router
}

// NewRouter instantiates and returns new Router
func NewRouter() *Router {

	return &Router{
		MuxRouter: mux.NewRouter(),
	}
}

// InitRoutes creates our routes
func (r *Router) InitRoutes(ctx context.Context) {

	fmt.Println("initializing routes")
	server := ctx.Value(ContextKey(ServerKey)).(*Server)

	// middlewares
	r.MuxRouter.Use(apmgorilla.Middleware()) //	apmgorilla.Instrument(r.MuxRouter) // elastic apm
	r.MuxRouter.Use(logger.MyLogger)         // ye-olde logger

	// health check endpoint. Not in a version path as it will seems to be a permanent endpoint (famous last words)
	h := handlers.NewHealth(ctx)
	r.MuxRouter.HandleFunc("/health", h.Handler)

	// handle swagger api static files as /docs.
	r.MuxRouter.PathPrefix("/docs").Handler(server.StaticFilter)

	// static file handler
	r.MuxRouter.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))

	// v1 APIs
	v1 := r.MuxRouter.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/hello", handlers.HandlerHello).Methods("GET")

	// display routes
	walk(*r.MuxRouter)

}

// StartServer starts listening at given port
func (ws *Server) StartServer(ctx context.Context) {

	var wait time.Duration

	log.Info("initializing server...")

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
	log.Info("shutting down........ byee")
	os.Exit(0)
}

// walk runs the mux.Router.Walk method to print all the registerd routes
func walk(r mux.Router) {
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
