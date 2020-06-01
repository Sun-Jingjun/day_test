package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Use(TestMiddleWare())

	router.Run()
}


/**
尝试编写一个中间件
 */

func TestMiddleWare() gin.HandlerFunc{
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		c.String(http.StatusOK,"this is a middleware demo")
	}
}