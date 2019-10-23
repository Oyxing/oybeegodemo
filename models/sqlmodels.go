package models

type User struct {
	Id       int    `json:"id" xorm:"pk int autoincr notnull unique 'id'"`
	Name     string `json:"name" xorm:"notnull"`
	Age      int64  `json:"age" xorm:"notnull"`
	Pwd      string `json:"pwd" xorm:"notnull"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mode     []Mode `json:"mode" xorm:"notnull"`
}
type Mode struct {
	Id   int    `json:"id" xorm:"pk int autoincr notnull unique 'id'"`
	Uuid string `json:"uuid" xorm:"notnull"`
	Msg  string `json:"msg" xorm:"notnull"`
	Code int64  `json:"code" xorm:"notnull"`
}

// 状态数据库

type Stats struct {
	Id                   int    `json:"id" xorm:"pk int autoincr notnull unique 'id'"`
	Name                 string `json:"name" xorm:"notnull"`
	Containerid          string `json:"containerid" xorm:"notnull"`
	SystemCpuUsage       int64  `json:"system_cpu_usage" xorm:"notnull"`
	OnlineCpus           int64  `json:"online_cpus" xorm:"notnull"`
	MemUsage             int64  `json:"mem_usage" xorm:"notnull"`
	MemMaxUsage          int64  `json:"mem_max_usage" xorm:"notnull"`
	MemLimit             int64  `json:"mem_limit" xorm:"notnull"`
	CpuTotalUsage        int64  `json:"cpu_total_usage" xorm:"notnull"`
	CpuPerUsageNum       int64  `json:"cpu_per_usage_num" xorm:"notnull"`
	CpuUsageInKernelmode int64  `json:"cpu_usage_in_kernelmode" xorm:"notnull"`
	CpuUsageInUsermode   int64  `json:"cpu_usage_in_usermode" xorm:"notnull"`
	Created              int64  `json:"created" xorm:"notnull"`
}

type JsonInfo struct {
	Id          int      `json:"id,omitempty"`
	Addr        string   `json:"addr,omitempty"`
	Containerid string   `json:"containerid,omitempty"`
	Port        string   `json:"port,omitempty"`
	Addrs       []string `json:"addrs,omitempty"`
	Syncbool    bool     `json:"syncbool,omitempty"`
	Nickname    string   `json:"nickname,omitempty"`
	Source      string   `json:"source,omitempty"`
	Target      string   `json:"target,omitempty"`
	Imagename   string   `json:"imagename,omitempty"`
	Password    string   `json:"password,omitempty"`
	Username    string   `json:"username,omitempty"`
	Serveraddr  string   `json:"serveraddr,omitempty"`
	Networkname string   `json:"networkname,omitempty"`
	Netid       string   `json:"netid,omitempty"`
	Repository  string   `json:"repository,omitempty"`
	Tag         string   `json:"tag,omitemptyg"`
	Allbool     string   `json:"allbool,omitempty"`
	Newname     string   `json:"newname,omitempty"`
	Name        string   `json:"name,omitempty"`
}

type Ordermsg struct {
	Id         string `json:"id" xorm:"pk notnull unique 'id'"`
	Productid  int    `json:"productid"`
	Outtradeno int    `json:"outtradeno"`
	Totalfee   int    `json:"totalfee"`
	Body       string `json:"body"`
	Attach     string `json:"attach"`
}
