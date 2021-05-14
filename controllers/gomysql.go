package controllers

import (
	"database/sql"
	"fmt"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/14 21:14
 **/
type MysqlController struct {
	beego.Controller
}
//数据库
func (this * MysqlController) showSql() {
	conn,err:=sql.Open("mysql","rootroot:root@tcp(127.0.0.1:3306)/db_test?charset=utf8")
	if err != nil {
		beego.Info(err)
		return
	}
	defer conn.Close()
	//创建数据表
	/*res,err:=conn.Exec("create table userInfo(id int,name varchar(20))")
	if err != nil {
		beego.Info(err)
		return
	}
	this.Ctx.WriteString("创建表成功！")
	beego.Info(res)*/
	//插入
	/*res,err:=conn.Exec("insert userInfo(id,name)values(?,?)",1,"bb")
	if err != nil {
		beego.Info(err)
		return
	}
	this.Ctx.WriteString("插入表成功！")
	beego.Info(res)*/
	//查询
	rows,err:=conn.Query("select * from userInfo")
	var id int
	var name string
	for rows.Next(){
		rows.Scan(&id,&name)
	}

}
