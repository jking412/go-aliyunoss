package utils

import (
	"aliyunoss/app/model"
	"aliyunoss/pkg/database"
)

func CreateUser(u *model.User) error {
	return database.DB.Create(u).Error
}

func GetUser(u *model.User) error {
	return database.DB.Where("username = ?", u.Username).First(u).Error
}
