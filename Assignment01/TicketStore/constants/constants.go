package constants

const TicketsCount int = 100

type Errorcodes = int

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
