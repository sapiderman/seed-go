package internal

import (
	"context"
	"fmt"
	"sapiderman/test-seed/internal/handlers"

	"github.com/gorilla/mux"
)

// Router stores the Mux instance.
type Router struct {
	MuxRouter *mux.Router
}

// NewRouter instantiates and returns new Router
func NewRouter(ctx context.Context) *Router {

	return &Router{
		MuxRouter: mux.NewRouter(),
	}

}

// InitRoutes creates our routes
func (r *Router) InitRoutes() {

	fmt.Println("initializing routes")

	r.MuxRouter.HandleFunc("/health", handlers.HandleHealth).Methods("GET")

}
