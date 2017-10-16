package routers

import (
	"github.com/lflxp/ips/controllers"
	"github.com/astaxie/beego"
)

func init() {
    	beego.Router("/", &controllers.MainController{})
    	beego.Router("/check/?:ip", &controllers.MainController{},"get,post:Check")
    	beego.Router("/debug/pprof", &controllers.ProfController{})
    	beego.Router("/debug/pprof/:pp([\\w]+)", &controllers.ProfController{})
}