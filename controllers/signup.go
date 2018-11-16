package controllers

import (
	"github.com/astaxie/beego"
	"github.com/liuyh73/cloudgo/models"
)

type SignupController struct {
	beego.Controller
}

func (c *SignupController) Get() {
	c.Data["user"] = models.User{
		Username:  "test",
		Password:  "123",
		Email:     "sd",
		Telephone: "15989067460",
		Id:        "12345678",
	}
	c.TplName = "detail.jade"
}
