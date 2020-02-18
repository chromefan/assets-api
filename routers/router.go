package routers

import (
	"assets-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/assets/save", &controllers.BaseController{}, "post:Save")
	beego.Router("/api/assets/list", &controllers.BaseController{}, "*:List")
	beego.Router("/api/work/day/save", &controllers.BaseController{}, "post:WorkDaySave")
	beego.Router("/api/work/group/save", &controllers.BaseController{}, "post:WorkGroupSave")
	beego.Router("/api/work/project/save", &controllers.BaseController{}, "post:WorkProjectSave")
	beego.Router("/api/work/list", &controllers.BaseController{}, "*:WorkList")
}
