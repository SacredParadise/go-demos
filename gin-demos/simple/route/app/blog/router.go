package blog

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/blog/comment", commentHandler)
	e.GET("/blog/post", postHandler)
}