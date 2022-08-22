package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type myClaim1 struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// 加密，解密 只使用字符串
func main() {
	// 初始化 claims 信息
	claims := myClaim1{
		UserName: "qingbing",
		Email:    "780042175@qq.com",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,      // 一分钟之前开始生效
			ExpiresAt: time.Now().Unix() + 60*60*2, // 两个小时后失效
			Issuer:    "qiyezhu",                   // 签发人
		},
	}

	// 指定密钥 key
	jwtKey := []byte("secret key")

	// 获取 token string
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("====== tokenStr ======")
	fmt.Println(tokenStr)

	// 解密token
	pToken, err := jwt.ParseWithClaims(tokenStr, &myClaim1{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("====== 解码后信息 ======")
	fmt.Println(pToken.Claims.(*myClaim1).UserName)
	fmt.Println(pToken.Claims.(*myClaim1).Email)
}
