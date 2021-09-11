package model

import (
	"github.com/debuginn/GoShortLink/dao"
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

//---- User Model CURD

// create a user function
func CreateAUser(user *User) (err error) {
	err = dao.DB.Create(&user).Error
	return
}

// get a user by telephone
func GetAUserByTelephone(telephone string) (user *User, err error) {
	user = &User{}
	if err := dao.DB.Where("telephone=?", telephone).First(&user).Error; err != nil {
		return nil, err
	}
	return
}

// get a user by userId
func GetAUserById(userId uint) (user *User, err error) {
	user = &User{}
	if err := dao.DB.Where("id=?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return
}
