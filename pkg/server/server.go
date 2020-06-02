package server

import "github.com/gin-gonic/gin"

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.New()
	return r
}


