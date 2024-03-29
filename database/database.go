package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-ApmCommon/logger"
	"go-ApmCommon/model"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

var config model.TomlConfig
var log *logrus.Entry

func Init() {
	config = *model.GetConfig()
	log = logger.New(context.Background())
}

const (
	DatabaseUser     = "admin"
	DatabasePassword = "admin"
	DatabaseName     = "postgres"
)

type DataAccess interface {
	Get(id string) (*model.User, error)
}

func NewOpenDB() *sql.DB {
	dbInfo := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		DatabaseUser,
		DatabasePassword,
		DatabaseName,
		config.Databases["postgres"].Server,
		config.Databases["postgres"].Port,
	)
	db, err := apmsql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
		panic("Invalid DB config")
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		panic("DB unreachable")
	}
	log.Debug("connected DB")
	return db
}
