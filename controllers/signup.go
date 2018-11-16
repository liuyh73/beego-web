package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/liuyh73/cloudgo/models"
)

type SignupController struct {
	beego.Controller
}

func (c *SignupController) Get() {
	fmt.Println("signup get")
	c.Data["user"] = models.User{}
	c.TplName = "signup.jade"
}

func (c *SignupController) Post() {
	signupUser := models.User {
		Username: c.Ctx.Request.FormValue("username"),
		Password: c.Ctx.Request.FormValue("password"),
		Id: c.Ctx.Request.FormValue("id"),
		Email: c.Ctx.Request.FormValue("email"),
		Telephone: c.Ctx.Request.FormValue("telephone"),
	}

	for _, user := range models.Users {
		if user.Username == signupUser.Username {
			c.Data["error"] = "用户名已被注册"
			c.Data["user"] = models.User{}
			c.TplName = "signup.jade"
			return
		}
	}
	models.Users = append(models.Users, signupUser) 
	c.Data["user"] = signupUser
	c.TplName = "detail.jade"
	c.Ctx.Redirect(302, "/detail/"+c.Ctx.Request.FormValue("username"))
}