package main

import (
	"fmt"
	"log"

	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
	CREATE TABLE projects (
		id text PRIMARY KEY,
		name text,
		status text
	);`

type Project struct {
	ID     string `db:"id"`
	Name   string `db:"name"`
	Status string `db:"status"`
}

func createConnDB(config *Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password='%s' dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.dbName,
	)

	con, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		glog.Fatalf("Failed to create a database connection: %v", err)
	}

	_, err = con.Exec(schema)
	log.Print(err)

	return con, nil
}

func buildDB(config *Config) *sqlx.DB {

	db, errDB := createConnDB(config)
	if errDB != nil {
		glog.Fatalf("failed to connect with database: %v", errDB)
	}

	//defer db.Close()

	errDB = db.Ping()
	if errDB != nil {
		glog.Fatalf("failed to connect with database: %v", errDB)
	}

	log.Println("Connection Successfully")

	return db
}
