package main

type Page struct {
	AppContext   *AppContext
	Guest        *Guest
	HotelList    []Hotel
	SelectStatus map[string]string
}
