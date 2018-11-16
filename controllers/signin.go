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
	password := c.Ctx.Request.FormValue("password")
	correctPwd := ""
	for _, user := range models.Users {
		if user.Username == username {
			correctPwd = user.Password
			break
		}
	}
	if correctPwd == password {
		c.TplName="detail.jade"
		c.Ctx.Redirect(302, "/detail/"+username)
	} else if correctPwd == "" {
		c.Data["user"] = models.User{}
		c.Data["error"] = username + "用户未注册"
		c.TplName = "signin.jade"
	} else {
		c.Data["user"] = models.User{}
		c.Data["error"] = "密码错误，请重新登录"
		c.TplName = "signin.jade"
	}
}
