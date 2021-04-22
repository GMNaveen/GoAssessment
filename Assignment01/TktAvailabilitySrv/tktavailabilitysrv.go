package main

import (
	"log"
	"net/http"
	"tktavailabilitysrv/requesthandlers"
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
	http.HandleFunc("/GetAvailableTickets", requesthandlers.GetAvailableTickets)
}
