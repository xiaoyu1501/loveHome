package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/adapter/orm"
)

type AreaCotroller struct {
	beego.Controller
}

func (c *AreaCotroller) GetArea() {
	beego.Info("connect success")
	//从session拿数据

	//从mysql拿到area数据
	area := models.area{}

	o := orm.NewOrm()
	o.Read
	//打包成json返回给前端

}
