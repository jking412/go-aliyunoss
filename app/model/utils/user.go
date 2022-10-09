package utils

import (
	"aliyunoss/app/model"
	"aliyunoss/pkg/database"
)

func CreateUser(u *model.User) error {
	return database.DB.Create(u).Error
}

func GetUser(u *model.User) (*model.User, error) {
	user := &model.User{}
	err := database.DB.Where("username = ?", u.Username).First(user).Error
	return user, err
}

func DeleteUser(u *model.User) error {
	return database.DB.Where("username = ?", u.Username).Delete(u).Error
}
