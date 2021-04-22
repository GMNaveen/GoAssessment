package constants

type Errorcodes = int

const GetAvailableTicketsURL string = "http://localhost:8070/GetAvailableTickets"
const Apikey string = "abcd@12345"

// Error codes
const (
	ErrorNoErrors              = 999
	ErrorCodeBadRequest        = 1001
	ErrorCodeTicktNotAvailable = 1002
)

// Error string
const (
	ErrorStringNoErrors               = "Success"
	ErrorStringCodeBadRequest         = "Bad Request"
	ErrorStringNotAvailableForBooking = "not available for booking"
)
