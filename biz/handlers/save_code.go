package handlers

import (
	"code_paste/biz/logic"
	"github.com/gin-gonic/gin"
)

type SaveCodeContext struct {
	Code string `json:"code"`
}

func SaveCode(c *gin.Context) map[string]string {
	req := SaveCodeContext{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return nil
	}
	key, err := logic.SaveCode(req.Code)
	if err != nil {
		return nil
	}
	return map[string]string{"key": key}
}
