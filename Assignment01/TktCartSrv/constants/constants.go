package constants

type Errorcodes = int

const GetAvailableTicketsURL string = "http://localhost:8050/GetAvailableTickets"

const ClientApikey string = "abcd@123456"
const TktStoreApikey string = "abcd@12345"

// Error codes
const (
	ErrorNoErrors              = 999
	ErrorCodeBadRequest        = 1001
	ErrorCodeTicktNotAvailable = 1002
	ErrorCodeAuthError         = 1003
)

// Error string
const (
	ErrorStringNoErrors               = "Success"
	ErrorStringCodeBadRequest         = "Bad Request"
	ErrorStringNotAvailableForBooking = "not available for booking"
	ErrorStringAuthError              = "Not authorized"
)
