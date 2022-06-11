package views

import (
	"example.com/mod/cache"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	cache.RedisInit()
	cache.Set("name", "xu756", 6000)
	c.JSONP(200, gin.H{
		"message": "测试",
		"data":    cache.Get("name"),
	})
}
