package views

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"example.com/mod/models"
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"time"
)

// UserLogin 登录页面
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
	data["password"] = hex.EncodeToString(h.Sum(nil)) // 对密码进行加密
	var user models.User
	db.Where("user_name = ?", data["username"]).First(&user)
	if user.Id == 0 {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "用户名不存在",
		})
		return
	}
	if user.Password != data["password"] {
		c.JSON(200, gin.H{
			"code":    202,
			"message": "密码错误",
		})
		return
	}
	// 生成token
	h = sha512.New()
	h.Write([]byte(user.UserName + user.Password + time.Now().String()))
	token := hex.EncodeToString(h.Sum(nil))
	user.Token = token // 将token存入数据库
	user.Frequency = user.Frequency + 1
	db.Save(&user)
	c.JSON(200, gin.H{
		"code":  200,
		"token": token,
		"msg":   "登录成功",
	})
}

// UserRegister 注册页面
func UserRegister(c *gin.Context) {
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
	data["password"] = hex.EncodeToString(h.Sum(nil)) // 对密码进行加密
	var user models.User
	db.Where("user_name = ?", data["username"]).First(&user)
	if user.Id != 0 {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "用户名已存在",
		})
		return
	}
	user.UserName = data["username"].(string)
	user.Password = data["password"].(string)
	user.Role = data["role"].(string)
	h = sha512.New()
	h.Write([]byte(user.UserName + user.Password + time.Now().String()))
	token := hex.EncodeToString(h.Sum(nil))
	user.Token = token // 将token存入数据库
	user.Frequency = 1
	db.Create(&user)
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"token": token,
		},
		"msg": "注册成功",
	})
}

// UserInfo 获取用户信息
func UserInfo(c *gin.Context) {
	models.InitMysqlDB()
	var db = models.DB
	var user models.User
	token := c.Param("token") // 获取token
	db.Where("token = ?", token).First(&user)
	if user.Id == 0 {
		c.JSON(200, gin.H{
			"code": 300,
			"msg":  "用户不存在，请重新登录",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"userinfo": gin.H{
			"username": user.UserName,
			"emial":    user.Emial,
		},
		"msg": "获取用户信息成功",
	})
}

// UploadPortrait 上传头像
func UploadPortrait(c *gin.Context) {
	// 获取上传文件
	file, err := c.FormFile("file")
	if err != nil {
		{
			log.Panicln("无法获取上传文件")
			return
		}

	}
	// 获取文件名
	filename := file.Filename
	// 获取文件后缀
	ext := path.Ext(filename)
	// 将文件写入
	err = c.SaveUploadedFile(file, "./media/upload/user/"+filename)
	if err != nil {
		log.Panicln("无法保存文件")
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"data": gin.H{
			"filename": filename,
			"ext":      ext,
		},
	})
}
