package requesthandlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"ticketstore/constants"
	"ticketstore/models"
	"ticketstore/ticketsrepo"
)

var errorrespponse models.ErrorResponse

func GetAvailableTickets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s GetAvailableTickets\n", r.Method)

	tickets := ticketsrepo.GetAvailableTickets()
	jsondata, _ := json.Marshal(tickets)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)

	log.Println("Available Tickets returned.")
}

func PreBookTickets(w http.ResponseWriter, r *http.Request) {
	var ticketsForPrebooking models.PreBookTicketList
	var preBookingId int
	var errorcode constants.Errorcodes

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s PreBookTickets\n", r.Method)

	payload, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(payload, &ticketsForPrebooking)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	preBookingId, errorcode, err = ticketsrepo.PreBookTickets(ticketsForPrebooking)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = errorcode
		errorrespponse.ErrorMsg = err.Error()
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	// On success, store prebooking id and return
	for i := 0; i < len(ticketsForPrebooking.PreBookingTickets); i++ {
		ticketsForPrebooking.PreBookingTickets[i].PreBookingId = preBookingId
	}

	jsondata, _ := json.Marshal(ticketsForPrebooking)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)
}
