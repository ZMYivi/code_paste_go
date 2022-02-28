package main

import (
	"code_paste/biz/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func routerRegister(r *gin.Engine) {
	r.GET("/code_paste/get_code/:key", dec(handlers.GetCode))
	r.POST("/code_paste/save_code/", dec(handlers.SaveCode))
}

func dec(f func(*gin.Context) map[string]string) gin.HandlerFunc {
	h := func(c *gin.Context) {
		res := f(c)
		if len(res) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "wrong",
			})
		} else {
			jsonMap := gin.H{}
			for k, v := range res {
				jsonMap[k] = v
			}
			c.JSON(http.StatusOK, jsonMap)
		}
	}
	return h
}
