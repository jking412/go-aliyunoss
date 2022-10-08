package jwt

import (
	"fmt"
	jwtpkg "github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

type JWTCustomClaims struct {
	UserId int
	jwtpkg.StandardClaims
}

func AuthJWT(r *http.Request) (*jwtpkg.Token, error) {
	tokenString, err := GetJwtTokenFromHeader(r)
	if err != nil {
		return nil, err
	}

	token, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ParseToken(tokenString string) (*jwtpkg.Token, error) {
	token, err := jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secret")), nil
	})
	return token, err
}

func GenerateToken(userId int) (string, error) {
	claims := JWTCustomClaims{
		userId,
		jwtpkg.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600,
			Issuer:    "aliyunoss",
		},
	}
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("jwt.secret")))
}

func GetJwtTokenFromHeader(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	fmt.Println(tokenString)

	tokenResult := strings.SplitN(tokenString, " ", 2)
	if len(tokenResult) != 2 || tokenResult[0] != "Bearer" {
		return "", fmt.Errorf("token格式错误")
	}

	return tokenResult[1], nil
}
