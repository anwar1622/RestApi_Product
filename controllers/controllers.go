package controllers

import (
	"RestApi_Product/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

var db *gorm.DB

func Connection() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("Connected successfully to the database")
	//db.Exec("CREATE DATABASE product")
	db.Exec("USE  product")
	db.AutoMigrate(&models.Barang{}, &models.Item{})
}

func CreateProduct(write http.ResponseWriter, r *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	var product models.Barang
	json.NewDecoder(r.Body).Decode(&product)
	db.Create(&product)
	json.NewEncoder(write).Encode(product)
}

func GetAllProduct(write http.ResponseWriter, r *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	var product []models.Barang
	db.Preload("Items").Find(&product)
	json.NewEncoder(write).Encode(product)
}

func GetProductById(write http.ResponseWriter, r *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputIdBarang := params["idBarang"]
	var product models.Barang
	db.Preload("Items").First(&product, inputIdBarang)
	json.NewEncoder(write).Encode(product)
}

func UpdateProduct(write http.ResponseWriter, r *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	var updatedProduct models.Barang
	json.NewDecoder(r.Body).Decode(updatedProduct)
	db.Save(&updatedProduct)
	json.NewEncoder(write).Encode(updatedProduct)
}

func DeleteProductById(write http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IdBarang := params["idBarang"]
	id64, _ := strconv.ParseUint(IdBarang, 10, 64)
	idToDelete := uint(id64)

	db.Where("id_barang = ?", idToDelete).Delete(&models.Item{})
	db.Where("id_barang = ?", idToDelete).Delete(&models.Barang{})
	write.WriteHeader(http.StatusNoContent)
}
