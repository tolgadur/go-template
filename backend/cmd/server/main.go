package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	httpv1 "github.com/tolgadur/email-project/backend/internal/api/v1/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	httpPort     = 8080
	host         = "postgresql.default.svc.cluster.local"
	postgresPort = 5432
	user         = "app1"
	password     = "AppPassword"
	dbname       = "app_db"
)

func main() {
	fx.New(
		fx.Provide(context.Background),
		fx.Provide(mux.NewRouter),
		fx.Provide(NewLogger),
		fx.Invoke(connectToDB),
		fx.Invoke(seedDB),
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
	logger.Infof("Starting HTTP server on port %d", httpPort)
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				http.Handle("/", server.Router)
				go http.ListenAndServe(fmt.Sprintf(":%s", httpPort), server.Router)
				return nil
			},
		},
	)
}

func connectToDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, postgresPort, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func seedDB() {

}
