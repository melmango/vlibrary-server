package routers

import (
	"github.com/melmango/vlibrary-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
