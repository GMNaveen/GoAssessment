package constants

type Errorcodes = int

const GetTicketsFromUserCartURL string = "http://localhost:8070/GetTicketsFromUserCart"
const PreBookTicketsForTicketStoreSrvURL string = "http://localhost:8050/PreBookTickets"
const ConfirmBookedTicketsForTicketStoreSrvURL string = "http://localhost:8050/ConfirmPreBookTickets"
const ClearPreBookTicketsForTicketStoreSrvURL string = "http://localhost:8050/RestorePreBookedTickets"

const TktStoreApikey string = "abcd@12345"
const TktCartSrvApikey string = "abcd@123456"

// Error codes
const (
	ErrorNoErrors                      = 999
	ErrorCodeBadRequest                = 1001
	ErrorCodeTicktNotAvailable         = 1002
	ErrorCodeAuthError                 = 1003
	ErrorCodeTicketServerNotResponsing = 1004
)

// Error string
const (
	ErrorStringNoErrors                  = "Success"
	ErrorStringCodeBadRequest            = "Bad Request"
	ErrorStringNotAvailableForBooking    = "Not available for booking"
	ErrorStringAuthError                 = "Not authorized"
	ErrorStringTicketServerNotResponsing = "Ticket Server not responding"
)
