package consts

import "errors"

var (
	ErrNoContextValue = errors.New("resource not found in request context")
)

const (
	ErrStrRetrieveData = "Failed to retrieve data from database."
	ErrStrDecode       = "Failed to read request body."
	ErrStrDB           = "Failed to update database."
	ErrStrUnauthorized = "User is not authorized to access this resource."
	ErrStrResourceDNE  = "Requested resource does not exist."
	ErrUnexpected      = "An Unexpected error has occurred. Please try again later."
)
