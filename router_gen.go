package main

import (
	"code_paste/biz/handlers"
	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	routerRegister(r)

	r.GET("/ping", handlers.Ping)
}
