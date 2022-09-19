package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	httpv1 "github.com/tolgadur/email-project/backend/internal/api/v1/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"path/filepath"

	_ "github.com/lib/pq"
)

const (
	httpPort     = 8080
	host         = "postgresql.default.svc.cluster.local" // replace with localhost for local development
	postgresPort = 5432
	user         = "postgres"
	password     = "AppPassword"
	dbname       = "app_db"
)

func main() {
	fx.New(
		fx.Provide(context.Background),
		fx.Provide(mux.NewRouter),
		fx.Provide(NewLogger),
		fx.Provide(connectToDB),
		fx.Invoke(createDBSchema),
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
				go startHttpServer(logger, server)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				closeDB(server.DB, logger)
				return nil
			},
		},
	)
}

func startHttpServer(logger *zap.SugaredLogger, server httpv1.Server) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), server.Router)
	if err != nil {
		logger.Errorf("Error while starting HTTP server: %s", err)
		panic(err)
	}
}

func closeDB(db *sql.DB, logger *zap.SugaredLogger) {
	logger.Info("Closing DB connection")
	err := db.Close()
	if err != nil {
		logger.Errorf("Error while closing DB connection: %s", err)
		panic(err)
	}
}

func connectToDB(logger *zap.SugaredLogger) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, postgresPort, user, password, dbname)
	logger.Infof("Connecting to DB: %s", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Errorf("Error while connecting to DB: %s", err)
		panic(err)
	}
	logger.Info("Successfully connected to DB!")
	return db
}

func createDBSchema(db *sql.DB, logger *zap.SugaredLogger) {
	path := filepath.Join("db", "schema.sql")
	c, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Errorf("Error while reading schema file: %s", err)
		panic(err)
	}

	schemaSql := string(c)
	_, err = db.Exec(schemaSql)
	if err != nil {
		logger.Errorf("Error while creating DB schema: %s", err)
		panic(err)
	}
}

func seedDB(db *sql.DB, logger *zap.SugaredLogger) {
	path := filepath.Join("db", "seeds.sql")
	c, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Errorf("Error while reading seed file: %s", err)
		panic(err)
	}

	schemaSql := string(c)
	_, err = db.Exec(schemaSql)
	if err != nil {
		logger.Errorf("Error while seeding DB: %s", err)
		panic(err)
	}
}
