package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	c              *gin.Context
	httpStatusCode int
	response       Response
}

func (j *JSONResponse) Send() {
	j.c.JSON(j.httpStatusCode, j.response)
}

func NewSuccessJSONResponse(c *gin.Context, data interface{}) AppHTTPResponse {
	httpStatusCode, resp := NewSuccessMessage(data)
	return &JSONResponse{
		c:              c,
		httpStatusCode: httpStatusCode,
		response:       resp,
	}
}

func NewErrorJSONResponse(c *gin.Context, err error) AppHTTPResponse {
	fmt.Println("===>", err)
	httpStatusCode, resp := NewErrorMessage(err)
	return &JSONResponse{
		c:              c,
		httpStatusCode: httpStatusCode,
		response:       resp,
	}
}

func NewGlobalJSONResponse(c *gin.Context, httpStatusCode int, response Response) AppHTTPResponse {
	return &JSONResponse{
		c:              c,
		httpStatusCode: httpStatusCode,
		response:       response,
	}
}
