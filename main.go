package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	r := gin.Default()
	DB := InitDB()
	defer DB.Close()

	r.POST("/api/auth/register", func(context *gin.Context) {
		// 获取参数
		username := context.PostForm("username")
		telephone := context.PostForm("telephone")
		password := context.PostForm("password")

		// 数据验证
		if len(telephone) != 11 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号必须为 11位",
			})
			return
		}
		if len(password) <= 6 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码不能少于 6位",
			})
			return
		}
		if len(username) == 0 {
			username = RandString(10)
		}
		// 判断手机号是否占用
		if isTelephoneExist(DB, telephone) {
			context.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "用户已经存在",
			})
			return
		}

		// 创建用户
		newUser := User{
			Name:      username,
			Telephone: telephone,
			Password:  password,
		}

		DB.Create(&newUser)

		// 反馈结果
		log.Println(username, password, telephone)
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户创建成功",
		})
	})

	panic(r.Run())
}

// 生成长度为 n 的随机字符串
func RandString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	// 使用 make 创建一个长度为 n 的字节类型的切片
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 连接数据库
func InitDB() *gorm.DB {
	driveName := "mysql"
	host := "localhost"
	port := "3306"
	database := "goshortlink"
	username := "root"
	password := "root"
	charset := "utf8"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driveName, args)
	if err != nil {
		panic("failed to connect database err:" + err.Error())
	}

	db.AutoMigrate(&User{})
	return db
}

// 判断手机号码是否占用
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)

	if user.ID != 0 {
		return true
	}
	return false
}
