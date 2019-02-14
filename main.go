package main

import (
	_ "wstshop/routers"
	_ "wstshop/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

