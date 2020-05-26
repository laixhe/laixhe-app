package utils

import (
	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func GinJsonResponse(c *gin.Context, code int, msg string, data interface{}) {

	res := &JsonResponse{Code: code}

	if msg != "" {
		res.Message = msg
	} else {
		res.Message = CodeText(code)
	}

	if data != nil {
		res.Data = data
	}

	c.JSON(200, res)
}

func GinJsonResponseData(c *gin.Context, data interface{}){
	GinJsonResponse(c, SUCCESS, "", data)
}

func GinJsonResponseMsg(c *gin.Context, code int, msg string){
	GinJsonResponse(c, code, msg, nil)
}

func GinJsonResponseCode(c *gin.Context, code int){
	GinJsonResponse(c, code, "", nil)
}
