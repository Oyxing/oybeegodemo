package controllers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
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
	beego.AppConfig.Set("wxmchid", "1486136622")
	beego.AppConfig.Set("wxappid", "wx154bc001e55c7805")
	beego.AppConfig.Set("wxappkey", "d02877925dcb9be25663f86070675024")
	beego.AppConfig.Set("AppID", "wx9e8dc9840454a617")
	beego.AppConfig.Set("AppSecret", "814ef8559cb6e3d4269898251b71d471")
	beego.AppConfig.Set("domainurl", "idcyw.cn")
	beego.AppConfig.Set("Spbill_create_ip", "125.76.228.213")

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
