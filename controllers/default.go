package controllers

import (
	"encoding/json"
	"exernew/models" //  引入路由表
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	d "gitee.com/countpoison/Doraemon"
	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
)

func init() {

}

// func (u Users) ReflectCall(name string, age int64) {
// 	fmt.Println("name:", name, "age:", age, "user.name", u.Name)
// }

// func (u Users) ReflectNoCall() {
// 	fmt.Println("ReflectNoCall")
// 	fmt.Println(u.Name)
// }

func (c *MainController) UploadingPost() {
	// 文章小图上传

	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	for _, v := range beego.AppConfig.Strings("configserip") {
		c.Ctx.Output.Header("Access-Control-Allow-Origin", v)
	}
	imgtype := c.GetString("action")
	uuid := c.GetString("uuid")
	var topfilename string
	var returnfilename string
	var hfilename string
	// licenselogo-执照 userlogo-头像 logocompany-企业logo platformlogo-系统配置图标 contentsimg-文章小图 setattrilogo-广告图
	if imgtype != "" {
		if uuid != "" {
			topfilename = imgtype + "\\" + uuid
			returnfilename = imgtype + "/" + uuid
		} else {
			returnfilename = imgtype
			topfilename = imgtype
		}
		hfilename = Updatelogo(c, topfilename, imgtype)
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/static/upload/" + returnfilename + "/" + hfilename, "title": hfilename, "original": hfilename}
	} else {
		c.Data["json"] = map[string]interface{}{"state": "DEFEATED", "url": "", "title": "", "original": ""}
	}
	c.ServeJSON()
}

func Updatelogo(c *MainController, topfilename, filetype string) string {
	fmt.Println(".\\static\\upload\\" + topfilename + "\\")
	err := os.MkdirAll(".\\static\\upload\\"+topfilename+"\\", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	_, h, err := c.GetFile(filetype)
	if err != nil {
		beego.Error(err)
	}
	path1 := ".\\static\\upload\\" + topfilename + "\\" + h.Filename
	err = c.SaveToFile(filetype, path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
	if err != nil {
		beego.Error(err)
	}
	return h.Filename
}

func jwts() {
	Info := d.Info{
		"name",
		"",
		time.Now(),
	}
	pp, err := Info.CreateToken("name")
	fmt.Println(pp)
	fmt.Println(err)
	Parsetoken(pp)
}

func Parsetoken(Token string) {
	info, err := d.ParseToken(Token, "name")
	fmt.Println(info)
	fmt.Println(err)
}

type MainController struct {
	beego.Controller
}
type tests interface {
	Type() int
}

func Type() int {
	return 1
}
func (c *MainController) Get() {
	_, err := models.AddUser("name", 2)
	if err == nil {
		c.Data["sessus"] = "成功1"
	}
	c.Data["Website"] = "beego.messs"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Username"] = c.GetSession("username")
	usermenuarr := c.GetSession("usermenuarr")
	fmt.Println("======1=====")

	fmt.Println(usermenuarr)
	fmt.Println("======2=====")

	c.TplName = "index.tpl"
}

func (c *MainController) GetApijson() {
	ages := make(map[string]string)
	ages["aaa"] = "ddddd"
	ages["ffff"] = "aaa"

	c.Data["json"] = ResApi{1, "aa", ages}
	c.ServeJSON()
}

func (c *MainController) UpdateGet() {
	user := new(models.User)
	user.Name = "aaaa"
	user.Age = 4112

	user.Mode = models.GetMode("1")
	_, err := models.UpdateUser(user)
	if err == nil {
		c.Data["sessus"] = "成功1"
	}
	fmt.Println("=====")
	fmt.Println(err)
	c.Data["Website"] = "UpdateGet.messs"
	c.Data["Email"] = "UpdateGet@gmail.com"
	c.TplName = "index.tpl"
}
func (c *MainController) GetApi() {

	_, err := models.AddMont("1", "你好", 2)
	if err == nil {
		c.Data["sessus"] = "成功GetApi"
	}

	c.Data["Website"] = "beego.messs"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *MainController) GetUsermsg() {
	p, err := models.GetUsermsg()
	intarr := []int{1, 2, 3}
	c.SetSession("username", "asdaad")
	c.SetSession("usermenuarr", intarr)
	if err == nil {
		c.Data["sessus"] = "成功GetUsermsgaaaa"
		c.Data["data"] = p
	}
	c.Data["Website"] = "beego.messs"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Clickfun"] = Clickfun()
	fmt.Println("-----")
	fmt.Println(time.Now())
	fmt.Println("-----")

	Postform := c.Ctx.Request.PostForm
	fmt.Println("==========")
	fmt.Printf("%+v", Postform)
	newjson := make(map[string]interface{})
	for k, v := range Postform {
		numint, Atoierr := strconv.Atoi(v[0])
		if Atoierr == nil {
			newjson[k] = numint
			continue
		} else if v[0] == "false" {
			newjson[k] = false
			continue
		} else if v[0] == "true" {
			newjson[k] = true
			continue
		}
		newjson[k] = v[0]
	}

	user := new(models.User)
	mjson, _ := json.Marshal(newjson)
	mString := string(mjson)
	json.Unmarshal([]byte(mString), user)
	fmt.Println("user")
	fmt.Println(user)

	c.TplName = "index.tpl"
}

func Clickfun() string {
	return "sadasda"
}

func (c *MainController) PostRegister() {
	fmt.Printf("%+v", c.Ctx.Request)
}

func (c *MainController) GetUser() {
	p, err := models.GetUser()
	fmt.Printf("%+v", p)
	if err == nil {
		c.Data["sessus"] = "成功GetUser"
		c.Data["json"] = p
	}

	for _, v := range p {
		for _, v1 := range v.Mode {
			fmt.Println("==")
			fmt.Println(v1.Id)
			fmt.Println("==")
		}
	}
	c.Data["Website"] = "beego.messs"
	c.Data["Email"] = "astaxie@gmail.com"

	c.ServeJSON()
}

// 开启服务
func WebSocket() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		fmt.Println("ListenAndServe", err)
	}
}

type ApiCont struct {
	Apiname string            `json:"apiname,omitempty"`
	Apidata map[string]string `json:"apidata,omitempty"`
}

type Demodata struct {
	User string `json:"user"`
	Age  string `json:"age"`
}
type Demodata1 struct {
	Useradmin string `json:"useradmin"`
	Ageadmin  string `json:"ageadmin"`
}

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("can't receive")
			break
		}
		apicont := new(ApiCont)
		err := json.Unmarshal([]byte(reply), apicont)
		if err != nil {
			fmt.Println(err)
		}
		demodata := new(Demodata)
		demodata1 := new(Demodata1)
		apidata, _ := json.Marshal(apicont.Apidata)
		err = json.Unmarshal([]byte(apidata), demodata)
		err = json.Unmarshal([]byte(apidata), demodata1)
		demodata.demotest(ws)
		demodata1.demotest(ws)
	}

}

func (d Demodata) demotest(ws *websocket.Conn) {
	if err := websocket.Message.Send(ws, d.Age+d.User); err != nil {
	}
}

func (d Demodata1) demotest(ws *websocket.Conn) {
	if err := websocket.Message.Send(ws, d.Ageadmin+d.Useradmin); err != nil {
	}
}

func test(r rune) bool {
	return false
}

type ResApi struct {
	Success int         `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type Usersapi map[string]Users

/*
	Funcuser    *Users `json:"funcuser,omitempty"`
	FuncuserDee *Users `json:"funcuserdee,omitempty"`
*/

type Users struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Deee string `json:"deee,omitempty"`
}

func (m *MainController) ApiJsonifo() {
	fmt.Println("==1======")
	// json := GetParams(m)
	var userapi Usersapi
	buf := make([]byte, 1024)
	n, _ := m.Ctx.Request.Body.Read(buf)
	fmt.Println("==---")
	fmt.Println(string(buf[0:n]))

	json.Unmarshal([]byte(string(buf[0:n])), &userapi)
	fmt.Printf("%+v", userapi)
	fmt.Println("===3=====")

	for aa, v := range userapi {
		fmt.Println("===5=====")
		Funcuser := reflect.ValueOf(v)
		v.Deee = "12312"
		v.Name = "qwqe"
		methodValue := Funcuser.MethodByName(aa)
		if aa == "Funcusers" {
			args := []reflect.Value{reflect.ValueOf("wudebao")}
			stra := methodValue.Call(args)
			fmt.Println(stra[0].Interface())
			fmt.Println("===7=====")

			m.Data["json"] = stra[0].Interface()
		} else if aa == "FuncuserDees" {
			fmt.Println("===11=====")
			stra := methodValue.Call(nil)
			fmt.Println(stra[0].Interface())
			fmt.Println("===7=====")

			m.Data["json"] = stra[0].Interface()
		}

	}
	// jwts()
	m.ServeJSON()
}

func (u Users) Funcusers() ResApi {
	fmt.Println("===9====")
	fmt.Println("\n")
	fmt.Println("===Name===")
	fmt.Println("dee", u.Name)
	fmt.Println("Deee", u.Deee)
	return ResApi{0, "Funcusers", 0}
}
func (u Users) FuncuserDees() ResApi {
	fmt.Println("\n")
	fmt.Println("===Deee===")
	fmt.Println("dee", u.Deee)
	return ResApi{0, "FuncuserDees", u}

}

//获取参数方法设置
func GetParams(this *MainController) models.JsonInfo {
	var msg models.JsonInfo
	buf := make([]byte, 1024)
	n, _ := this.Ctx.Request.Body.Read(buf)
	json.Unmarshal([]byte(string(buf[0:n])), &msg)
	return msg
}

var (
	mu      sync.Mutex
	balance int
)

func Balance() int {
	mu.Lock()
	b := 1
	mu.Unlock()
	return b
}

//获取参数方法设置
func (c *MainController) SetSessions() {
	c.SetSession("value", "存session")
	value := c.GetSession("value")
	fmt.Println("value1")
	fmt.Println(value)
	c.Data["json"] = "存session"
	c.ServeJSON()
}

//获取参数方法设置
func (c *MainController) GetSessions() {
	value := c.GetSession("value")
	fmt.Println("value")
	fmt.Println(value)
	c.Data["json"] = value
	c.ServeJSON()
}

// 结构体 转 map
func StructToMapDemo(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[strings.ToLower(obj1.Field(i).Name)] = obj2.Field(i).Interface()
	}
	return data
}
