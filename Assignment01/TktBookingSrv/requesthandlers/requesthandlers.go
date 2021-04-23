package requesthandlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"tktbookingsrv/bookingsrv"
	"tktbookingsrv/constants"
	"tktbookingsrv/models"
)

var errorrespponse models.ErrorResponse

func GetTicketsFromCart(w http.ResponseWriter, r *http.Request) {
	var cartId int
	if r.Method != http.MethodGet {
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s GetTicketsFromCart\n", r.Method)

	m, _ := url.ParseQuery(r.URL.RawQuery)
	cartId, _ = strconv.Atoi(m["cartid"][0])

	ticketsjsondata, err := bookingsrv.GetTicketsFromCart(cartId)

	if (err != nil) && (ticketsjsondata == nil) {
		errorrespponse.ErrorCode = constants.ErrorCodeTicketServerNotResponsing
		errorrespponse.ErrorMsg = constants.ErrorStringTicketServerNotResponsing
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	// We will be closing the response after processing received data
	defer ticketsjsondata.Body.Close()

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
