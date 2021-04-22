package ticketsrepo

import (
	"fmt"
	"sync"
	"ticketstore/constants"
	"ticketstore/models"
)

var availabletickets []models.Ticket
var reservedTickets []models.Ticket
var soldoutTickets []models.Ticket
var prebookedtickets []models.PreBookedTicket
var bookedtickets []models.BookedTicket
var prebookingIdGenerator = 1

var lock = sync.RWMutex{}

// Getters
func GetAvailableTickets() []models.Ticket {
	var tickets []models.Ticket

	lock.RLock()
	tickets = availabletickets
	lock.RUnlock()

	return tickets
}

func isTicketBooked(pTicketID int) bool {
	for i := 0; i < len(bookedtickets); i++ {
		if bookedtickets[i].TicketId == pTicketID {
			return true
		}
	}
	return false
}

func isAnyOfTicketBooked(pTicketsForPrebooking models.PreBookTicketList) bool {

	for i := 0; i < len(pTicketsForPrebooking.PreBookingTickets); i++ {
		if isTicketBooked(pTicketsForPrebooking.PreBookingTickets[i].TicketId) {
			return true
		}
	}
	return false
}

func isTicketPreBooked(pTicketID int) bool {
	for i := 0; i < len(prebookedtickets); i++ {
		if prebookedtickets[i].TicketId == pTicketID {
			return true
		}
	}
	return false
}

func isAnyOfTicketPreBooked(pTicketsForPrebooking models.PreBookTicketList) bool {

	for i := 0; i < len(pTicketsForPrebooking.PreBookingTickets); i++ {
		if isTicketPreBooked(pTicketsForPrebooking.PreBookingTickets[i].TicketId) {
			return true
		}
	}
	return false
}

func markTicketAsReserved(pTicketID int) {
	for i := 0; i < len(availabletickets); i++ {
		if availabletickets[i].TicketID == pTicketID {
			reservedTickets = append(reservedTickets, availabletickets[i])
			availabletickets = append(availabletickets[:i], availabletickets[i+1:]...)
			return
		}
	}
}

func PreBookTickets(pTicketsForPrebooking models.PreBookTicketList) (int, constants.Errorcodes, error) {
	var prebookid int
	lock.RLock()
	// Check if tickets are already booked
	if isAnyOfTicketBooked(pTicketsForPrebooking) {
		lock.RUnlock()
		return -1, constants.ErrorCodeTicktNotAvailable, fmt.Errorf(constants.ErrorStringNotAvailableForBooking)
	}

	// Check if tickets are in prebooking
	if isAnyOfTicketPreBooked(pTicketsForPrebooking) {
		lock.RUnlock()
		return -1, constants.ErrorCodeTicktNotAvailable, fmt.Errorf(constants.ErrorStringNotAvailableForBooking)
	}
	lock.RUnlock()

	// Tickets are available for booking
	lock.Lock()
	prebookid = prebookingIdGenerator
	for i := 0; i < len(pTicketsForPrebooking.PreBookingTickets); i++ {
		pTicketsForPrebooking.PreBookingTickets[i].PreBookingId = prebookid
		markTicketAsReserved(pTicketsForPrebooking.PreBookingTickets[i].TicketId)
		prebookedtickets = append(prebookedtickets, pTicketsForPrebooking.PreBookingTickets[i])
	}
	lock.Unlock()

	prebookingIdGenerator++

	return prebookid, constants.ErrorNoErrors, nil
}

// Helper functions
func GenerateTickets(pTheatreName string, pShowDateTime string, pTicketCost float64, pNumberofTickets int) {
	var ticket models.Ticket

	for i := 1; i < pNumberofTickets; i++ {
		ticket = models.Ticket{
			TicketID:     i,
			TheatreName:  pTheatreName,
			ShowDateTime: pShowDateTime,
			TicketCost:   pTicketCost,
		}

		availabletickets = append(availabletickets, ticket)
	}

}
