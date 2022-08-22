package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

type myClaim2 struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func (receiver myClaim2) Valid() error {
	return nil
}

// 加密，解密 只使用字符串
func main() {
	// 初始化 claims 信息
	claims := myClaim2{
		UserName: "qingbing",
		Email:    "780042175@qq.com",
		StandardClaims: jwt.StandardClaims{
			//NotBefore: time.Now().Unix() - 60, // 一分钟之前开始生效
			NotBefore: time.Now().Unix() + 60,      // 一分钟之后开始生效
			ExpiresAt: time.Now().Unix() + 60*60*2, // 两个小时后失效
			Issuer:    "qiyezhu",                   // 签发人
		},
	}

	priKeyBs, err := ioutil.ReadFile("openssl-key/rsa_private.pem")
	if err != nil {
		log.Fatal("读取私钥失败", err)
	}
	rsaPriKey, err := jwt.ParseRSAPrivateKeyFromPEMWithPassword(priKeyBs, "111111")
	if err != nil {
		log.Fatal("解析私钥失败", err)
	}

	// 获取 token string
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenStr, err := token.SignedString(rsaPriKey) // key 要和 NewWithClaims 的 method 对应
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("====== tokenStr ======")
	fmt.Println(tokenStr)

	// 解密token
	pToken, err := jwt.ParseWithClaims(tokenStr, &myClaim2{}, func(token *jwt.Token) (interface{}, error) {
		pubKeyBs, err := ioutil.ReadFile("openssl-key/rsa_public.pem")
		if err != nil {
			return "", errors.New("读取公钥失败")
		}
		return jwt.ParseRSAPublicKeyFromPEM(pubKeyBs)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("====== 解码后信息 ======")
	if err := pToken.Claims.(*myClaim2).Valid(); err != nil {
		log.Fatal(err)
	}
	//if pToken.Valid {
	//	fmt.Println("解析后数据有效")
	//	fmt.Println(pToken.Claims.(*myClaim2).UserName)
	//	fmt.Println(pToken.Claims.(*myClaim2).Email)
	//} else {
	//	fmt.Println("解析后数据无效")
	//	fmt.Println(pToken.Claims.(*myClaim2).UserName)
	//	fmt.Println(pToken.Claims.(*myClaim2).Email)
	//}
}
