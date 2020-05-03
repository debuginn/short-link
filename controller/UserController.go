package controller

import (
	"GoShortLink/common"
	"GoShortLink/dto"
	"GoShortLink/model"
	"GoShortLink/response"
	"GoShortLink/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// 注册用户
func Register(context *gin.Context) {
	var requestUser = model.User{}
	err := context.Bind(&requestUser)
	if err != nil {
		response.ServerError(context, nil, "请求参数错误")
		return
	}

	// 获取参数
	username := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	// 数据验证
	if len(telephone) != 11 {
		response.UnprocessableEntity(context, nil, "手机号必须为 11位")
		return
	}
	if len(password) <= 6 {
		response.UnprocessableEntity(context, nil, "密码不能少于 6位")
		return
	}
	if len(username) == 0 {
		username = util.RandString(10)
	}
	// 判断手机号是否占用
	if isTelephoneExist(telephone) {
		response.UnprocessableEntity(context, nil, "用户已经存在")
		return
	}

	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(context, nil, "加密密码错误")
		return
	}

	newUser := model.User{
		Name:      username,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}

	errCreate := model.CreateAUser(&newUser)
	if errCreate != nil {
		response.Created(context, nil, "用户创建失败")
	} else {
		// 发放 token
		token, err := common.ReleaseToken(&newUser)
		if err != nil {
			response.Fail(context, nil, "生成 Token 失败")
			return
		}
		response.Success(context, &gin.H{"token": token}, "用户创建成功")
	}
}

// 用户登录
func Login(context *gin.Context) {
	// 获取参数
	var requestUser = model.User{}
	err := context.Bind(&requestUser)
	if err != nil {
		response.ServerError(context, nil, "请求参数错误")
		return
	}
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 数据验证
	if len(telephone) != 11 {
		response.UnprocessableEntity(context, nil, "手机号必须为 11位")
		return
	}
	if len(password) <= 6 {
		response.UnprocessableEntity(context, nil, "密码不能少于 6位")
		return
	}
	// 判断手机号是否注册
	user, err := model.GetAUserByTelephone(telephone)
	if err != nil || user.ID == 0 {
		response.UnprocessableEntity(context, nil, "手机号未注册")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.ServerError(context, nil, "密码错误")
		log.Printf("token generate err: %v", err)
		return
	}
	// 发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Fail(context, nil, "生成 Token 失败")
		return
	}
	// 返回结果
	response.Success(context, &gin.H{"token": token}, "登录成功")
}

// 获取信息
func Info(context *gin.Context) {
	// 从上下文中读取 user 信息
	user, _ := context.Get("user")
	// 处理 dto 数据
	dtoUser := dto.ToUserDto(user.(*model.User))
	// 返回封装
	response.Success(context, &gin.H{"user": dtoUser}, "")
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
