package internal

import (
	"context"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/sapiderman/test-seed/internal/handlers"
	"github.com/sapiderman/test-seed/internal/logger"

	"go.elastic.co/apm/module/apmgorilla"
)

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

	// middleware
	apmgorilla.Instrument(r.MuxRouter)
	r.MuxRouter.Use(logger.MyLogger)

	r.MuxRouter.HandleFunc("/health", handlers.HandlerHealth).Methods("GET")

	// v1 APIs
	v1 := r.MuxRouter.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/hello", handlers.HandlerHello).Methods("GET")

}
