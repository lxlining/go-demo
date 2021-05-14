package routers

import (
	"demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/index", &controllers.IndexController{},"get:ShowFunc;post:Post")//高级路由
	//正则路由
	beego.Router("/index/?:id", &controllers.IndexController{}) //  :id将数据赋值给id 为空 ？：id 可以
	beego.Router("/index/*.*", &controllers.IndexController{})

    //数据库
	beego.Router("/mysql", &controllers.MysqlController{},"get:showSql")
	beego.Router("/orml", &controllers.OrmController{},"get:showSql")
	beego.Router("/orml", &controllers.OrmController{},"get:showInsert")
	beego.Router("/orml", &controllers.OrmController{},"get:showList")
	beego.Router("/orml", &controllers.OrmController{},"get:showUpdate")
}
