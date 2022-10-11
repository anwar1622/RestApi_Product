package main

import (
	"RestApi_Product/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := routers.Router()
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
