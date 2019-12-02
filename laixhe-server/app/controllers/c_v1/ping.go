package c_v1

import (
	"github.com/gin-gonic/gin"

	"github.com/laixhe/laixhe-app/laixhe-server/config"
	"github.com/laixhe/laixhe-app/laixhe-server/utils"
)

func Ping(c *gin.Context) {

	token, err := utils.TokenGin(c)
	if err != nil {
		utils.GinJsonResponseMsg(c, utils.ERROR_AUTH, err.Error())
		return
	}

	utils.GinJsonResponseData(c, token)

}

func Token(c *gin.Context) {

	token := &utils.Token{Account:"aaaaaa-", No:123654}

	tokenString, err := utils.TokenGen(token, config.GetApp().Token)
	if err != nil {
		utils.GinJsonResponseMsg(c, utils.ERROR_AUTH, err.Error())
		return
	}

	utils.GinJsonResponseData(c, tokenString)

}

