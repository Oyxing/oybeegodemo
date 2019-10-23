package controllers

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

type SendState struct {
	Name    string
	Url     string
	Types   string
	Message string
	Email   string
}

//1.添加邮件服务器
func SendMail(tomail, url string) bool {
	// 邮箱地址
	UserEmail := "dev@idcyw.cn"
	// 端口号，:25也行 587
	Mail_Smtp_Port := "587"
	//邮箱的授权码，去邮箱自己获取
	Mail_Password := "kaifaCESHI123"
	// 此处填写SMTP服务器
	Mail_Smtp_Host := "smtp.exmail.qq.com"
	//用户昵称
	nickname := "柚备idcyw.com"
	user := UserEmail
	name := "柚备账户激活链接"
	message := "点击链接激活账户"
	auth := smtp.PlainAuth("", UserEmail, Mail_Password, Mail_Smtp_Host)
	t, _ := template.ParseFiles("views/mail/ceshi.html")
	var body bytes.Buffer
	headers := "MIME-version:1.0;\nContent-Type: text/html;"
	body.Write([]byte(fmt.Sprintf("From: %s<%s>\r\nSubject: "+name+"\r\n%s\r\n\r\n", nickname, user, headers)))
	var sendstate SendState
	sendstate.Email = tomail
	sendstate.Name = name
	sendstate.Url = url
	sendstate.Message = message
	sendstate.Types = "激活账户"
	err := t.Execute(&body, sendstate)
	fmt.Println(err)
	to := []string{tomail}
	err = smtp.SendMail(Mail_Smtp_Host+":"+Mail_Smtp_Port, auth, user, to, body.Bytes())
	if err != nil {
		fmt.Printf("send mail error: %v", err)
		return false
	}
	return true
}
