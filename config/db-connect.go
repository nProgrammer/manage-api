package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	pgUrl, _ := pq.ParseURL(os.Getenv("ELEPHANT_URL"))
	fmt.Println(os.Getenv("ELEPHANT_URL"))
	db, _ = sql.Open("postgres", pgUrl)
	db.Ping()

	return db
}
