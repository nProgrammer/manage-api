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

func GetMagazineDB(db *sql.DB, magazine models.Magazine, magazines []models.Magazine, id string) []models.Magazine {
	rows, _ := db.Query("select * from magazines where compID=$1;", id)
	for rows.Next() {
		_ = rows.Scan(&magazine.ID, &magazine.Holder, &magazine.IsReserved, &magazine.Price, &magazine.DateTillReserved, &magazine.CompanyID)
		magazines = append(magazines, magazine)
	}
	return magazines
}

func GetMagazineReservedByDB(db *sql.DB, magazine models.Magazine, magazines []models.Magazine, name string) []models.Magazine {
	rows, _ := db.Query("select * from magazines where holder=$1;", name)
	for rows.Next() {
		_ = rows.Scan(&magazine.ID, &magazine.Holder, &magazine.IsReserved, &magazine.Price, &magazine.DateTillReserved, &magazine.CompanyID)
		magazines = append(magazines, magazine)
	}
	return magazines
}

func GetReservedMagazinesDB(db *sql.DB, magazine models.Magazine, magazines []models.Magazine) []models.Magazine {
	rows, _ := db.Query("select * from magazines where holder != '';")
	for rows.Next() {
		_ = rows.Scan(&magazine.ID, &magazine.Holder, &magazine.IsReserved, &magazine.Price, &magazine.DateTillReserved, &magazine.CompanyID)
		magazines = append(magazines, magazine)
	}
	return magazines
}

func ReserveMagazineDB(db *sql.DB, magazine models.Magazine) int64 {
	result, _ := db.Exec("update magazines set holder=$1, isReserved=$2, dateTillReserves=$3 where compID=$4",
		&magazine.Holder, &magazine.IsReserved, &magazine.DateTillReserved, &magazine.CompanyID)

	rowsUpd, _ := result.RowsAffected()

	return rowsUpd
}

func DeleteMagazineDB(db *sql.DB, compID string) {
	result, _ := db.Exec("delete from magazines where compID=$1", compID)

	result.RowsAffected()
}

func FindMagazinesDB(db *sql.DB, magazineToFind models.Magazine, magazines []models.Magazine) []models.Magazine {
	var magazine models.Magazine
	rows, _ := db.Query("select * from magazines where price=$1;", magazineToFind.Price)
	for rows.Next() {
		_ = rows.Scan(&magazine.ID, &magazine.Holder, &magazine.IsReserved, &magazine.Price, &magazine.DateTillReserved, &magazine.CompanyID)
		magazines = append(magazines, magazine)
	}
	return magazines
}
