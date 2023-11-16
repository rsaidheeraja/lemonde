// Go connection Sample Code:
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initDB()

	// Initialize router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/sales", insertCustomerAndSale).Methods("POST")

	// Start the server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initDB() {
	panic("unimplemented")
}
