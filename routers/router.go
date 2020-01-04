package routers

import (
	"assets-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/assets/save", &controllers.BaseController{}, "post:Save")
	beego.Router("/api/assets/list", &controllers.BaseController{}, "*:List")
}
