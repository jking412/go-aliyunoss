package model

type User struct {
	Id       int    `json:"id" gorm:"primary_key;auto_increment;colum:id"`
	Username string `json:"username" gorm:"unique;not null;colum:username"`
	Password string `json:"password" gorm:"not null;colum:password"`
	Salt     string `json:"salt;" gorm:"not null;colum:salt"`
}
