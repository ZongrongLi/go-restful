package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrSQL              = &Errno{Code: 10003, Message: "sql error."}
	ErrParam            = &Errno{Code: 10004, Message: "param error."}
	ErrEmpty            = &Errno{Code: 10005, Message: "empty error."}

	// user errors
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
)
