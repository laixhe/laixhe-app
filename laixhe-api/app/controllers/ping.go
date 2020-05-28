package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/laixhe/goutil/zap_log"

	"github.com/laixhe/laixhe-app/laixhe-api/app/models"
	"github.com/laixhe/laixhe-app/laixhe-api/config"
	"github.com/laixhe/laixhe-app/laixhe-api/utils"
)

func Ping(c *gin.Context) {

	//token, err := utils.TokenGin(c)
	//if err != nil {
	//	utils.GinJsonResponseMsg(c, utils.ERROR_AUTH, err.Error())
	//	return
	//}

	data , err := models.UserList()
	if err != nil {
		zap_log.Errorf("ping err:%v", err)
		utils.GinJsonResponseMsg(c, utils.ERROR_DB, err.Error())
		return
	}

	utils.GinJsonResponseData(c, data)

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

