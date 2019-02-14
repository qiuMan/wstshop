package routers

import (
	"wstmart/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //user
    beego.Router("/users/index", &controllers.Users{}, "Get,Post:Index")
    beego.Router("/users/?:id", &controllers.Users{})

    //菜单管理
    beego.Router("/menus/index", &controllers.Menus{}, "Get,Post:Index")

    //权限管理
    beego.Router("/privileges/index/:id", &controllers.Privileges{}, "Get,Post:Index")
    beego.Router("/privileges/:id", &controllers.Privileges{})
}
