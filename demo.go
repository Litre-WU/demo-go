package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
            "msg": "来了,老弟！",
            "result": "你看这个面它又长又宽，就像这个碗它又大又圆",
		})
	})
	r.Run(":80")
}
