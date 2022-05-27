package views

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSONP(200, gin.H{
		"message": "测试",
	})

}
