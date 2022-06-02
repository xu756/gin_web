package router

import (
	"example.com/mod/views"
	"github.com/gin-gonic/gin"
)

// UserRouter 用户模块
func UserRouter(r *gin.Engine) *gin.Engine {
	rr := r.Group("/api/user")
	rr.GET("/", views.Index)                                           // 首页
	rr.POST("/login", views.UserLogin)                                 // 登录
	rr.POST("/register", views.UserRegister)                           // 注册
	rr.GET("/userinfo/user=:token", views.UserInfo)                    // 获取用户信息
	rr.GET("/set/portrait", views.UploadPortrait)                      // 上传头像
	rr.StaticFS("/get/portrait", gin.Dir("./media/upload/user", true)) // 获取头像
	return r
}
