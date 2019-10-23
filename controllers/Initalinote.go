package controllers

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

// 生成 验证码
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// 阿里短信
func Initalinote() {
	client, err := dysmsapi.NewClientWithAccessKey(region, accessKeyId, accessSecret)
	if err != nil {
		// Handle exceptions
		panic(err)
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	// 电话
	request.PhoneNumbers = "18309298895"
	// 标签名
	request.SignName = "中光电信"
	// 模板code
	request.TemplateCode = "SMS_173945502"
	// 短信参数 验证码
	// strarr = `{"company": sendmsg.Companyname ,"dues":sendmsg.Name ,"date":sendmsg.Date,"renewal": renewaltime,"yearnum":sendmsg.Agenum ,"money":` + sendmsg.Totalfeesum + `}`

	request.TemplateParam = "{'code':'" + GenValidateCode(6) + "'}"
	// 发送
	response, err := client.SendSms(request)

	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
