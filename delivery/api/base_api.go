package api

import (
	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-api-with-gin/delivery/api/response"
	"github.com/mitchellh/mapstructure"
)

type BaseAPI struct{}

func (b *BaseAPI) ParseRequestBody(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(body); err != nil {
		return err
	}
	return nil
}

func (b *BaseAPI) ParseRequestFormData(c *gin.Context, requestModel interface{}, postFormKey ...string) error {
	mapRes := make(map[string]interface{})
	for _, v := range postFormKey {
		mapRes[v] = c.PostForm(v)
	}
	if err := mapstructure.Decode(mapRes, &requestModel); err != nil {
		return err
	}
	return nil
}

func (b *BaseAPI) Success(c *gin.Context, data interface{}) {
	response.NewSuccessJSONResponse(c, data).Send()
}

func (b *BaseAPI) SuccessDownload(c *gin.Context, filePath string) {
	response.NewSuccessFileResponse(c, filePath)
}

func (b *BaseAPI) Failed(c *gin.Context, err error) {
	response.NewErrorJSONResponse(c, err).Send()
}
