package controllers

import (
	"github.com/liuyh73/cloudgo/models"
	"github.com/astaxie/beego"
	_ "github.com/Joker/jade"
)

type DetailController struct {
	beego.Controller
}

func (c *DetailController) Jade() {
	user := models.User{}
	c.Data[user] = models.User{
		Username: "test",
		Password: "123",
		Email: "sd",
		Telephone: "15989067460",
	}
	c.TplName = "detail.jade"
}
