package controllers

import (
	"api/models"
	"api/repositories"
	"api/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

func CreateWarehouse(r *http.Request, db *sql.DB, rw http.ResponseWriter) {
	var newMag models.Warehouse
	json.NewDecoder(r.Body).Decode(&newMag)
	_ = repositories.CreateWarehouseDB(db, newMag)
	json.NewEncoder(rw).Encode(newMag)
}

func GetWarehouse(db *sql.DB, rw http.ResponseWriter) {
	var warehouses []models.Warehouse
	var warehouse models.Warehouse
	warehouses = repositories.GetWarehousesDB(db, warehouse, warehouses)
	json.NewEncoder(rw).Encode(warehouses)
}

func GetWarehouseCT(db *sql.DB, rw http.ResponseWriter, id string) {
	var warehouses []models.Warehouse
	var warehouse models.Warehouse
	warehouses = repositories.GetWarehouseDB(db, warehouse, warehouses, id)
	json.NewEncoder(rw).Encode(warehouses)
}

func GetWarehouseReservedByCT(db *sql.DB, rw http.ResponseWriter, name string) {
	var warehouses []models.Warehouse
	var warehouse models.Warehouse
	warehouses = repositories.GetWarehouseReservedByDB(db, warehouse, warehouses, name)
	json.NewEncoder(rw).Encode(warehouses)
}

func ReserveWarehouseCT(db *sql.DB, warehouse models.Warehouse) int64 {
	rowsUpd := repositories.ReserveWarehouseDB(db, warehouse)
	return rowsUpd
}

func UpdateWarehousePriceCT(db *sql.DB, warehouse models.Warehouse) int64 {
	rowsUpd := repositories.UpdateWarehousePriceDB(db, warehouse)
	return rowsUpd
}

func GetReservedWarehouses(db *sql.DB, rw http.ResponseWriter) {
	var warehouses []models.Warehouse
	var warehouse models.Warehouse
	warehouses = repositories.GetReservedWarehousesDB(db, warehouse, warehouses)
	json.NewEncoder(rw).Encode(warehouses)
}

func DeleteWarehouseCT(db *sql.DB, rw http.ResponseWriter, compID string) {
	text := "Deleted " + compID
	repositories.DeleteWarehouseDB(db, compID)
	utils.JsonResponse(text, true, rw)
}

func FindWarehouseCT(db *sql.DB, warehouseToFind models.Warehouse, rw http.ResponseWriter) {
	var warehouses []models.Warehouse
	warehouses = repositories.FindWarehousesDB(db, warehouseToFind, warehouses)
	json.NewEncoder(rw).Encode(warehouses)
}
