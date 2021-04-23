package models

import "tktbookingsrv/constants"

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
