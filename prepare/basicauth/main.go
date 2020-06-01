package main

import "github.com/gin-gonic/gin"

/**
测试基本的http用户登陆验证
 */

func main1() {
	r := gin.Default()
	r.Use(gin.BasicAuth(gin.Accounts{
		"sun":"123",
		"sun2":"456",
	}))
	r.Run()
}
