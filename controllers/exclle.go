package controllers

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

//  用户
type User struct {
	Id           string    `json:"id" xorm:"pk notnull unique 'id'"`
	Parentid     string    `json:"parentid"`        // 上级uuid
	Enable       string    `json:"enable"`          // 是否激活  "yes"  "no"
	Nickname     string    `json:"nickname"`        // 昵称
	Username     string    `json:"username"`        // 用户名
	Password     string    `json:"password"`        // 密码
	Email        string    `json:"email"`           // 邮箱
	Qq           int64     `json:"qq"`              // qq
	Phonenum     string    `json:"phonenum"`        // 电话
	Logo         string    `json:"logo"`            // 头像
	Groups       string    `json:"groups"`          // 角色级别id
	Groupsname   string    `json:"groupsname"`      // 角色级别
	Gender       int       `json:"grnder"`          // 性别
	Province     string    `json:"province"`        // 籍贯
	Duty         string    `json:"duty"`            // 职务
	Birthday     time.Time `json:"birthday"`        // 生日
	Thetitle     string    `json:"thetitle"`        // 职称
	Societytitle string    `json:"societytitle"`    // 社会职务
	Address      string    `json:"address"`         // 家庭住址
	Deed         string    `json:"deed"`            // 个人 事迹
	Honor        string    `json:"honor"`           // 个人 荣誉
	Honorimg     []string  `json:"honorimg"`        // 个人 荣誉图片
	Comments     string    `json:"comments"`        // 参与公益事业情况
	Education    string    `json:"education"`       // 个人 学历
	Position     string    `json:"position"`        // 申请职务类型
	Enddate      time.Time `json:"enddate"`         // 过期时间
	Inittime     time.Time `json:"inittime"`        // 入会时间
	Tags         []string  `json:"tags"`            // 标签
	Tagrollout   []string  `json:"industryrollout"` // 需求的资源
	Tagbelow     []string  `json:"industrybelow"`   // 提供资源
	Integral     int       `json:"integral"`        // 交钱数
	Created      time.Time `json:"created"`
}

//  企业用户
type Company struct {
	Id                  string    `json:"id" xorm:"pk notnull unique 'id'"`
	Useruuid            string    `json:"useruuid"`            // 用户id
	Logo                string    `json:"logo"`                // 企业
	Legalperson         string    `json:"legalperson"`         // 法人名称
	Legalpersonphonenum string    `json:"legalpersonphonenum"` // 法人电话
	Idnumber            string    `json:"idnumber"`            // 企业法人身份证件号
	Enable              string    `json:"enable"`              // 激活  "start" "close"
	Mainproducts        string    `json:"mainproducts"`        // 主推企业 "yes" "ni"
	Linkman             string    `json:"linkman"`             // 企业 联系人
	Linkmanphonenum     string    `json:"linkmanphonenum"`     // 企业 联系人 电话
	Companyduty         string    `json:"companyduty"`         // 企业 联系人 职务
	Companyname         string    `json:"companyname"`         // 企业名
	Companynum          string    `json:"companynum"`          // 企业人数
	Companynumphonenum  string    `json:"companynumphonenum"`  // 企业电话
	Industrys           []string  `json:"industrys"`           // 公司行业
	Directtag           []string  `json:"directtag"`           // 公司主营范围
	Email               string    `json:"email"`               // 企业邮箱
	Website             string    `json:"website"`             // 企业网站
	Setuptime           time.Time `json:"setuptime"`           // 成立时间
	Partymembernum      string    `json:"partymembernum"`      // 党员人数
	Regiscode           string    `json:"regiscode"`           // 注册码
	Logincapital        string    `json:"logincapital"`        // 注册资金
	Assetcapital        string    `json:"assetcapital"`        // 企业资产
	Annualcapital       string    `json:"annualcapital"`       // 年销售额
	Province            string    `json:"province"`            // 省
	City                string    `json:"city"`                // 市
	District            string    `json:"district"`            // 区
	Address             string    `json:"address"`             // 企业地址
	Addresscoord        []string  `json:"addresscoord"`        // 企业地址 坐标
	Tags                []string  `json:"tags"`                // 标签
	Honor               string    `json:"honor"`               // 企业荣誉
	Honorimg            []string  `json:"honorimg"`            // 企业荣誉图片
	Synopsis            string    `json:"synopsis"`            // 企业简介
	Influence           string    `json:"influence"`           // 企业影响力
	Comments            string    `json:"comments"`            // 企业参与公益事业情况
	Username            string    `json:"username"`            // 用户名称
	License             string    `json:"license"`             // 营业执照
	Clicknum            int       `json:"clicknum"`            // 查看数
	Groupsid            string    `json:"groupsid"`            // 权限
	Integral            int       `json:"integral"`            // 交钱数
	Created             time.Time `json:"created"`
}

func Main_test() {
	filename := "static/xlsx/b17689c172643de3a21829d4da03d76f22.xlsx"
	fmt.Println(filename)
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			var user User
			var company Company
			user.Enable = "yes"
			for k, v := range row.Cells {
				if k <= 13 {
					user = getUser(user, k, v)
				} else {
					company = getCompany(company, k, v)
				}

			}
			fmt.Printf("%+v", user)
			fmt.Printf("%+v", company)
		}
	}
}

func getCompany(company Company, k int, v *xlsx.Cell) Company {
	if v != nil {
		switch k {
		case 14:
			company.Companyname = v.String()
		case 15:
			company.Website = v.String()
		case 16:
			splitval := split(v.String())
			company.Province = splitval[0]
			company.City = splitval[1]
			company.District = splitval[2]
			company.Address = splitval[3]
		case 17:
			company.Setuptime = getTime(v.String())
		case 18:
			company.Companynumphonenum = v.String()
		case 19:
			company.Companynum = v.String()
		case 20:
			company.Partymembernum = v.String()
		case 21:
			company.Legalperson = v.String()
		case 22:
			company.Idnumber = v.String()
		case 23:
			company.Industrys = split(v.String())
		case 24:
			company.Directtag = split(v.String())
		case 25:
			company.Logincapital = v.String()
		case 26:
			company.Assetcapital = v.String()
		case 27:
			company.Linkman = v.String()
		case 28:
			company.Linkmanphonenum = v.String()
		case 29:
			company.Companyduty = v.String()
		case 30:
			company.Comments = v.String()
		case 31:
			company.Regiscode = v.String()
		}
	}
	return company
}

func getUser(user User, k int, v *xlsx.Cell) User {
	if v != nil {
		switch k {
		case 0:
			user.Nickname = v.String()
		case 1:
			user.Phonenum = v.String()
			user.Password = sha512Str(md5Str(v.String()))
		case 2:
			user.Thetitle = v.String()
		case 3:
			user.Birthday = getTime(v.String())
		case 4:
			user.Groupsname = v.String()
		case 5:
			user.Inittime = getTime(v.String())
		case 6:
			user.Gender = getGender(v.String())
		case 7:
			user.Address = v.String()
		case 8:
			user.Duty = v.String()
		case 9:
			user.Email = v.String()
			user.Username = v.String()
		case 10:
			ints64, _ := toInt64Ceil(v.String())
			user.Qq = ints64
		case 11:
			user.Education = v.String()
		case 12:
			user.Tagrollout = split(v.String())
		case 13:
			user.Tagbelow = split(v.String())
		}
	}
	return user
}

func getTime(times string) time.Time {
	formatTime, _ := time.Parse("2006-01-02 15:04:05", times)
	return formatTime
}

func getGender(gender string) int {
	if gender == "男" {
		return 1
	} else {
		return 2
	}
}

//转int64，向上取整
func toInt64Ceil(a interface{}) (int64, error) {
	switch a.(type) {
	case int:
		x, _ := a.(int)
		return int64(x), nil
	case string:
		x, _ := a.(string)
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return int64(0), err
		}
		return int64(math.Ceil(f)), err
	case float32:
		x, _ := a.(float32)
		return int64(math.Ceil(float64(x))), nil
	case float64:
		x, _ := a.(float64)
		return int64(math.Ceil(x)), nil
	case []byte:
		x, _ := a.([]byte)
		bt64 := math.Float64frombits(binary.BigEndian.Uint64(x))
		return int64(math.Ceil(bt64)), nil
	default:
		return int64(0), errors.New("must be base type")
	}
}

//md5验证
func md5Str(src string) string {
	h := md5.New()
	h.Write([]byte(src)) //
	return hex.EncodeToString(h.Sum(nil))
}

//sha512验证
func sha512Str(src string) string {
	h := sha512.New()
	h.Write([]byte(src)) //
	return hex.EncodeToString(h.Sum(nil))
}

func split(splitstr string) []string {
	return strings.Split(splitstr, "/")
}
