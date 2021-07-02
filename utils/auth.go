package utils

import (
	"api/models"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"net/http"
)

func AuthorizeFunc(db *sql.DB) []models.Authorize {
	var auth models.Authorize
	var auths []models.Authorize

	rows, _ := db.Query("select * from authorizedata;")
	for rows.Next() {
		_ = rows.Scan(&auth.ID, &auth.AuthCL, &auth.AuthPV, &auth.MainAuth)
		auths = append(auths, auth)
	}
	return auths
}

func AuthorizeMethod(r *http.Request, authDB []models.Authorize) string {
	au := r.Header.Get("Auth-Client")
	outau := au + authDB[0].AuthPV
	mainau := sha256.Sum256([]byte(outau))
	mainauS := hex.EncodeToString(mainau[:])

	return mainauS
}
