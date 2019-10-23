package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type OperationController struct {
	beego.Controller
}

type Useroper struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (c *OperationController) GetOperation() {
	fmt.Println(c.Ctx.Input.Param(":id"))
	fmt.Printf("%+v", c.Ctx)
	c.Data["json"] = "GetOperation"
	c.ServeJSON()
}
func (c *OperationController) PostOperation() {
	fmt.Println("=====PostOperation======")
	useroper := new(Useroper)
	buf := make([]byte, 1024)
	n, _ := c.Ctx.Request.Body.Read(buf)
	json.Unmarshal([]byte(string(buf[0:n])), &useroper)
	fmt.Printf("%+v", useroper)
	c.Data["json"] = "PostOperation"
	c.ServeJSON()
}

func (c *OperationController) PutOperation() {
	fmt.Println("=====PutOperation======")
	fmt.Println(c.Ctx.Input.Param(":id"))
	useroper := new(Useroper)
	buf := make([]byte, 1024)
	n, _ := c.Ctx.Request.Body.Read(buf)
	json.Unmarshal([]byte(string(buf[0:n])), &useroper)
	fmt.Printf("%+v", useroper)
	c.Data["json"] = "PutOperation"
	c.ServeJSON()
}

func (c *OperationController) DeleteOperation() {
	fmt.Println("=====DeleteOperation======")
	fmt.Println(c.Ctx.Input.Param(":id"))
	fmt.Printf("%+v", c.Ctx.Request)
	c.Data["json"] = "DeleteOperation"
	c.ServeJSON()
}

type Delarr struct {
	Delarr []string `json:"delarr"`
}

func (c *OperationController) DeletearrOperation() {
	fmt.Println("=====DeletearrOperation======")
	delarr := new(Delarr)
	buf := make([]byte, 1024)
	n, _ := c.Ctx.Request.Body.Read(buf)
	json.Unmarshal([]byte(string(buf[0:n])), &delarr)
	fmt.Println(delarr.Delarr)
	c.Data["json"] = "DeletearrOperation"
	c.ServeJSON()
}
