package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	httpv1 "github.com/tolgadur/email-project/backend/internal/api/v1/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func main() {
	fx.New(
		fx.Provide(context.Background),
		fx.Provide(mux.NewRouter),
		fx.Provide(NewLogger),
		fx.Invoke(getConfigurations),

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
	logger.Infof("Starting HTTP server on port %d", viper.Get("http.port"))
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

func closeDB(db *sql.DB, logger *zap.SugaredLogger) {
	logger.Info("Closing DB connection")
	err := db.Close()
	if err != nil {
		logger.Errorf("Error while closing DB connection: %s", err)
		panic(err)
	}
}

func connectToDB(logger *zap.SugaredLogger) *sql.DB {
	host := viper.Get("db.host")
	postgresPort := viper.Get("db.port")
	user := viper.Get("db.user")
	password := viper.Get("db.password")
	dbname := viper.Get("db.name")

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
