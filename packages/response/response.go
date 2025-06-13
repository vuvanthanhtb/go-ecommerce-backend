package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// success reponse
func SuccessResponse(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// error reponse
func ErrorResponse(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusBadGateway, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
