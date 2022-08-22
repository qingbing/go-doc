package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("openssl-key/rsa_private.pem")
	if err != nil {
		log.Fatal("文件读取失败")
	}
	if rsaPrivate, err := jwt.ParseRSAPrivateKeyFromPEMWithPassword(data, "111111"); err != nil {
		fmt.Println("解析失败1")
	} else {
		fmt.Printf("%+#v\n\n", rsaPrivate)
	}
	data, err = ioutil.ReadFile("openssl-key/rsa_public.pem")
	if err != nil {
		log.Fatal("文件读取失败")
	}
	if rsaPub, err := jwt.ParseRSAPublicKeyFromPEM(data); err != nil {
		fmt.Println("解析失败1")
	} else {
		fmt.Printf("%+#v\n\n", rsaPub)
	}

	// token: New
	toke := jwt.New(jwt.SigningMethodHS256)
	toke.Claims = jwt.MapClaims{
		"name": "qing",
		"sex":  "male",
		"age":  19,
	}
	fmt.Println(toke)

	// token: NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": "qing",
		"sex":  "male",
		"age":  19,
	})
	fmt.Println(token)

	me := jwt.GetSigningMethod("HS384")
	fmt.Println(me)
	bs01 := []byte("http://www.baidu.com/?name=sss&sex=12")
	enStr := jwt.EncodeSegment(bs01)
	fmt.Println(enStr)
	if deBs, err := jwt.DecodeSegment(enStr); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(deBs))
	}
}
