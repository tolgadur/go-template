package http

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/tolgadur/email-project/backend/internal/api/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	fx.In

	Logger    *zap.SugaredLogger
	Router    *mux.Router
	Codec     Codec
	Endpoints v1.Endpoints
	DB        *sql.DB
}

func RegisterHttpServer(server Server) error {
	return server.RegisterRoutes()
}
