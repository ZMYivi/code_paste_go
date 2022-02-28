package handlers

import (
	"code_paste/biz/logic"
	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) map[string]string {
	key := c.Params.ByName("key")
	code, err := logic.GetCode(key)
	if err != nil {
		return nil
	}
	return map[string]string{"code": code}
}
