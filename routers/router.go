package routers

import (
	"github.com/astaxie/beego"
	"github.com/liuyh73/cloudgo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/detail", &controllers.DetailController{}, "Get:Jade")
}
