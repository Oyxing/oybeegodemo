package routers

import (
	"exernew/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// AllowAllOrigins: true,
		AllowOrigins:     beego.AppConfig.Strings("configserip"),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.SetStaticPath("/static", "static")
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/operation/?:id", &controllers.OperationController{}, "get:GetOperation;post:PostOperation;put:PutOperation;delete:DeleteOperation")
	beego.Router("/operation/delarr", &controllers.OperationController{}, "post:DeletearrOperation")
	beego.Router("/api/get", &controllers.MainController{}, "*:GetApi")
	beego.Router("/api/json", &controllers.MainController{}, "*:GetApijson")
	beego.Router("/api/updateget", &controllers.MainController{}, "*:UpdateGet")
	beego.Router("/api/getuser", &controllers.MainController{}, "*:GetUser")
	beego.Router("/api/getusermag", &controllers.MainController{}, "*:GetUsermsg")
	beego.Router("/api/jsonifo", &controllers.MainController{}, "*:ApiJsonifo")
	beego.Router("/api/login", &controllers.MainController{}, "*:Login")
	beego.Router("/api/getlogin", &controllers.MainController{}, "*:GetLogin")
	beego.Router("/api/register", &controllers.MainController{}, "*:PostRegister")
	beego.Router("/api/Websocke", &controllers.MyWebSocketController{}, "*:Websocketmsg")
	// beego.Router("/api/contpay", &controllers.MainController{}, "*:Contpay")
	beego.Router("/api/native", &controllers.WxpayController{}, "*:Native")
	beego.Router("/api/uploading", &controllers.MainController{}, "*:UploadingPost")
	beego.Router("/api/userinfo/?:code", &controllers.WxCallbackController{}, "*:Getuserinfo")
	beego.Router("/api/setsessions", &controllers.MainController{}, "*:SetSessions")
	beego.Router("/api/getsessions", &controllers.MainController{}, "*:GetSessions")

}
