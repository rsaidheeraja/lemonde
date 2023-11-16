package main

import (
	"REST_API/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func insertCustomerAndSale(w http.ResponseWriter, r *http.Request) {
	var customerData models.Customer
	err := json.NewDecoder(r.Body).Decode(&customerData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the customer already exists in the database
	var customerID int
	err = db.QueryRow("SELECT Id FROM Customer WHERE Email = ?", customerData.Email).Scan(&customerID)

	// If the customer does not exist, insert a new record
	if err == sql.ErrNoRows {
		result, err := db.Exec(
			"INSERT INTO Customer (FirstName, LastName, ContactNumber, Email) VALUES (?, ?, ?, ?)",
			customerData.FirstName, customerData.LastName, customerData.ContactNumber, customerData.Email,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Retrieve the auto-incremented ID of the newly inserted customer
		customerID, err = result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Now, insert data into the Sales table
	var saleData models.Sales
	err = json.NewDecoder(r.Body).Decode(&saleData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"INSERT INTO Sales (CustomerID, OriginLocationID, DestinationLocationID, ClassOfServiceID) VALUES (?, ?, ?, ?)",
		customerID, saleData.OriginLocationID, saleData.DestinationLocation, saleData.ClassOfServiceID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
