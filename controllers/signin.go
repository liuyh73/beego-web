package controllers

import (
	"github.com/astaxie/beego"
	"github.com/liuyh73/cloudgo/models"
)

type SigninController struct {
	beego.Controller
}

func (c *SigninController) Get() {
	c.Data["user"] = models.User{}
	c.Data["error"] = ""
	c.TplName="signin.jade"
}

func (c *SigninController) Post() {
	username := c.Ctx.Request.FormValue("username")
	c.TplName="detail.jade"
	c.Ctx.Redirect(302, "/detail/"+username)
}
