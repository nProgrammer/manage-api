package routes

import (
	"api/controllers"
	"database/sql"
	"net/http"
)

func GetAllMagazines(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		controllers.GetMagazine(db, rw)
	}
}
