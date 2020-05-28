package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/laixhe/goutil"
)

type Token struct {
	Account   string `json:"account"`       // 账号
	No        int    `json:"no"`            // ?
	IssuedAt  int64  `json:"iat,omitempty"` // 签发时间
	ExpiresAt int64  `json:"exp,omitempty"` // 过期时间
}

// 产生 token
func TokenGen(data *Token, tokenKey string) (string, error) {

	token, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// 加密
	cryted, err := goutil.AesEncrypt(token, []byte(tokenKey))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cryted), nil
}

// 校验 token
func TokenCheck(token string, tokenKey string) (*Token, error) {

	// 转成字节数组
	cryted, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	// 解密
	data, err := goutil.AesDecrypt(cryted, []byte(tokenKey))
	if err != nil {
		return nil, err
	}

	t := new(Token)
	err = json.Unmarshal(data, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func TokenGin(c *gin.Context) (*Token, error) {

	token, exists := c.Get("token")
	if !exists {
		return nil, errors.New("gin token empty")
	}

	return token.(*Token), nil
}