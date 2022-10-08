package user

import (
	"aliyunoss/app/model"
	"aliyunoss/app/model/utils"
	"github.com/gin-gonic/gin"
)

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	registerReq := &RegisterReq{}
	if err := c.ShouldBindJSON(registerReq); err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
			"error":   err.Error(),
		})
		return
	}
	user := &model.User{
		Username: registerReq.Username,
		Password: registerReq.Password,
		Salt:     "123456",
	}
	if err := utils.CreateUser(user); err != nil {
		c.JSON(400, gin.H{
			"message": "注册失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func Login(c *gin.Context) {
	loginReq := &LoginReq{}
	if err := c.ShouldBindJSON(loginReq); err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
			"error":   err.Error(),
		})
		return
	}

	user := &model.User{
		Username: loginReq.Username,
		Password: loginReq.Password,
	}

	if err := utils.GetUser(user); err != nil {
		c.JSON(400, gin.H{
			"message": "登录失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "登录成功",
	})
}
