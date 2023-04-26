package controllers

import (
	"loveHome/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AreaCotroller struct {
	beego.Controller
}

func (this *AreaCotroller) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaCotroller) GetArea() {
	beego.Info("connect success")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer c.RetData(resp)

	//从session拿数据

	//从mysql拿到area数据
	var areas []models.Area

	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)

	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		c.RetData(resp)
		return
	}

	if num == 0 {
		beego.Info("数据错误")
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = areas
	//打包成json返回给前端

	beego.Info("query data success, resp =", resp)
}
