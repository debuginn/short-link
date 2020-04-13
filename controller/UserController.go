package controller

import (
	"GoShortLink/model"
	"GoShortLink/util"
	"github.com/gin-gonic/gin"
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
	newUser := model.User{
		Name:      username,
		Telephone: telephone,
		Password:  password,
	}

	err := model.CreateAUser(&newUser)
	if err != nil {
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
