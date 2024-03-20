package ginupload

import (
	"fmt"
	"go-food-delivery/common"
	"go-food-delivery/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(appctx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
			Id:        0,
			Url:       "http://localhost:8080/static/" + fileHeader.Filename,
			Width:     0,
			Height:    0,
			CloudName: "",
			Extension: "",
		}))

	}
}
