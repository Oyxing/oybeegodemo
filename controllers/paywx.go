/*
@author:wanjianning@qq.com
@create_time: 2018-11-15 11:37:29
@description:weixin的统-下单处理操作
*/
package controllers

import (
	"encoding/base64"
	"exernew/models/Wxpay"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/odeke-em/rsc/qr"
)

type WxpayController struct {
	beego.Controller
}

func (c *WxpayController) Native() {
	// uuid := c.GetSession("uuid").(string)
	// order_id := c.GetString("order_id")
	// order_id_str, err := base64.URLEncoding.DecodeString(order_id)
	// if err != nil {
	// 	fmt.Println(err)

	// }
	// orderid := string(order_id_str)
	// order := models.QueryOrderInfoById(orderid)
	//orderNumber := this.Ctx.Input.Param(":id") //获取订单号
	// orderNumber := order.Order_sn
	// payAmount := order.Total_fee
	params := make(map[string]interface{})
	params["body"] = "柚备数据备份软件_1" //显示标题
	params["out_trade_no"] = 1123131
	params["total_fee"] = 1
	params["product_id"] = 1
	params["attach"] = "www.idcyw.cn" //自定义参数
	var modwx Wxpay.UnifyOrderReq
	res := modwx.CreateOrder(c.Ctx, params)
	c.Data["data"] = res
	//fmt.Println(res)
	//统一下单后生成数据，转成二维码。
	b, _ := base64.StdEncoding.DecodeString(Img(res.Code_url))
	fmt.Printf("%T", b)
	// c.Layout = "public/Index.tpl"
	c.Data["Website"] = "beego.messs"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Image"] = Img(res.Code_url)

	c.TplName = "payindex.tpl"

}

func (c *WxpayController) Notify() {
	var notifyReq Wxpay.WXPayNotifyReq
	res := notifyReq.WxpayCallback(c.Ctx)
	//beego.Debug("res",res)
	if res != nil {
		//这里可以组织res的数据 处理自己的业务逻辑：
		sendData := make(map[string]interface{})
		sendData["id"] = res["out_trade_no"]
		sendData["trade_no"] = res["transaction_id"]
		paid_time, _ := time.Parse("20060102150405", res["time_end"].(string))
		paid_timestr := paid_time.Format("2006-01-02 15:04:05")
		sendData["paid_time"] = paid_timestr
		sendData["payment_type"] = "wxpay"
		intfee := res["cash_fee"].(int)
		floatfee := float64(intfee)
		cashfee := floatfee / 100
		sendData["payment_amount"] = strconv.FormatFloat(cashfee, 'f', 2, 32)
		//处理业务逻辑
		fmt.Println("-------------------------异步通知--------------")
	}
	c.Data["json"] = "cccc"
	c.ServeJSON()
}

func Img(url string) string {
	code, err := qr.Encode(url, qr.H)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	imgByte := code.PNG()

	str := base64.StdEncoding.EncodeToString(imgByte)

	return str
}
