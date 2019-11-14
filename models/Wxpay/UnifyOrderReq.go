package Wxpay

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type UnifyOrderReq struct {
	Attach           string `xml:"attach" json:"attach"`
	Appid            string `xml:"appid" json:"appid"`
	Body             string `xml:"body" json:"body"`
	Detail           string `xml:"detail" json:"detail"`
	Fee_type         string `xml:"fee_type" json:"fee_type"`
	Goods_tag        string `xml:"goods_tag" json:"goods_tag"`
	Mch_id           string `xml:"mch_id" json:"mch_id"`
	Openid           string `xml:"openid" json:"openid"`
	Nonce_str        string `xml:"nonce_str" json:"nonce_str"`
	Notify_url       string `xml:"notify_url" json:"notify_url"`
	Product_id       string `xml:"product_id" json:"product_id"`
	Time_start       string `xml:"time_start" json:"time_start"`
	Time_expire      string `xml:"time_expire" json:"time_expire"`
	Trade_type       string `xml:"trade_type" json:"trade_type"`
	Spbill_create_ip string `xml:"spbill_create_ip" json:"spbill_create_ip"`
	Total_fee        int    `xml:"total_fee" json:"total_fee"`
	Out_trade_no     string `xml:"out_trade_no" json:"out_trade_no"`
	Sign             string `xml:"sign" json:"sign"`
}

type UnifyOrderResp struct {
	Return_code string `xml:"return_code" json:"return_code"`
	Return_msg  string `xml:"return_msg" json:"return_msg"`
	Attach      string `xml:"attach" json:"attach"`
	Appid       string `xml:"appid" json:"appid"`
	Mch_id      string `xml:"mch_id" json:"mch_id"`
	Nonce_str   string `xml:"nonce_str" json:"nonce_str"`
	Sign        string `xml:"sign" json:"sign"`
	Result_code string `xml:"result_code" json:"result_code"`
	Prepay_id   string `xml:"prepay_id" json:"prepay_id"`
	Trade_type  string `xml:"trade_type" json:"trade_type"`
	Code_url    string `xml:"code_url" json:"code_url"`
}

//国内时间转换一致
func GetCurrentTime() time.Time {
	location := time.FixedZone("Asia/Shanghai", +8*60*60)
	abnow := time.Now().UTC()
	_now := abnow.In(location)
	return _now
}
func TimeConvert(span int) string {
	_now := GetCurrentTime()
	if span == 2 {
		_now = _now.Add(time.Hour * 2)
	}
	return _now.Format("20060102150405")
}

func (o *UnifyOrderReq) CreateOrder(ctx *context.Context, param map[string]interface{}) UnifyOrderResp {
	fmt.Println("====11===")
	xmlResp := UnifyOrderResp{}
	unify_order_req := "https://api.mch.weixin.qq.com/pay/unifiedorder"
	var yourReq UnifyOrderReq
	yourReq.Attach = param["attach"].(string)
	yourReq.Appid = beego.AppConfig.String("wxappid") //微信开放平台我们创建出来的app的app id
	yourReq.Body = param["body"].(string)
	yourReq.Detail = "西安东美信息科技有限公司"
	yourReq.Fee_type = "CNY"
	yourReq.Goods_tag = "WXG"
	yourReq.Mch_id = beego.AppConfig.String("wxmchid")
	yourReq.Nonce_str = randStr(32, "alphanum")
	yourReq.Notify_url = "http://" + beego.AppConfig.String("domainurl") + "/wxpay/notify" //异步返回的地址
	yourReq.Product_id = "111231"
	yourReq.Time_start = TimeConvert(1)
	yourReq.Time_expire = TimeConvert(2)
	yourReq.Trade_type = "JSAPI"
	yourReq.Openid = "o1aT74lbX1f4BRk9tlnho-P7iHvc"
	//yourReq.Spbill_create_ip = ctx.Input.IP()
	yourReq.Spbill_create_ip = "125.76.228.183"
	totalFee, _ := strconv.ParseFloat("10", 64)
	totalfeeint := totalFee * 100
	yourReq.Total_fee = int(totalfeeint) //单位是分
	//beego.Debug("totalfee*100",totalfeeint)
	//beego.Debug("totalfee*int",int(totalfeeint))
	//beego.Debug("totalfee*1your",yourReq.Total_fee)
	yourReq.Out_trade_no = "20191012053634523"

	fmt.Printf("%+v", yourReq)

	//beego.Debug("yourReqEntity", yourReq)
	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["attach"] = yourReq.Attach
	m["appid"] = yourReq.Appid
	m["body"] = yourReq.Body
	m["detail"] = yourReq.Detail
	m["fee_type"] = yourReq.Fee_type
	m["goods_tag"] = yourReq.Goods_tag
	m["mch_id"] = yourReq.Mch_id
	m["openid"] = yourReq.Openid
	m["nonce_str"] = yourReq.Nonce_str
	m["notify_url"] = yourReq.Notify_url
	m["product_id"] = yourReq.Product_id
	m["time_start"] = yourReq.Time_start
	m["time_expire"] = yourReq.Time_expire
	m["trade_type"] = yourReq.Trade_type
	m["spbill_create_ip"] = yourReq.Spbill_create_ip
	m["total_fee"] = yourReq.Total_fee
	m["out_trade_no"] = yourReq.Out_trade_no
	appkey := beego.AppConfig.String("wxappkey")
	yourReq.Sign = WxpayCalcSign(m, appkey) //这个是计算wxpay签名的函数上面已贴出
	fmt.Println("\n")
	fmt.Printf("%+v", yourReq)
	fmt.Println("\n")

	//beego.Debug("yourReq",yourReq)
	bytes_req, err := xml.Marshal(yourReq)
	if err != nil {
		fmt.Println("以xml形式编码发送错误, 原因:", err)
		return xmlResp
	}

	str_req := string(bytes_req)
	//wxpay的unifiedorder接口需要http body中xmldoc的根节点是<xml></xml>这种，所以这里需要replace一下
	str_req = strings.Replace(str_req, "UnifyOrderReq", "xml", -1)
	fmt.Println(str_req)
	bytes_req = []byte(str_req)

	//发送unified order请求.
	req, err := http.NewRequest("POST", unify_order_req, bytes.NewReader(bytes_req))
	if err != nil {
		fmt.Println("New Http Request发生错误，原因:", err)
		return xmlResp
	}
	req.Header.Set("Accept", "application/xml")
	//这里的http header的设置是必须设置的.
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")

	c := http.Client{}
	resp, _err := c.Do(req)
	if _err != nil {
		fmt.Println("请求微信支付统一下单接口发送错误, 原因:", _err)
		return xmlResp
	}

	//xmlResp :=UnifyOrderResp{}
	body, _ := ioutil.ReadAll(resp.Body)
	_err = xml.Unmarshal(body, &xmlResp)
	fmt.Printf("%+v", xmlResp)
	if xmlResp.Return_code == "FAIL" {
		fmt.Println("微信支付统一下单不成功，原因:", xmlResp.Return_msg)
		return xmlResp
	}

	//beego.Debug("xmlResp",xmlResp)
	//这里已经得到微信支付的prepay id，需要返给客户端，由客户端继续完成支付流程
	fmt.Println("微信支付统一下单成功，预支付单号:", xmlResp.Prepay_id)

	return xmlResp

}

func randStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func WxpayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	fmt.Println("微信支付签名计算, API KEY:", key)
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

// {Attach:飞洒发生的非䔯基本原理 艺术形式䤂
// 	Appid:wx7aa4851a39b7c99a
// 	Body:幅度还是对符合条件儿童游客脱衣和腺癌而入而入而已
// 	Detail: Fee_type:CNY
// 	Goods_tag:WXG
// 	Mch_id:1560223251
// 	Openid:
// 	Nonce_str:Cj5HIfHrzK82tYn3GwlDw1frpcxqW6aD
// 	Notify_url:https://www.5ul.cn/wxpay/notify
// 	Product_id:20191012053646017
// 	Time_start:20191030104924
// 	Time_expire:20191030124924
// 	Trade_type:NATIVE
// 	 Spbill_create_ip:125.76.228.183
// 	 Total_fee:100
// 	 Out_trade_no:20191012053646017
// 	 Sign:D8EABECB480DC2CDFE0C7FC3BDF831F0
// }
// {Attach:www.idcyw.cn
// 	Appid:wx7aa4851a39b7c99a
// 	Body:柚备数据备份软件_1
// 	Detail:西安东美信息科技有限公司
// 	Fee_type:CNY
// 	Goods_tag:WXG
// 	Mch_id:1560223251
// 	Openid:
// 	Nonce_str:PVv6OSA2qa6rj6xmaW2UAjsILRgU87iF
// 	Notify_url:http://www.5ul.cn/wxpay/notify
// 	Product_id:1
// 	Time_start:20191030104945
// 	Time_expire:20191030124945
// 	Trade_type:NATIVE
// 	Spbill_create_ip:125.76.228.183
// 	Total_fee:1000
// 	Out_trade_no:201910120536460
// 	Sign:905FFE521E0E291C324E70A6D261EB43}
