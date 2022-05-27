package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int       `primaryKey:"true"`                     // 自增ID
	UserName  string    `gorm:"type:varchar(100);unique_index"` // 用户名
	Password  string    `gorm:"type:varchar(256)"`              // 密码
	Emial     string    `gorm:"type:varchar(100);unique_index"` // 邮箱
	Token     string    `gorm:"type:varchar(100)"`              // token
	Role      int       `gorm:"type:int(1);default:1"`          // 角色
	Frequency int       `gorm:"type:int(1000)"`                 // 访问频率
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`     // 创建时间
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`     // 更新时间
}

var DB *gorm.DB

func InitMysqlDB() {
	//dsn := "xu:xjx756756@tcp(121.5.132.57:5700)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	DB = db
	err = DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("用户表创建失败")
		return
	}
}
