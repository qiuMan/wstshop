/*
 * @Author: v_wenlqiu 
 * @Date: 2019-02-02 11:55:24 
 * @Last Modified by: v_wenlqiu
 * @Last Modified time: 2019-02-04 00:46:26
 */
package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"wstmart/models"
)

type Base struct {
	beego.Controller
	controllerName string             //当前控制名称
	actionName     string             //当前action名称
	curUser        models.Users //当前用户信息
	Login		   bool			//是否是登录
}

func (this *Base) Prepare() {
	//模板后者设置
	this.TplExt = "html"
	
	//附值
	this.controllerName, this.actionName = this.GetControllerAndAction()
	//从Session里获取数据 设置用户信息
	this.adapterUserInfo()
	this.checkLogin()
	this.Login = true
	
	if this.Login == true && !this.Ctx.Input.IsAjax() {
		fmt.Println("Base/Prepare")
		this.SetTpl()
	}
}

//从session里取用户信息
func (this *Base) adapterUserInfo() {
	a := this.GetSession("userInfo")
	
	if a != nil {
		this.curUser = a.(models.Users)
		this.Data["userInfo"] = a
	}
}

func (this *Base) checkLogin() {
	if this.curUser.Id == 0 {
		ctrls := this.ctrls()
		if !this.inArray(this.controllerName, ctrls) {
			// this.Redirect("/", 302)
		}
	}else {
		this.Login = true
	}
}

func (this *Base) SetTpl() {	
	this.Layout = "layout/index.html"
}

func (this *Base) ctrls() []string {
	 var ctrls []string
	 ctrls = append(ctrls, "MainController")
	 return ctrls
}

func (this *Base) inArray(str string, array []string) bool {
	size := len(array)

	for i := 0; i < size; i++ {
		if array[i] == str {
			return true
		}
	}  

	return false
}

func (this *Base) JsonOut(Params ...interface{}) {
	json := make(map[string]interface{})

	json["code"] = 0
	json["msg"] = "success"
	json["count"] = 0
	json["data"] = make(map[string]interface{})

	for key, value := range Params {
		switch key {
		case 0:
			json["data"] = value
		case 1:
			json["count"] = value
		case 2:
			json["code"] = value
		case 3:
			json["msg"] = value
			
		}
	} 
	
	this.Data["json"] = json

	this.ServeJSON()

}



func Echo(value ...interface{}) {
	fmt.Println(value)
}
