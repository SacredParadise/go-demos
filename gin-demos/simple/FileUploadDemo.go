package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传文件出错")
		}

		err = c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			c.String(500, fmt.Sprintf("保存文件错误: %s", err))
		}
		fmt.Println(file.Filename)
		c.String(http.StatusOK, file.Filename)
	})

	r.Run()
}
