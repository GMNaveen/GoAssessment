package models

import "tktavailabilitysrv/constants"

type ErrorResponse struct {
	ErrorCode constants.Errorcodes `json:"errorcode"`
	ErrorMsg  string               `json:"errormessage"`
}

type Ticket struct {
	TicketID     int     `json:"ticketid"`
	TheatreName  string  `json:"theatrename"`
	ShowDateTime string  `json:"showdatetime"`
	TicketCost   float64 `json:"cost"`
}
