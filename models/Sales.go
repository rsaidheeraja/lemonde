package models

type Sales struct {
	CustomerID          int `json:"customerid"`
	OriginLocationID    int `json:"originlocationid"`
	DestinationLocation int `json:"destinationlocationid"`
	ClassOfServiceID    int `json:"classofserviceid"`
}
