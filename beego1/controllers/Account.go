package controllers

import (
	"github.com/astaxie/beego"
)

type AccountController struct {
	beego.Controller
}

func (c *AccountController) Login() {
    p0 := c.Ctx.Input.Param("0")
	p1 := c.Ctx.Input.Param("2")
	c.Data["Website"] = p0
	c.Data["Email"] = p1
	c.TplNames = "index.tpl"
}

func (c *AccountController) Logout() {
	c.Data["Website"] = "logout"
	c.Data["Email"] = "logout@gmail.com"
	c.TplNames = "index.tpl"
}