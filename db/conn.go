package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "go_db"
	port = 5432
	user = "docker"
	password = "docker"
	db_name = "my-store-go"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + db_name)

	return db, nil
}