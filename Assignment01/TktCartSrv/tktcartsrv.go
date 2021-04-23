package main

import (
	"log"
	"net/http"
	"tktcartsrv/cartsrv"
	"tktcartsrv/requesthandlers"
)

func main() {
	initializeCartTickets()
	startServing()
}

func startServing() {

	registerRouts()

	log.Println("Starting Server on port : 8070")
	err := http.ListenAndServe(":8070", nil)

	if err != nil {
		log.Println(err)
		return
	}
}

func registerRouts() {
	http.HandleFunc("/GetAllTicketsFromUserCart", requesthandlers.GetAllTicketsFromUserCart)
	http.HandleFunc("/GetTicketsFromUserCart", requesthandlers.GetTicketsFromUserCart)
	http.HandleFunc("/AddTicketsToUserCart", requesthandlers.AddTicketsToUserCart)
	http.HandleFunc("/ClearCartTicketsForUserCart", requesthandlers.ClearCartTicketsForUser)
}

// Genetare Tickets for a show
func initializeCartTickets() {
	cartsrv.ClearCartTickets()
}
