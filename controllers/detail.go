package controllers

import (
	"github.com/astaxie/beego"
	"github.com/liuyh73/cloudgo/models"
)

type DetailController struct {
	beego.Controller
}

func (c *DetailController) Jade() {
	user := models.User{}
	c.Data[user] = models.User{
		Username:  "test",
		Password:  "123",
		Email:     "sd",
		Telephone: "15989067460",
		Id:        "12345678",
	}
	c.TplName = "detail.jade"
}
