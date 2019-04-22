package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrEmpty            = &Errno{Code: 10005, Message: "empty error."}

	// user errors
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrBlack             = &Errno{Code: 20103, Message: ""}
	ErrTokenInvalid      = &Errno{Code: 20104, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20105, Message: "The password was incorrect."}
	ErrParam             = &Errno{Code: 20106, Message: "param error."}
	ErrSQL               = &Errno{Code: 20107, Message: "sql error."}
)
