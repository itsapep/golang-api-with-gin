package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileResponse struct {
	c              *gin.Context
	httpStatusCode int
	fileName       string
}

func (j *FileResponse) Send() {
	j.c.File(j.fileName)
}

func NewSuccessFileResponse(c *gin.Context, fileName string) AppHTTPResponse {
	return &FileResponse{
		c:              c,
		httpStatusCode: http.StatusOK,
		fileName:       fileName,
	}
}

// func NewErrorFileResponse(c *gin.Context, err error) AppHTTPResponse {
// 	fmt.Println("===>", err)
// 	httpStatusCode, resp := NewErrorMessage(err)
// 	return &FileResponse{
// 		c:              c,
// 		httpStatusCode: httpStatusCode,
// 	}
// }

// func NewGlobalFileResponse(c *gin.Context, httpStatusCode int, response Response) AppHTTPResponse {
// 	return &FileResponse{
// 		c:              c,
// 		httpStatusCode: httpStatusCode,
// 	}
// }
