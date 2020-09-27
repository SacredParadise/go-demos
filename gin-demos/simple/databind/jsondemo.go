package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User string `form:"username" json:"user" uri:"user" bingding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("loginJson", func(c *gin.Context) {
		var json Login

		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadGateway, gin.H{"status": "304"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.Run()
}
