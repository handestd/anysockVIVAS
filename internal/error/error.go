package error

import (
	"encoding/json"
	"errors"
)

type Error struct {
	Code    int
	Message string
	Error   error
}

func NewError(code int, message string) *Error {
	return &Error{Code: code, Message: message, Error: errors.New(message)}
}
func JsonError(code int, message string) string {
	json, err := json.Marshal(&Error{Code: code, Message: message})
	if err != nil {
		return err.Error()
	} else {
		return string(json)
	}
}

func ExposeError(err error, es ...*Error) *Error {
	for _, e := range es {
		if errors.Is(err, e.Error) {
			return e
		}
	}
	return Unknown
}

var (
	Unknown           = NewError(1, "Unknown error")
	UserExist         = NewError(2, "User Exist")
	IncorrectPassword = NewError(3, "Incorrect Password")
	NoMatchPassword   = NewError(4, "No Match Password")
	Unauthorized      = NewError(5, "Unauthorized")
	EmptyField        = NewError(6, "Empty Field")
	InValidUserID     = NewError(7, "Invalid User ID")
	RecordNotFound    = NewError(8, "Record Not Found")
	EmptyTable        = NewError(9, "Empty Table")
)
var (
	LoginRequired = JsonError(401, "Login Required")
)
