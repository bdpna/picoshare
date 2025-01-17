package handlers

import (
	"github.com/gorilla/mux"

	"github.com/mtlynch/picoshare/v2/garbagecollect"
	"github.com/mtlynch/picoshare/v2/handlers/auth"
	"github.com/mtlynch/picoshare/v2/store"
)

type Server struct {
	router        *mux.Router
	authenticator auth.Authenticator
	store         store.Store
	collector     *garbagecollect.Collector
}

// Router returns the underlying router interface for the server.
func (s Server) Router() *mux.Router {
	return s.router
}

// New creates a new server with all the state it needs to satisfy HTTP
// requests.
func New(authenticator auth.Authenticator, store store.Store, collector *garbagecollect.Collector) Server {
	s := Server{
		router:        mux.NewRouter(),
		authenticator: authenticator,
		store:         store,
		collector:     collector,
	}

	s.routes()
	return s
}
