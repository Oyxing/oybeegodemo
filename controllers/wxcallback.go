package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

type Access_token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	SessionKey   string `json:"session_key"`
	Openid       string
	Scope        string
	Unionid      string
}

//access_token 出错相应
type AccessTokenErrorResponse struct {
	Errcode float64
	Errmsg  string
}
type WxUserInfo struct {
	Nickname   string
	Language   string
	City       string
	Headimgurl string
	Openid     string
	Sex        int
	Province   string
	Country    string
	Privilege  []string
	Unionid    string
}

type WxRes struct {
	State int
	Res   WxUserInfo
	Msg   interface{}
}

type WxCallbackController struct {
	beego.Controller
}

func (c *WxCallbackController) Getuserinfo() {
	code := c.Ctx.Input.Param(":code")
	res := getInfoByOauth(code, "wx35383c55c7dd277e", "5450bdbd4ed58b7641261dbc488561ad")
	c.Data["json"] = res
	c.ServeJSON()
}

func getInfoByOauth(code, appid, secret string) *WxRes {
	fmt.Println(code)
	fmt.Println(appid)
	fmt.Println(secret)
	var refreshtoken string
	token_url := strings.Join([]string{
		"https://api.weixin.qq.com/cgi-bin/token",
		"?appid=", appid,
		"&secret=", secret,
		"&js_code=", code,
		"&grant_type=authorization_code"}, "")
	resp, err := http.Get(token_url)
	if err != nil || resp.StatusCode != http.StatusOK {

		return &WxRes{1, WxUserInfo{}, fmt.Sprintln("发送get请求access_token地址获取错误", err)}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return &WxRes{1, WxUserInfo{}, fmt.Sprintln("发送get请求access_token地址,读取返回body错误", err)}
	}
	if bytes.Contains(body, []byte("errcode")) {
		var access_token_err_struct AccessTokenErrorResponse
		err = json.Unmarshal(body, &access_token_err_struct)
		if err != nil {
			return &WxRes{1, WxUserInfo{}, fmt.Sprintln("发送get请求获取access_token地址 的错误信息:" + err.Error())}
		} else {
			return &WxRes{1, WxUserInfo{}, fmt.Sprintln("access_token获取失败1" + access_token_err_struct.Errmsg)}
		}
	} else { //正确获取到信息
		var access_token_struct Access_token
		fmt.Println(string(body))
		err = json.Unmarshal(body, &access_token_struct)
		if err != nil {

			return &WxRes{1, WxUserInfo{}, fmt.Sprintln("发送get请求access_token地址,返回数据json解析错误", err)}
		}
		fmt.Println(access_token_struct)
		refreshtoken = access_token_struct.RefreshToken
	}
	var new_access_token string
	refresh_token_url := strings.Join([]string{"https://api.weixin.qq.com/sns/oauth2/refresh_token",
		"?appid=", appid,
		"&refresh_token=", refreshtoken,
		"&grant_type=refresh_token"}, "")
	resp, err = http.Get(refresh_token_url)
	if err != nil || resp.StatusCode != http.StatusOK {

		return &WxRes{1, WxUserInfo{}, fmt.Sprintln("发送get请求access_token地址获取错误", err)}
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {

		return &WxRes{1, WxUserInfo{}, fmt.Sprintln("发送get请求refresh_token地址,读取返回body错误", err)}
	}
	var refresh_token_struct Access_token
	if bytes.Contains(body, []byte("errcode")) {
		var access_token_err_struct AccessTokenErrorResponse
		err = json.Unmarshal(body, &access_token_err_struct)
		if err != nil {
			fmt.Println()
			return &WxRes{1, WxUserInfo{}, fmt.Sprintln("发送get请求获取refresh_token地址 的错误信息:" + err.Error())}
		} else {

			return &WxRes{1, WxUserInfo{}, fmt.Sprintln("access_token获取失败2" + access_token_err_struct.Errmsg)}
		}
	} else { //正确获取到信息
		fmt.Println(string(body))
		err = json.Unmarshal(body, &refresh_token_struct)
		if err != nil {

			return &WxRes{1, WxUserInfo{}, fmt.Sprintln("发送get请求refresh_token地址,返回数据json解析错误", err)}
		}
		new_access_token = refresh_token_struct.AccessToken
	}
	user_info_url := strings.Join([]string{
		"https://api.weixin.qq.com/sns/userinfo",
		"?access_token=", new_access_token,
		"&openid=", refresh_token_struct.Openid,
		"&lang=zh_CN"}, "")
	resp, err = http.Get(user_info_url)
	if err != nil || resp.StatusCode != http.StatusOK {

		return &WxRes{1, WxUserInfo{}, fmt.Sprintln("请求user_info地址获取错误", err)}
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {

		return &WxRes{1, WxUserInfo{}, fmt.Sprintln("发送get请求user_info 地址，读取返回body错误", err)}
	}
	var userinfo WxUserInfo
	if bytes.Contains(body, []byte("errcode")) {
		var access_token_err_struct AccessTokenErrorResponse
		err = json.Unmarshal(body, &access_token_err_struct)
		if err != nil {

			return &WxRes{1, WxUserInfo{}, fmt.Sprintln("发送get请求获取user_info地址 的错误信息:" + err.Error())}
		} else {

			return &WxRes{1, WxUserInfo{}, fmt.Sprintln("access_token获取失败3" + access_token_err_struct.Errmsg)}
		}
	} else { //正确获取到信息
		err = json.Unmarshal(body, &userinfo)
		if err != nil {

			return &WxRes{1, WxUserInfo{}, fmt.Sprintln("发送get请求user_info地址,返回数据json解析错误", err)}
		}
		return &WxRes{0, userinfo, "成功"}
	}

}
