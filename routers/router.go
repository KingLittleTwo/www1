package routers

import (
	"ww1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
    beego.Router("/login", &controllers.AuthController{}, "get:Login")
}
