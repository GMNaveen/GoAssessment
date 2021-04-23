package requesthandlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"tktcartsrv/cartsrv"
	"tktcartsrv/constants"
	"tktcartsrv/models"
)

var errorrespponse models.ErrorResponse

func GetAllTicketsFromUserCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s GetAllTicketsFromUserCart\n", r.Method)

	cartTickets := cartsrv.GetAllCartTickets()
	jsondata, _ := json.Marshal(cartTickets)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)

	log.Println("All Cart Tickets returned.")
}

func GetTicketsFromUserCart(w http.ResponseWriter, r *http.Request) {
	var cartId int
	if r.Method != http.MethodGet {
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s GetTicketsFromUserCart\n", r.Method)

	m, _ := url.ParseQuery(r.URL.RawQuery)

	//log.Println(m["cartid"][0])
	cartId, _ = strconv.Atoi(m["cartid"][0])

	cartTickets := cartsrv.GetCartTickets(cartId)
	jsondata, _ := json.Marshal(cartTickets)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)

	log.Println("Cart Tickets returned.")
}

func AddTicketsToUserCart(w http.ResponseWriter, r *http.Request) {
	var ticketsForCarting models.TicketsListForCart
	var cartId int
	var errorcode constants.Errorcodes

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s AddTicketsToUserCart\n", r.Method)

	payload, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(payload, &ticketsForCarting)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	cartId, errorcode, err = cartsrv.AddTicketsToUserCart(ticketsForCarting)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = errorcode
		errorrespponse.ErrorMsg = err.Error()
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	// On success, store cart id and return
	for i := 0; i < len(ticketsForCarting.Tickets); i++ {
		ticketsForCarting.Tickets[i].CartID = cartId
	}

	jsondata, _ := json.Marshal(ticketsForCarting)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)

}

func ClearCartTicketsForUser(w http.ResponseWriter, r *http.Request) {
	var ticketsForUnCarting models.TicketsListForCart
	var errorcode constants.Errorcodes

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s ClearCartTicketsForUser\n", r.Method)

	payload, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(payload, &ticketsForUnCarting)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	errorcode, err = cartsrv.ClearCartTicketsForUser(ticketsForUnCarting)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = errorcode
		errorrespponse.ErrorMsg = err.Error()
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cart Tickets Cleared"))

}
