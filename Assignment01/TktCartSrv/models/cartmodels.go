package models

import "tktcartsrv/constants"

type ErrorResponse struct {
	ErrorCode constants.Errorcodes `json:"errorcode"`
	ErrorMsg  string               `json:"errormessage"`
}

type TicketsListForCart struct {
	Tickets []TicketsCart `json:"ticketstocart"`
}

type TicketsCart struct {
	CartID           int    `json:"cartid"`
	TicketID         int    `json:"ticketid"`
	CustomerMobileNo string `json:"customermobile"`
}
