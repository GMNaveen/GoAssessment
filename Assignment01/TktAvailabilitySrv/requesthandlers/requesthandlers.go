package requesthandlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"tktavailabilitysrv/constants"
	"tktavailabilitysrv/models"
	"tktavailabilitysrv/ticketsrv"
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

	ticketsjsondata, err := ticketsrv.GetAvailableTickets()
	jsondata, _ := ioutil.ReadAll(ticketsjsondata.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsondata)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)

	log.Println("Available Tickets returned.")
}
