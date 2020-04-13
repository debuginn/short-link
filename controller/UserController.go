package controller

import (
	"GoShortLink/common"
	"GoShortLink/dto"
	"GoShortLink/model"
	"GoShortLink/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// 注册用户
func Register(context *gin.Context) {
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
		username = util.RandString(10)
	}
	// 判断手机号是否占用
	if isTelephoneExist(telephone) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已经存在",
		})
		return
	}

	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "加密密码错误",
		})
		return
	}

	newUser := model.User{
		Name:      username,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}

	errCreate := model.CreateAUser(&newUser)
	if errCreate != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg":  "用户创建失败",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户创建成功",
		})
	}
}

// 用户登录
func Login(context *gin.Context) {
	// 获取参数
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

	// 判断手机号是否注册
	user, err := model.GetAUserByTelephone(telephone)
	if err != nil || user.ID == 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号未注册",
		})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		log.Printf("token generate err: %v", err)
		return
	}

	// 发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "生成 Token 失败"})
		return
	}

	// 返回结果
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{"token": token},
	})
}

// 获取信息
func Info(context *gin.Context) {
	user, _ := context.Get("user")
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": gin.H{
			"user": dto.ToUserDto(user.(*model.User)),
		},
	})
}

// 判断手机号码是否占用
func isTelephoneExist(telephone string) bool {
	user, err := model.GetAUserByTelephone(telephone)
	if err != nil {
		return false
	}
	if user.ID != 0 {
		return true
	}

	return false
}
