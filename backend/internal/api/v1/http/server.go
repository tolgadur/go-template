package http

import (
	"github.com/gorilla/mux"
	"github.com/tolgadur/email-project/backend/internal/api/v1"
	"go.uber.org/fx"
)

type Server struct {
	fx.In

	Router    *mux.Router
	Codec     Codec
	Endpoints v1.Endpoints
}

func RegisterHttpServer(server Server) error {
	return server.RegisterRoutes()
}
