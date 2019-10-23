package controllers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
	"github.com/go-gomail/gomail"
)

var (
	appid        string = "1400252999"
	appkey       string = "328ef04289bde56c971e94a9c33fcb62"
	sign         string = "中光电信"
	region       string = "cn-hangzhou"
	accessKeyId  string = "LTAI4FsM6XHzMgBuskrLJ8zt"
	accessSecret string = "TiL41sDlQRCclSEhhbn0DJjAX5eKRY"
)

type QcloudController struct {
	beego.Controller
}

// 腾讯 短信

func Initqcloud() {
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
	request.TemplateCode = "SMS_173475615"
	// 短信参数 验证码
	request.TemplateParam = "{'code':'1111'}"
	// 发送
	response, err := client.SendSms(request)

	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

type EmailParam struct {
	// ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerHost string
	// ServerPort 邮箱服务器端口，如腾讯企业邮箱为465
	ServerPort int
	// FromEmail　发件人邮箱地址
	FromEmail string
	// FromPasswd 发件人邮箱密码（注意，这里是明文形式），TODO：如果设置成密文？
	FromPasswd string
	// Toers 接收者邮件，如有多个，则以英文逗号(“,”)隔开，不能为空
	Toers string
	// CCers 抄送者邮件，如有多个，则以英文逗号(“,”)隔开，可以为空
	CCers string
}

func TestEmail(t *testing.T) {
	serverHost := "smtp.exmail.qq.com"
	serverPort := 465
	fromEmail := "cicd@latelee.org"
	fromPasswd := "1qaz@WSX"

	myToers := "li@latelee.org, latelee@163.com" // 逗号隔开
	myCCers := ""                                //"readchy@163.com"

	subject := "这是主题"
	body := `这是正文<br>
            <h3>这是标题</h3>
             Hello <a href = "http://www.latelee.org">主页</a><br>`
	// 结构体赋值
	myEmail := &EmailParam{
		ServerHost: serverHost,
		ServerPort: serverPort,
		FromEmail:  fromEmail,
		FromPasswd: fromPasswd,
		Toers:      myToers,
		CCers:      myCCers,
	}
	t.Logf("init email.\n")
	InitEmail(myEmail)
	SendEmail(subject, body)
}

// 全局变量，因为发件人账号、密码，需要在发送时才指定
// 注意，由于是小写，外面的包无法使用
var serverHost, fromEmail, fromPasswd string
var serverPort int

var m *gomail.Message

func InitEmail(ep *EmailParam) {
	toers := []string{}

	serverHost = ep.ServerHost
	serverPort = ep.ServerPort
	fromEmail = ep.FromEmail
	fromPasswd = ep.FromPasswd

	m = gomail.NewMessage()

	if len(ep.Toers) == 0 {
		return
	}

	for _, tmp := range strings.Split(ep.Toers, ",") {
		toers = append(toers, strings.TrimSpace(tmp))
	}

	// 收件人可以有多个，故用此方式
	m.SetHeader("To", toers...)

	//抄送列表
	if len(ep.CCers) != 0 {
		for _, tmp := range strings.Split(ep.CCers, ",") {
			toers = append(toers, strings.TrimSpace(tmp))
		}
		m.SetHeader("Cc", toers...)
	}

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	m.SetAddressHeader("From", fromEmail, "")
}

// SendEmail body支持html格式字符串
func SendEmail(subject, body string) {
	// 主题
	m.SetHeader("Subject", subject)

	// 正文
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer(serverHost, serverPort, fromEmail, fromPasswd)
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}
