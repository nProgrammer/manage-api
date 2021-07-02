package config

import (
	"api/utils"
	"database/sql"
	"fmt"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANT_URL"))
	utils.ErrorM(err)
	fmt.Println(os.Getenv("ELEPHANT_URL"))
	db, err = sql.Open("postgres", pgUrl)
	utils.ErrorM(err)
	db.Ping()

	return db
}
