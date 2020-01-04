package main

import (
	_ "assets-api/routers"
	_ "assets-api/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

