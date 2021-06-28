package repositories

import (
	"api/models"
	"database/sql"
)

func CreateMagazineDB(db *sql.DB, newMag models.Magazine) int {
	_ = db.QueryRow("insert into magazines(holder, isReserved, price, dateTillReserves, compID) values ($1 ,$2, $3, $4, $5)",
		newMag.Holder, newMag.IsReserved, newMag.Price, newMag.DateTillReserved, newMag.CompanyID).Scan(&newMag.ID)
	return newMag.ID
}

func GetMagazinesDB(db *sql.DB, magazine models.Magazine, magazines []models.Magazine) []models.Magazine {
	rows, _ := db.Query("select * from magazines;")
	for rows.Next() {
		_ = rows.Scan(&magazine.ID, &magazine.Holder, &magazine.IsReserved, &magazine.Price, &magazine.DateTillReserved, &magazine.CompanyID)
		magazines = append(magazines, magazine)
	}
	return magazines
}
