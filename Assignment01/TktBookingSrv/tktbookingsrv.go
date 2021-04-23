package main

import (
	"log"
	"net/http"
	"tktbookingsrv/requesthandlers"
)

func main() {
	startServing()
}

func startServing() {

	registerRouts()

	log.Println("Starting Server on port : 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println(err)
		return
	}
}

func registerRouts() {
	//http.HandleFunc("/GetTicketsFromCart", requesthandlers.GetTicketsFromCart)
	http.HandleFunc("/InitiateBookingForCartTickets", requesthandlers.InitiateBookingForCartTickets)
}
