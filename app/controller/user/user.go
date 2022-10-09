package user

import (
	"aliyunoss/app/model"
	"aliyunoss/app/model/utils"
	"aliyunoss/pkg/jwt"
	"aliyunoss/pkg/response"
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
		response.ErrorJSON(c, "参数错误")
		return
	}
	user := &model.User{
		Username: registerReq.Username,
		Password: registerReq.Password,
		Salt:     "123456",
	}
	if err := utils.CreateUser(user); err != nil {
		response.ErrorJSON(c, "注册失败")
		return
	}
	token, _ := jwt.GenerateToken(user.Id)
	response.SuccessJSONWithField(c, "注册成功", response.Field{
		"token": token,
	})
}

func Login(c *gin.Context) {
	loginReq := &LoginReq{}
	if err := c.ShouldBindJSON(loginReq); err != nil {
		response.ErrorJSON(c, "参数错误")
		return
	}

	user := &model.User{}

	if _user, err := utils.GetUser(&model.User{Username: loginReq.Username}); err != nil {
		response.ErrorJSON(c, "登录失败")
		return
	} else {
		user = _user
		if user.Password != loginReq.Password {
			response.ErrorJSON(c, "登录失败")
			return
		}
	}

	token, _ := jwt.GenerateToken(user.Id)
	response.SuccessJSONWithField(c, "登录成功", response.Field{
		"token": token,
	})
}
