package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		addr := c.Query("addr")
		c.JSON(200, gin.H{
			"messus":   "ok",
			"username": username,
			"addr":     addr,
		})
	})
	r.Run()
}

//访问http://127.0.0.1:8080/user/search?username=bb&addr=aa
//http://127.0.0.1:8080/user/search?addr=aa
