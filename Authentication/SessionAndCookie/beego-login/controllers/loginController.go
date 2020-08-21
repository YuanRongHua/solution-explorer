// loginController
package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("Hello World!")
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) LoginPage() {

	username := this.GetSession("username")
	if username != nil {
		this.TplName = "success.html"
		beego.Informational("Have login")
		return
	}

	this.TplName = "login.html"
	this.Data["loginUrl"] = "http://localhost:8080/login"
}

func (this *LoginController) Login() {

	this.TplName = "success.html"
	username := this.GetString("username")
	password := this.GetString("password")
	this.SetSession("username", username)
	beego.Informational("username: " + username)
	beego.Informational("password: " + password)
}
