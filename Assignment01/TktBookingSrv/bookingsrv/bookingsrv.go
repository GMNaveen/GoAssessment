package bookingsrv

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"tktbookingsrv/constants"
	"tktbookingsrv/models"
	"tktbookingsrv/paymentgateway"
)

func GetTicketsFromCart(pCartId int) (*http.Response, error) {

	request, _ := http.NewRequest("GET", constants.GetTicketsFromUserCartURL, nil)

	q := request.URL.Query()
	q.Add("cartid", strconv.Itoa(pCartId))
	request.URL.RawQuery = q.Encode()

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", constants.TktCartSrvApikey)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

func ConfirmBookTicketsTktSrv(pTicketsForBookingConfirmation models.BookedTicketList) (*http.Response, error) {

	jsondata, _ := json.Marshal(pTicketsForBookingConfirmation)
	requestBody := strings.NewReader(string(jsondata))

	request, _ := http.NewRequest("POST", constants.ConfirmBookedTicketsForTicketStoreSrvURL, requestBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", constants.TktStoreApikey)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

func PreBookTicketsTktSrv(pTicketsForPrebooking models.PreBookTicketList) (*http.Response, error) {

	jsondata, _ := json.Marshal(pTicketsForPrebooking)
	requestBody := strings.NewReader(string(jsondata))

	request, _ := http.NewRequest("POST", constants.PreBookTicketsForTicketStoreSrvURL, requestBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", constants.TktStoreApikey)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

func ClearPreBookTicketsTktSrv(pTicketsForPrebooking models.PreBookTicketList) (*http.Response, error) {

	jsondata, _ := json.Marshal(pTicketsForPrebooking)
	requestBody := strings.NewReader(string(jsondata))

	request, _ := http.NewRequest("POST", constants.ClearPreBookTicketsForTicketStoreSrvURL, requestBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", constants.TktStoreApikey)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

func GetTicketsFromCartIntoModel(pCartId int) (models.TicketsListForCart, error) {
	var ticketsFromCartForBooking models.TicketsListForCart

	response, err := GetTicketsFromCart(pCartId)

	if err != nil {
		return models.TicketsListForCart{}, err
	}

	defer response.Body.Close()

	payload, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(payload, &ticketsFromCartForBooking)

	if err != nil {
		return models.TicketsListForCart{}, err
	}

	return ticketsFromCartForBooking, nil
}

func InitiateBookingForCartTickets(pCartId int) (*http.Response, error) {
	var ticketsFromCartForBooking models.TicketsListForCart
	var ticketsForPrebooking models.PreBookTicketList
	var ticketsForBookingConfirmation models.BookedTicketList
	var errorrespponse models.ErrorResponse
	var err error

	ticketsFromCartForBooking, err = GetTicketsFromCartIntoModel(pCartId)

	if err != nil {
		return nil, err
	}

	// Tickets received from cart.

	// Attempt to pre-book the tickets
	for i := 0; i < len(ticketsFromCartForBooking.Tickets); i++ {
		var prebookticket models.PreBookedTicket

		prebookticket.PreBookingId = -1
		prebookticket.TicketId = ticketsFromCartForBooking.Tickets[i].TicketID
		prebookticket.CustomerMobileNo = ticketsFromCartForBooking.Tickets[i].CustomerMobileNo

		ticketsForPrebooking.PreBookingTickets = append(ticketsForPrebooking.PreBookingTickets, prebookticket)
	}

	res, err := PreBookTicketsTktSrv(ticketsForPrebooking)
	if err != nil {
		return res, err
	}

	if res.StatusCode != http.StatusOK {
		return res, err
	}

	defer res.Body.Close()

	// Check if body received is of type ErrorResponse. If so send back the error response as it is
	payload, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(payload, &errorrespponse)
	if err == nil {
		if errorrespponse.ErrorCode != 0 {
			// This means it is ErrorResponse.
			if errorrespponse.ErrorCode == constants.ErrorCodeTicktNotAvailable {
				ClearPreBookTicketsTktSrv(ticketsForPrebooking)
			}
			return res, err
		}
	}

	// Prebooking has succeded.
	payRef, _ := paymentgateway.MakePayment(176 * 2)

	for i := 0; i < len(ticketsForPrebooking.PreBookingTickets); i++ {
		var bookedticket models.BookedTicket

		bookedticket.BookingId = -1
		bookedticket.TicketId = ticketsForPrebooking.PreBookingTickets[i].TicketId
		bookedticket.CustomerMobileNo = ticketsForPrebooking.PreBookingTickets[i].CustomerMobileNo
		bookedticket.PaymentReference = payRef

		ticketsForBookingConfirmation.BookedTickets = append(ticketsForBookingConfirmation.BookedTickets, bookedticket)
	}

	res, err = ConfirmBookTicketsTktSrv(ticketsForBookingConfirmation)
	if err == nil {
		ClearPreBookTicketsTktSrv(ticketsForPrebooking)
	}

	return res, err
}
