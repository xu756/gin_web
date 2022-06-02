package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()    // 初始化路由
	UserRouter(r)         // 用户模块
	err := r.Run(":7000") // 监听端口
	fmt.Println("启动成功")
	if err != nil {
		fmt.Println("启动失败") // 启动失败
		return
	}
}
