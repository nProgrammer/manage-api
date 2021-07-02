package routes

import (
	"api/controllers"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteMagazine(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["compID"]
		controllers.DeleteMagazineCT(db, rw, id)
	}
}
