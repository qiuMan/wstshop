package controllers

import (
	"fmt"
	"crypto/md5"
	"wstmart/models"

)



type MainController struct {
	Base
}

func (this *MainController) Get() {	
	this.login()
}

func (this *MainController) Post() {
	user := models.NewUsers()
	userName := this.GetString("username")
	u, err := user.GetUserInfoByName(userName)
	pwd := this.GetString("password")
	data := []byte(pwd)
	has := md5.Sum(data)
	password := fmt.Sprintf("%x", has)
	
	if err ==nil && u.LoginPwd != password {
		this.SetSession("userInfo", u)
		this.Login = true
	} else {
		this.Data["username"] = userName
	}

	this.login()  
}

func (this *MainController) login() {
	if this.Login == true {
		this.Redirect("user/index", 302)
	} else {
		this.TplName = "user/login.html"
	}
}

