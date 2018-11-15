package routers

import (
	"github.com/liuyh73/cloudgo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
