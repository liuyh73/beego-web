package routers

import (
	"github.com/astaxie/beego"
	"github.com/liuyh73/cloudgo/controllers"
)

func init() {
	beego.Router("/beego", &controllers.MainController{})
	beego.Router("/detail/:name", &controllers.DetailController{}, "get:Get")
	beego.Router("/signin", &controllers.SigninController{},"get:Get;post:Post")
	beego.Router("/signup", &controllers.SignupController{}, "get:Get;post:Post")
}
