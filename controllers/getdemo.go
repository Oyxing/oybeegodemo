package controllers

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//获取参数方法设置
// func GetParams(this beego.Controller) models.JsonInfo {
// 	var msg models.JsonInfo
// 	jsoninfo := this.GetString("jsoninfo")
// 	json.Unmarshal([]byte(jsoninfo), &msg)
// 	return msg
// }
type Token struct {
	Token string `json:"token"`
}

const (
	SecretKey = "wwwqqqqqqq"
)

func (m *MainController) Login() {
	msg := GetParams(m)
	fmt.Println("=====5======")
	fmt.Println(msg.Password)
	fmt.Println(msg.Name)
	fmt.Println("=====5======")
	m.Data["json"] = ResApi{0, "成功" + m.Ctx.Input.IP(), LoginHandler()}
	m.ServeJSON()
}
func (m *MainController) GetLogin() {
	m.Data["json"] = ResApi{0, "成功" + m.Ctx.Input.IP(), GetTocker()}
	m.ServeJSON()
}

type MyCustomClaims struct {
	Audience bool   `json:"audience"`
	Name     string `json:"name"`
	Created  int64  `json:"created"`
	jwt.StandardClaims
}

func LoginHandler() string {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := MyCustomClaims{
		true,
		"name",
		time.Now().Unix(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 150000,
			Issuer:    "tdddest",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	return ss
}

func GetTocker() string {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWRpZW5jZSI6dHJ1ZSwibmFtZSI6Im5hbWUiLCJjcmVhdGVkIjoxNTYyMTM4MjU1LCJleHAiOjE1NzcxMzgyNTUsImlzcyI6InRkZGRlc3QifQ.0nITwBr8uAHYP6DM0kAO25SDyQBI0Tt8VbLk8GN1HZM"
	// at(time.Unix(0, 0), func() {
	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	fmt.Println("====1====")
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Println(claims.StandardClaims.ExpiresAt)
		fmt.Println("\n")
		fmt.Println(claims.StandardClaims.Issuer)
		// return claims.StandardClaims.Issuer
	} else {
		fmt.Println("====3====")
		fmt.Println(err)
	}
	// })
	return "1112"
}
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
func Prepare(tokenString string) (ResApi, bool) {
	//在执行所有方法之前 会先执行它
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if token.Valid {
		//验证通过没有错误 可以在这个里面做取值 操作 如果不需要可以跳过
		return ResApi{0, "验证通过", 0}, true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			//这不是一个令牌
			return ResApi{1, "这不是一个令牌", "That's not even a token"}, false

		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// 令牌已过期或者尚未激活
			return ResApi{1, "令牌已过期或者尚未激活", "Token is either expired or not active yet"}, false

		} else {
			//无法处理这个token
			return ResApi{1, "无法处理这个token", err}, false
		}
	} else {
		//无法处理这个token
		return ResApi{1, "无法处理这个token", err}, false
	}

}

// // 获取 token
// func GetToken(name, facilityname string) string {
// 	// Create the Claims
// 	mySigningKey := []byte("AllEzDocker")
// 	claims := MyCustomClaims{
// 		true,
// 		"name",
// 		time.Now().Unix(),
// 		facilityname,
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Unix() + 360000,
// 			Issuer:    "EzDocker",
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	ss, err := token.SignedString(mySigningKey)
// 	if err != nil {
// 		return "token 失败"
// 	}
// 	return ss
// }
