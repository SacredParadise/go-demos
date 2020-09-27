package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Hanmeimei")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	r.Run()    //default 8080 port
}
