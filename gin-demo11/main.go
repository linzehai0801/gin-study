package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("index", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(file.Filename)

		// 上传文件到指定的目录
		dst := fmt.Sprint("./", file.Filename)
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"messge": fmt.Sprintf("'%s' upload!", file.Filename),
		})
	})
	r.Run()
}
