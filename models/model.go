package models

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/14 21:40
 **/

//放结构体  表设计
//方法初始化
type user struct {
	id int
	name string
}
func init() {
	//1.连接数据库
	orm.RegisterDataBase("user","mysql","rootroot:root@tcp(127.0.0.1:3306)/db_test?charset=utf8")
	//注册表
	orm.RegisterModel(new(user))
	//生成表
	orm.RunSyncdb("user",false,true)//args 2 强制更新-》false 3 创表视图
	c.Ctx.WriteString("orm创建成功")
}
