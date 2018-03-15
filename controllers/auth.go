package controllers

import (
	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) Login() {
	//if this.Input().
	this.Layout = "layout.html"
	this.TplName = "signin.html"
}
