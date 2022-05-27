package main

import (
	"example.com/mod/views"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", views.Index)
	r.POST("/api/login", views.UserLogin)
	r.POST("/api/register", views.UserRegister)
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("启动失败")
		return
	}
}
