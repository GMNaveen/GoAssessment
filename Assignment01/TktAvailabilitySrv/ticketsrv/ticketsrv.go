package ticketsrv

import (
	"net/http"
	"tktavailabilitysrv/constants"
)

func GetAvailableTickets() (*http.Response, error) {

	request, _ := http.NewRequest("GET", constants.GetAvailableTicketsURL, nil)

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", constants.TktStoreApikey)

	client := &http.Client{}

	response, err := client.Do(request)

	return response, err
}
