package response

import (
	"errors"
	"net/http"

	"github.com/itsapep/golang-api-with-gin/utils"
)

type Status struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type Response struct {
	Status
	Data interface{} `json:"data,omitempty"`
}

func NewSuccessMessage(data interface{}) (httpStatusCode int, APIresponse Response) {
	status := Status{
		ResponseCode:    SuccessCode,
		ResponseMessage: SuccessMessage,
	}
	httpStatusCode = http.StatusOK
	APIresponse = Response{
		Status: status,
		Data:   data,
	}
	return
}

func NewErrorMessage(err error) (httpStatusCode int, APIresponse Response) {
	var userError utils.AppError
	var status Status
	if errors.As(err, &userError) {
		status = Status{
			ResponseCode:    userError.ErrorCode,
			ResponseMessage: userError.ErrorMessage,
		}
	} else {
		status = Status{
			ResponseCode:    DefaultErrorCode,
			ResponseMessage: DefaultErrorMessage,
		}
		httpStatusCode = http.StatusInternalServerError
	}
	APIresponse = Response{
		Status: status,
		Data:   nil,
	}
	return
}
