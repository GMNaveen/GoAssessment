package constants

type Errorcodes = int

const GetAvailableTicketsURL string = "http://localhost:8050/GetAvailableTickets"
const TktStoreApikey string = "abcd@12345"

// Error codes
const (
	ErrorNoErrors                      = 999
	ErrorCodeBadRequest                = 1001
	ErrorCodeTicktNotAvailable         = 1002
	ErrorCodeTicketServerNotResponsing = 1003
)

// Error string
const (
	ErrorStringNoErrors                  = "Success"
	ErrorStringCodeBadRequest            = "Bad Request"
	ErrorStringNotAvailableForBooking    = "Not available for booking"
	ErrorStringTicketServerNotResponsing = "Ticket Server not responding"
)
