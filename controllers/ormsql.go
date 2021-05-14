package controllers

import "demo/models"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/14 21:40
 **/

type OrmController struct {
	beego.Controller
}

type user struct {
	id int
	name string
}
func (c * OrmController)showSql() {
	//1.连接数据库
	orm.RegisterDataBase("user","mysql","rootroot:root@tcp(127.0.0.1:3306)/db_test?charset=utf8")
	//注册表
	orm.RegisterModel(new(user))
	//生成表
	orm.RunSyncdb("user",false,true)//args 2 强制更新-》false 3 创表视图
	c.Ctx.WriteString("orm创建成功")
}

//增
func (c * OrmController)showInsert() {
	o:=orm.NewOrm()
	user:=user{}
	//user:=models.user{}
	user.name="aaa"
	user.id=1001
	n,err:=o.Insert(&user)
	if err != nil {
		beego.Info("插入失败",err)
		return
	}
	c.Ctx.WriteString("orm插入成功")
}

//查
func (c * OrmController)showList() {
	o:=orm.NewOrm()
	user:=user{id:1}
	//user1:=models.user{id:1}
	err:=o.Read(&user)
	if err != nil {
		beego.Info("查询失败",err)
		return
	}
	beego.Info(user.name)
	c.Ctx.WriteString("orm查询成功")
}
//改
func (c * OrmController)showUpdate() {
	o:=orm.NewOrm()
	user:=models.user{id:1}
	//user1:=models.user{id:1}
	err:=o.Read(&user)
	if err != nil {
		beego.Info("更新失败",err)
		return
	}
	user.name="bbbb"
	_,err=o.Update(&user)
	if err != nil {
		beego.Info("更新失败",err)
		return
	}

	c.Ctx.WriteString("orm更新成功")
}