package controllers

import (
	"github.com/astaxie/beego"
	"github.com/liuyh73/cloudgo/models"
)

type DetailController struct {
	beego.Controller
}

func (c *DetailController) Get() {
	username := c.Ctx.Input.Param(":name")
	var signinUser models.User
	for _, user := range models.Users {
		if user.Username == username {
			signinUser = models.User{
				Username:  user.Username,
				Password:  user.Password,
				Email:     user.Email,
				Telephone: user.Telephone,
				Id:        user.Id,
			}
			break
		}
	}
	if signinUser.Username == "" {
		c.Data["message"] = username + "用户未注册"
		c.TplName = "error.jade"
		//c.Ctx.Redirect(302, "/error")
	} else {
		c.Data["user"] = signinUser
		c.TplName = "detail.jade"
	}
}
