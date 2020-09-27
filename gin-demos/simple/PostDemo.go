package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		name := c.PostForm("name")
		password := c.PostForm("password")

		c.String(http.StatusOK, fmt.Sprintf("username:%s, password:%s, type:%s", name, password, types))
	})

	r.Run(":8000")
}
