package cartsrv

import (
	"sync"
	"tktcartsrv/constants"
	"tktcartsrv/models"
)

var cartTicketList []models.TicketsCart
var cartIdGenerator = 1

var lock = sync.RWMutex{}

func GetAllCartTickets() []models.TicketsCart {
	var cartTickets []models.TicketsCart

	lock.RLock()
	cartTickets = cartTicketList
	lock.RUnlock()

	return cartTickets
}

func GetCartTickets(pCartId int) models.TicketsListForCart {
	var cartTickets models.TicketsListForCart

	cartTickets.Tickets = make([]models.TicketsCart, 0)
	lock.RLock()
	for i := 0; i < len(cartTicketList); i++ {
		if cartTicketList[i].CartID == pCartId {
			cartTickets.Tickets = append(cartTickets.Tickets, cartTicketList[i])
		}
	}
	lock.RUnlock()

	return cartTickets
}

func AddTicketsToUserCart(pTicketsForCarting models.TicketsListForCart) (int, constants.Errorcodes, error) {
	var cartId int

	// Note : Currently not validating if tickets are already in cart

	// Add Tickets to cart
	lock.Lock()
	cartId = cartIdGenerator
	for i := 0; i < len(pTicketsForCarting.Tickets); i++ {
		pTicketsForCarting.Tickets[i].CartID = cartId
		cartTicketList = append(cartTicketList, pTicketsForCarting.Tickets[i])
	}

	lock.Unlock()

	cartIdGenerator++

	return cartId, constants.ErrorNoErrors, nil
}

func clearTicketFromCart(pTicketID int) {
	for i := 0; i < len(cartTicketList); i++ {
		if cartTicketList[i].TicketID == pTicketID {
			cartTicketList = append(cartTicketList[:i], cartTicketList[i+1:]...)
			break
		}
	}
}

func ClearCartTicketsForUser(pTicketsForUnCarting models.TicketsListForCart) (constants.Errorcodes, error) {

	lock.Lock()

	for i := 0; i < len(pTicketsForUnCarting.Tickets); i++ {
		clearTicketFromCart(pTicketsForUnCarting.Tickets[i].TicketID)
	}

	lock.Unlock()

	return constants.ErrorNoErrors, nil
}

// Helper functions
func ClearCartTickets() {
	cartTicketList = make([]models.TicketsCart, 0)
}
