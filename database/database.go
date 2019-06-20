package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/nodias/go-ApmCommon/model"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

const (
	DatabaseUser     = "admin"
	DatabasePassword = "admin"
	DatabaseName     = "postgres"
)

type DataAccess interface {
	Get(id string) (*model.User, error)
}

type PostgreAccess struct {
	*sql.DB
}

func GetUserInfo(ctx context.Context, id string) (*model.User, error) {
	db := NewOpenDB()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	user := model.User{}
	defer db.Close()
	row := tx.QueryRowContext(ctx, "SELECT * FROM schema_user.user WHERE id = $1", id)
	err = row.Scan(&user.Id, &user.Name)
	if err != nil {
		log.Printf("database.PostgreAccess.Get - %s", err)
		return nil, err
	}
	log.Printf("database.PostgreAccess.Get - id: %s, name: %s", user.Id, user.Name)
	return &user, nil
}

func NewOpenDB() *sql.DB {
	dbInfo := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=54.180.2.92 port=5432",
		DatabaseUser,
		DatabasePassword,
		DatabaseName,
	)
	db, err := apmsql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Invalid DB config : ", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DB unreachable : ", err)
	}
	log.Println("connected DB")
	return db
}
