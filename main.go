package main

import (
	"encoding/json"
	"exernew/controllers"
	_ "exernew/routers"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

type A struct {
	Created  time.Time
	Time     string
	Hardware string
}

func main() {
	// controllers.WebSocket()
	// go pi_pie()
	// controllers.Balance()
	// fmt.Println(controllers.GenValidateCode(6))
	// controllers.Initqcloud()
	// wxRes := controllers.GetInfoByOauth("001mOmb40EYtIJ119wb40CDnb40mOmb3", "wx9e8dc9840454a617", "814ef8559cb6e3d4269898251b71d471")

	// controllers.SendMail("897346175@qq.com", "http://localhost:8083/api/get")
	str := "****qU9A=="
	key := "******P5Y2b9SfejeA=="
	iv := "*****FLgzU09FtANlRw=="
	src, err := controllers.Dncrypt(str, key, iv)
	fmt.Println(err)
	var s = map[string]interface{}{}
	json.Unmarshal([]byte(src), &s)
	fmt.Printf("== %+v", src)
	fmt.Printf("cc== %+v", s)
	strs := "aaaaa"
	pathstr := []string{
		".",
		"static",
		"upload",
		strs,
	}

	fmt.Println(strings.Replace(strings.Trim(fmt.Sprint(pathstr), "[]"), " ", "//", -1))

	// controllers.Main_test()
	// fmt.Println(t)

	// //获取当前时间戳
	// fmt.Println(t.Unix()) //1531293019

	// //获得当前的时间
	// fmt.Println(t.Format("2006年01月02日"))
	//2018-7-15 15:23:00

	//时间 to 时间戳
	// loc, _ := time.LoadLocation("Asia/Shanghai")                     //设置时区
	// tt, _ := time.ParseInLocation("2006年01月02日", "2018年07月11日", loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	// fmt.Println(tt.Unix())                                           //1531292871

	// //时间戳 to 时间
	// tm := time.Unix(1531293019, 0)
	// fmt.Println(tm.Format("2006年01月02日")) //2018-07-11 15:10:19

	// //获取当前年月日,时分秒
	// y := t.Year()                 //年

	// d := t.Day()                  //日
	// h := t.Hour()                 //小时
	// i := t.Minute()               //分钟
	// s := t.Second()               //秒
	// fmt.Println(y, m, d, h, i, s) //2018 July 11 15 24 59

	t := time.Unix(1576893901, 0) //2018-07-11 15:07:51.8858085 +0800 CST m=+0.004000001

	m := t.Month() //月
	fmt.Println(m)

	// h := md5.New()
	// h.Write([]byte("asdasasdasda" + "idcyw2018")) // 需要加密的字符串为 加盐加密 idcyw2018
	// cipherStr := h.Sum(nil)
	// usercode := hex.EncodeToString(cipherStr)
	// url := "http://mail.qq.com/user/activation?email=asdasasdasda&code=" + usercode
	// fmt.Println(url)

	// haha := A{}

	// a, e := d.CreateHardware()
	// if e != nil {
	// 	fmt.Println(e)
	// 	return
	// }
	// haha.Hardware = a
	// haha.Time = "aaaaa"
	// // aescode := "hgfedcba87654321"
	// w, _ := json.Marshal(haha)
	// fmt.Println(w)
	// //生成秘钥对
	// bits := 2048
	// err = d.GenerateRSAKey(bits, "./key/private/private.pem", "./key/public/public.pem")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// tk := toolbox.NewTask("myTask", "0 35 15 * * 0-6", func() error { fmt.Println("hello worldddddddd"); return nil })
	tk := toolbox.NewTask("myTask", "0 0 16 * * *", func() error { fmt.Println("hello worldddddddd"); return nil })
	// err = tk.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	toolbox.AddTask("myTask", tk)
	toolbox.StartTask()
	formatTime, err := time.Parse("2006年1月2", "2002年8月26")
	fmt.Println(formatTime)
	fmt.Println(err)
	beego.Run()
}

// // 通道使用
// func pi_pie() {
// 	pipe := make(chan string, 1024) //make关键字创建一个管道（关键字chan），管道内装int类型的数据，并且管道大小能装10个数字，超过则阻塞
// 	pipe <- "10"                    //向管道内放入数据
// 	pipe <- "8"
// 	pipe <- "9"
// 	pipe <- "7" email=asdasasdasda&code=abcf0ef43c3bf336db065f277feb6b10

// 	fmt.Println("p1", <-pipe)
// 	fmt.Println("p2", <-pipe)
// 	fmt.Println("p3", <-pipe)
// 	fmt.Println("p4", <-pipe)

// }
