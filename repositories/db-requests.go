package repositories

import (
	"api/models"
	"database/sql"
	"log"
)

func CreateWarehouseDB(db *sql.DB, newMag models.Warehouse) int {
	_ = db.QueryRow("insert into warehouses(holder, isReserved, price, dateTillReserves, compID) values ($1 ,$2, $3, $4, $5)",
		newMag.Holder, newMag.IsReserved, newMag.Price, newMag.DateTillReserved, newMag.CompanyID).Scan(&newMag.ID)
	return newMag.ID
}

func CreateClientDB(db *sql.DB, client models.Client) int {
	log.Println(client)
	_ = db.QueryRow("insert into clients(name, login, phone, email) values ($1, $2, $3, $4)",
		client.Name, client.Login, client.Phone, client.Email).Scan(&client.ID)
	return client.ID
}

func GetWarehousesDB(db *sql.DB, warehouse models.Warehouse, warehouses []models.Warehouse) []models.Warehouse {
	rows, _ := db.Query("select * from warehouses;")
	for rows.Next() {
		_ = rows.Scan(&warehouse.ID, &warehouse.Holder, &warehouse.IsReserved, &warehouse.Price, &warehouse.DateTillReserved, &warehouse.CompanyID)
		warehouses = append(warehouses, warehouse)
	}
	return warehouses
}

func GetClientsDB(db *sql.DB, client models.Client, clients []models.Client) []models.Client {
	rows, _ := db.Query("select * from clients;")
	for rows.Next() {
		_ = rows.Scan(&client.ID, &client.Name, &client.Login, &client.Phone, &client.Email)
		clients = append(clients, client)
	}
	return clients
}

func GetWarehouseDB(db *sql.DB, warehouse models.Warehouse, warehouses []models.Warehouse, id string) []models.Warehouse {
	rows, _ := db.Query("select * from warehouses where compID=$1;", id)
	for rows.Next() {
		_ = rows.Scan(&warehouse.ID, &warehouse.Holder, &warehouse.IsReserved, &warehouse.Price, &warehouse.DateTillReserved, &warehouse.CompanyID)
		warehouses = append(warehouses, warehouse)
	}
	return warehouses
}

func GetWarehouseReservedByDB(db *sql.DB, warehouse models.Warehouse, warehouses []models.Warehouse, name string) []models.Warehouse {
	rows, _ := db.Query("select * from warehouses where holder=$1;", name)
	for rows.Next() {
		_ = rows.Scan(&warehouse.ID, &warehouse.Holder, &warehouse.IsReserved, &warehouse.Price, &warehouse.DateTillReserved, &warehouse.CompanyID)
		warehouses = append(warehouses, warehouse)
	}
	return warehouses
}

func GetReservedWarehousesDB(db *sql.DB, warehouse models.Warehouse, warehouses []models.Warehouse) []models.Warehouse {
	rows, _ := db.Query("select * from warehouses where holder != '';")
	for rows.Next() {
		_ = rows.Scan(&warehouse.ID, &warehouse.Holder, &warehouse.IsReserved, &warehouse.Price, &warehouse.DateTillReserved, &warehouse.CompanyID)
		warehouses = append(warehouses, warehouse)
	}
	return warehouses
}

func ReserveWarehouseDB(db *sql.DB, warehouse models.Warehouse) int64 {
	result, _ := db.Exec("update warehouses set holder=$1, isReserved=$2, dateTillReserves=$3 where compID=$4",
		&warehouse.Holder, &warehouse.IsReserved, &warehouse.DateTillReserved, &warehouse.CompanyID)

	rowsUpd, _ := result.RowsAffected()

	return rowsUpd
}

func UpdateClientDB(db *sql.DB, client models.Client) int64 {
	result, _ := db.Exec("update clients set name=$1, login=$2, phone=$3, email=$4 where login=$5",
		&client.Name, &client.Login, &client.Phone, &client.Email, &client.Login)

	rowsUpd, _ := result.RowsAffected()

	return rowsUpd
}

func UpdateWarehousePriceDB(db *sql.DB, warehouse models.Warehouse) int64 {
	result, _ := db.Exec("update warehouses set price=$1 where compID=$2",
		&warehouse.Price, &warehouse.CompanyID)

	rowsUpd, _ := result.RowsAffected()

	return rowsUpd
}

func DeleteWarehouseDB(db *sql.DB, compID string) {
	result, _ := db.Exec("delete from warehouses where compID=$1", compID)

	result.RowsAffected()
}

func FindWarehousesDB(db *sql.DB, warehouseToFind models.Warehouse, warehouses []models.Warehouse) []models.Warehouse {
	var warehouse models.Warehouse
	rows, _ := db.Query("select * from warehouses where price=$1;", warehouseToFind.Price)
	for rows.Next() {
		_ = rows.Scan(&warehouse.ID, &warehouse.Holder, &warehouse.IsReserved, &warehouse.Price, &warehouse.DateTillReserved, &warehouse.CompanyID)
		warehouses = append(warehouses, warehouse)
	}
	return warehouses
}

func FindClientDB(login string, db *sql.DB, clients []models.Client) []models.Client {
	var client models.Client
	rows, _ := db.Query("select * from clients where login=$1;", login)
	for rows.Next() {
		_ = rows.Scan(&client.ID, &client.Name, &client.Login, &client.Phone, &client.Email)
		clients = append(clients, client)
	}
	return clients
}

func DeleteClientDB(db *sql.DB, login string) {
	result, _ := db.Exec("delete from clients where login=$1", login)

	result.RowsAffected()
}
