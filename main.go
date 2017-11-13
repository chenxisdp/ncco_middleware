package main

import (
	"github.com/astaxie/beego"
	"github.com/ncco_middleware/controllers"
)

func main() {
	beego.Router("/ncco", &controllers.NccoController{})
	beego.Run()

}
