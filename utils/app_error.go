package utils

import (
	"fmt"
	"net/http"
)

type AppError struct {
	ErrorCode    string
	ErrorMessage string
	ErrorType    int
}

func (a AppError) Error() string {
	return fmt.Sprintf("type: %d, code: %s, err: %s", a.ErrorType, a.ErrorCode, a.ErrorMessage)
}

// error classification
// shown to user as the error coming from them

func RequiredError() error {
	return AppError{
		ErrorCode:    "X01",
		ErrorMessage: "Input should never be empty",
		ErrorType:    http.StatusBadRequest,
	}
}

func UnauthorisedError() error {
	return AppError{
		ErrorCode:    "X04",
		ErrorMessage: "Unauthorised user",
		ErrorType:    http.StatusUnauthorized,
	}
}

func DataNotFoundError() error {
	return AppError{
		ErrorCode:    "X02",
		ErrorMessage: "Data not found!",
	}
}
