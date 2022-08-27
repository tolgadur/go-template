package main

import (
	"context"
	"github.com/gorilla/mux"
	httpv1 "github.com/tolgadur/email-project/backend/internal/api/v1/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(context.Background),
		fx.Provide(mux.NewRouter),
		fx.Provide(NewLogger),
		fx.Invoke(httpv1.RegisterHttpServer),
		fx.Invoke(registerHooks),
	).Run()
}

func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}

func registerHooks(
	lifecycle fx.Lifecycle, server httpv1.Server,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				http.Handle("/", server.Router)
				go http.ListenAndServe(":8080", server.Router)
				return nil
			},
		},
	)
}
