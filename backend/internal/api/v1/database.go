package v1

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io/ioutil"
	"path/filepath"
)

func CloseDB(db *sql.DB, logger *zap.SugaredLogger) {
	logger.Info("Closing DB connection")
	err := db.Close()
	if err != nil {
		logger.Errorf("Error while closing DB connection: %s", err)
		panic(err)
	}
}

func ConnectToDB(logger *zap.SugaredLogger) *sql.DB {
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

func CreateDBSchema(db *sql.DB, logger *zap.SugaredLogger) {
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

func SeedDB(db *sql.DB, logger *zap.SugaredLogger) {
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
