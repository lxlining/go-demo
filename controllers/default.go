package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	//传输数据
	c.Data["address"]="重庆"
	c.TplName = "index.tpl"  //  test.html
}

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Post()  {
	this.Data["address"]=" go  web "
	this.TplName="test.html"
}

func (this * IndexController) ShowFunc() {
	id:=this.GetString(":id")
	/*path:=this.GetString(":path")//文件名
	ext:=this.GetString(":ext")//后缀
	splat:=this.GetString(":splat")*/
	beego.Info(id)
	this.Data["test"]="这是get请求对应方法"
	this.TplName="test.html"
}
