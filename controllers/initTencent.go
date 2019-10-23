package controllers

import (
	"fmt"

	qcloudsms "github.com/qichengzx/qcloudsms_go"
)

func InitTencent() {
	var smsreq qcloudsms.SMSSingleReq
	opt := qcloudsms.NewOptions(appid, appkey, sign)
	var client = qcloudsms.NewClient(opt)
	params := []string{
		client.NewRandom(6).Random,
	}
	smsreq.Params = params
	smsreq.Sign = "中光电信"
	smsreq.Time = 10000000
	smsreq.Tel.Nationcode = "86"
	smsreq.Tel.Mobile = "18309298895"
	smsreq.TplID = 412035
	_, err := client.SendSMSSingle(smsreq)
	fmt.Println(err)

}
