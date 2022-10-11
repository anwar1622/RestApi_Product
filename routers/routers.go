package routers

import (
	"RestApi_Product/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/product", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/product", controllers.GetAllProduct).Methods("GET")
	router.HandleFunc("/api/product/{idBarang}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/product/{idBarang}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/product/{idBarang}", controllers.DeleteProductById).Methods("DELETE")

	controllers.Connection()
	return router
}
