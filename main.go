package main

import (
	. "example.com/mod/cache"
	. "example.com/mod/config"
	. "example.com/mod/models"
	. "example.com/mod/router"
)

func main() {
	UploadConfig() //上传配置
	InitRouter()   //路由配置
	InitMysqlDB()  //数据库配置
	RedisInit()    //redis配置
}
