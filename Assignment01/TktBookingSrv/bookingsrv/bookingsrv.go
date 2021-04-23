package bookingsrv

import (
	"net/http"
	"strconv"
	"tktbookingsrv/constants"
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
