package routes

import (
	"api/controllers"
	"database/sql"
	"encoding/json"
	"net/http"
)

func CreateMagazine(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id := controllers.CreateMagazine(r, db)
		json.NewEncoder(rw).Encode(id)
	}
}
