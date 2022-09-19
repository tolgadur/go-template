package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	v1 "github.com/tolgadur/email-project/backend/internal/api/v1"
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
		fx.Invoke(getConfigurations),

		fx.Provide(v1.ConnectToDB),
		fx.Invoke(v1.CreateDBSchema),
		fx.Invoke(v1.SeedDB),
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
	logger *zap.SugaredLogger,
) {
	logger.Infof("Starting HTTP server on port %d", viper.Get("http.port"))
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				http.Handle("/", server.Router)
				go startHttpServer(logger, server)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				v1.CloseDB(server.DB, logger)
				return nil
			},
		},
	)
}

func getConfigurations(logger *zap.SugaredLogger) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("Error while reading config file: %s", err)
		panic(err)
	}
}

func startHttpServer(logger *zap.SugaredLogger, server httpv1.Server) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", viper.Get("http.port")), server.Router)
	if err != nil {
		logger.Errorf("Error while starting HTTP server: %s", err)
		panic(err)
	}
}
