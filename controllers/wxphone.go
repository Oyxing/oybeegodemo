package controllers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	d "gitee.com/countpoison/Doraemon"

	"net/http"
)

type Phonestruct struct {
	CountryCode     string `json:"countryCode"`
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
}

func (c *MainController) BoundPhone() {
	code := c.GetString("code")
	rawData := c.GetString("rawData")
	iv := c.GetString("iv")
	opedid, key, err := GetOpenid(code)
	fmt.Println("err", err)
	fmt.Println("opedid", opedid)
	phonesrc, phoneerr := Dncrypt(rawData, key, iv)
	fmt.Println("phoneerr", phoneerr)
	fmt.Printf("== %+v", phonesrc)
	c.Data["json"] = phonesrc
	c.ServeJSON()
}

//  获取用户 openid
func GetOpenid(code string) (string, string, string) {
	appid := "wx7aa4851a39b7c99a"                // 小程序 id
	secret := "6253ec2a794c23b9ba2e7a9fe95e5bf9" // 小程序私钥
	token_url := strings.Join([]string{
		"https://api.weixin.qq.com/sns/jscode2session",
		"?appid=", appid,
		"&secret=", secret,
		"&js_code=", code,
		"&grant_type=authorization_code"}, "")
	resp, err := http.Get(token_url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", "发送get请求access_token地址获取错误", fmt.Sprintln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "发送get请求access_token地址,读取返回body错误", fmt.Sprintln(err)
	}
	if bytes.Contains(body, []byte("errcode")) {
		var access_token_err_struct AccessTokenErrorResponse
		err = json.Unmarshal(body, &access_token_err_struct)
		if err != nil {
			return "", "发送get请求获取access_token地址 的错误信息", fmt.Sprintln(err)
		} else {
			errcode, _ := d.ToString(access_token_err_struct.Errcode)
			return "", "access_token获取失败2,", access_token_err_struct.Errmsg + errcode
		}
	}
	var access_token_struct Access_token
	err = json.Unmarshal(body, &access_token_struct)
	if err != nil {
		return "", "发送get请求access_token地址,返回数据json解析错误", fmt.Sprintln(err)
	}
	fmt.Println("access_token_struct")
	fmt.Printf("%+v", access_token_struct)
	return access_token_struct.Openid, access_token_struct.SessionKey, ""
}

// 微信 电话号 解析
func Dncrypt(rawData, key, iv string) (Phonestruct, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	key_b, err_1 := base64.StdEncoding.DecodeString(key)
	iv_b, errs := base64.StdEncoding.DecodeString(iv)
	if errs != nil {
		return Phonestruct{}, errs
	}
	if err != nil {
		return Phonestruct{}, err
	}
	if err_1 != nil {
		return Phonestruct{}, err_1
	}
	fmt.Println("data", data)
	dnData, err := AesCBCDncrypt(data, key_b, iv_b)
	if err != nil {
		return Phonestruct{}, err
	}
	return dnData, nil
}

var wx_try_times = 3

// 解密
func AesCBCDncrypt(encryptData, key, iv []byte) (Phonestruct, error) {
	if wx_try_times <= 0 {
		wx_try_times = 3
		return Phonestruct{}, errors.New("time out 3")
	}
	wx_try_times = wx_try_times - 1
	var phones = Phonestruct{}
	encryptDataf := encryptData
	block, err := aes.NewCipher(key)
	if err != nil {
		return phones, err
	}
	blockSize := block.BlockSize()
	if len(encryptDataf) < blockSize {
		return phones, errors.New("ciphertext too short")
	}
	if len(encryptDataf)%blockSize != 0 {
		return phones, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.BlockSize()
	mode.CryptBlocks(encryptDataf, encryptDataf)

	length := len(encryptDataf)

	unpadding := int(encryptDataf[length-1])
	if length-unpadding < 0 {
		d.Sleep(1000)
		return AesCBCDncrypt(encryptData, key, iv)
	}
	encryptDataf = encryptDataf[:(length - unpadding)]
	err = json.Unmarshal(encryptDataf, &phones)
	if err != nil {
		d.Sleep(1000)
		return AesCBCDncrypt(encryptData, key, iv)
	}
	return phones, err
}
