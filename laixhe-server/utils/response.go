package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/laixhe/laixhe-app/laixhe-server/protoapi"
)

func GinJsonResponse(c *gin.Context, code protoapi.ErrorCode, msg string) {

	res := &protoapi.ErrorResponse{Code: code}

	if msg != "" {
		res.Message = msg
	} else {
		res.Message = protoapi.ErrorCode_name[int32(code)]
	}

	c.JSON(http.StatusOK, res)
}

func GinJsonResponseMsg(c *gin.Context, code protoapi.ErrorCode, msg string){
	GinJsonResponse(c, code, msg)
}

func GinJsonResponseCode(c *gin.Context, code protoapi.ErrorCode){
	GinJsonResponse(c, code, "")
}
