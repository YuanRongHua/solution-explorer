package routers

import (
	"beego-login/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/loginPage", &controllers.LoginController{}, "get:LoginPage")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
}
