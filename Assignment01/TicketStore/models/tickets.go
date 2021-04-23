package models

import "ticketstore/constants"

type ErrorResponse struct {
	ErrorCode constants.Errorcodes `json:"errorcode"`
	ErrorMsg  string               `json:"errormessage"`
}

type Ticket struct {
	TicketID     int     `json:"ticketid"` // This is the unique key in our system
	TheatreName  string  `json:"theatrename"`
	ShowDateTime string  `json:"showdatetime"`
	TicketCost   float64 `json:"cost"`
}

type PreBookTicketList struct {
	PreBookingTickets []PreBookedTicket `json:"prebookticketlist"`
}

type PreBookedTicket struct {
	PreBookingId     int    `json:"preboookingid"`
	TicketId         int    `json:"ticketid"`
	CustomerMobileNo string `json:"customermobile"`
}

type BookedTicketList struct {
	BookedTickets []BookedTicket `json:"bookedtickets"`
}

type BookedTicket struct {
	BookingId        int    `json:"boookingid"`
	TicketId         int    `json:"ticketid"`
	CustomerMobileNo string `json:"customermobile"`
	PaymentReference string `json:"paymentreference"`
}
