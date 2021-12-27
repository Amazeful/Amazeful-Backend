package consts

import "errors"

var (
	ErrNoContextValue = errors.New("resource not found in request context")
)

const (
	ErrStrRetrieveData = "Failed to retrieve data from database."
	ErrStrDecode       = "Failed to decode request data."
	ErrStrInsert       = "Failed to insert data into database."
	ErrStrUnauthorized = "User is not authorized to access this resource."
	ErrStrResourceDNE  = "Requested resource does not exist."
)
