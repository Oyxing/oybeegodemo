package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/objcoding/wxpay"
	"github.com/odeke-em/rsc/qr"
)

func init() {
	Setcongif()
	str := RandomOsn()
	fmt.Println("===111==")
	fmt.Println(str)
	fmt.Println("===11  ==")
}

/*
#微信支付参数配置
wxmchid=1486136622
wxappid=wx154bc001e55c7805
wxappkey=d02877925dcb9be25663f86070675024


#微信授权登录参数配置
AppID=wx9e8dc9840454a617
AppSecret=814ef8559cb6e3d4269898251b71d471
*/

func Setcongif() {
	beego.AppConfig.Set("wxmchid", "1560223251")
	beego.AppConfig.Set("wxappid", "wx7aa4851a39b7c99a")
	beego.AppConfig.Set("wxappkey", "DsWalAlBeUDa6jlDLAZ0CVriX8rNZsMr")
	beego.AppConfig.Set("AppID", "wx7aa4851a39b7c99a")
	beego.AppConfig.Set("AppSecret", "6253ec2a794c23b9ba2e7a9fe95e5bf9")
	beego.AppConfig.Set("domainurl", "www.5ul.cn")
	beego.AppConfig.Set("Spbill_create_ip", "125.76.228.183")
}

func RandomOsn() string {

	tm := time.Unix(time.Now().Unix(), 0)
	strtime := tm.Format("20060102030405")
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	order_sn := strtime + string(result)
	return order_sn
}

type WxPay struct {
	Appid      string `json:"appid"`
	ResultCode string `json:"result_code"`
	Sign       string `json:"sign"`
	PrepayId   string `json:"prepay_id"`
	TradeType  string `json:"trade_type"`
	CodeUrl    string `json:"code_url"`
	ReturnCode string `json:"return_code"`
	ErrCode    string `json:"err_code"`
	ErrCodeDes string `json:"err_code_des"`
	ReturnMsg  string `json:"return_msg"`
	MchId      string `json:"mch_id"`
	NonceStr   string `json:"nonce_str"`
}

func (c *MainController) Wxpayment() {
	client := wxpay.NewClient(wxpay.NewAccount("wx7aa4851a39b7c99a", "1560223251", "zNlC7qFDPmvnoygzVip46I5tkWQrJGxE", false))
	// 统一下单
	/*
		result_code:FAIL
		appid:wx7aa4851a39b7c99a
		mch_id:1560223251
		nonce_str:xgqHDwfEc7rKHHj9
		sign:B0B7A6A0C7C614F8EA63B18B5A4EDF84
		err_code:INVALID_REQUEST
		err_code_des:201 商户订单号重复
		return_code:SUCCESS
		return_msg:OK
	*/
	params := make(wxpay.Params)
	params.SetString("body", "test222").
		SetString("out_trade_no", "201910281251244").
		SetInt("total_fee", 10).
		SetString("spbill_create_ip", "125.76.228.183").
		SetString("notify_url", "https://www.5ul.cn/wxpay/notify").
		SetString("trade_type", "NATIVE").
		SetString("openid", "o1aT74lX564FKPTP9pcJDbWak3_E").
		SetString("product_id", "wwddd")

	p, _ := client.UnifiedOrder(params)

	fmt.Println("Wxpayment")
	fmt.Println(p)
	fmt.Println(p["prepay_id"])
	fmt.Println("Wxpayment")
	jsonStr, err := json.Marshal(p)
	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(jsonStr))
	var wxpays WxPay
	err = json.Unmarshal(jsonStr, &wxpays)
	fmt.Println(err)
	fmt.Println(wxpays)
	if wxpays.ReturnCode == "FAIL" {
		// appid参数长度有误
		fmt.Println(wxpays.ReturnMsg)
	}
	if wxpays.ResultCode == "FATL" {
		fmt.Println(wxpays.ErrCodeDes)
	}
	wxpays.CodeUrl = Imgs(wxpays.CodeUrl)
	stringA := "appid=wx7aa4851a39b7c99a&body=test&device_info=1000&mch_id=1560223251&nonce_str=" + wxpays.NonceStr
	stringSignTemp := stringA + "&key=zNlC7qFDPmvnoygzVip46I5tkWQrJGxE" //注：key为商户平台设置的密钥key
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(stringSignTemp))
	cipherStr := md5Ctx.Sum(nil)
	sign := strings.ToUpper(hex.EncodeToString(cipherStr))
	//注：MD5签名方式
	fmt.Println("sign")
	fmt.Println(sign)
	c.Data["data"] = wxpays
	c.Data["Image"] = wxpays.CodeUrl

	c.TplName = "payindex.tpl"

}

func Imgs(url string) string {
	code, err := qr.Encode(url, qr.H)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	imgByte := code.PNG()
	str := base64.StdEncoding.EncodeToString(imgByte)
	return str
}
