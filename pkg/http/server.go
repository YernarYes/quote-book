package httpserver

import (
	"fmt"
	"net/http"
	"quotes/internal/core/handler"

	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	handler *handler.Handler
}

func NewHTTPServer(handler *handler.Handler, router *mux.Router) *Server {
	RoutesRegister(handler, router)

	return &Server{
		handler: handler,
		router:  router,
	}
}

func (s *Server) StartHTTPServer(host, port string) error {
	address := fmt.Sprintf("%s:%s", host, port)
	if err := http.ListenAndServe(address, s.router); err != nil {
		return fmt.Errorf("failed to start HTTPServer %v", err)
	}

	return nil
}
