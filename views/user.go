package views

import (
	"crypto/sha256"
	"encoding/hex"
	"example.com/mod/models"
	"github.com/gin-gonic/gin"
	"log"
)

func UserLogin(c *gin.Context) {
	models.InitMysqlDB()
	var db = models.DB
	data := make(map[string]interface{}) // 定义一个map 存储客户端数据
	err := c.BindJSON(&data)             // 获取客户端数据
	if err != nil {
		log.Panicln("无法解析数据") // 如果解析失败 则报错
		return
	}
	// 对密码进行sha256加密
	h := sha256.New()
	h.Write([]byte(data["password"].(string)))
	data["password"] = hex.EncodeToString(h.Sum(nil))
	var user models.User
	db.Where("user_name = ?", data["username"]).Where("password = ?", data["password"]).First(&user)
	if user.Id == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "用户名或密码错误",
		})
		return
	}
	user.Password = data["password"].(string)
	db.Debug().Save(&user)
	c.JSON(200, gin.H{
		"code": 200,
		"ID":   user.Id,
		"msg":  "修改成功",
	})

}
