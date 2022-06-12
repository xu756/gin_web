package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default() // 初始化路由
	//404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "请求路径不存在",
		})

	})
	UserRouter(r)         // 用户模块
	err := r.Run(":7000") // 监听端口
	fmt.Println("启动成功")
	if err != nil {
		fmt.Println("启动失败") // 启动失败
		return
	}
}
