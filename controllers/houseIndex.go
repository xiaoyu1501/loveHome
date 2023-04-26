package controllers

import (
	"loveHome/models"
	_ "loveHome/models"

	_ "github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

type HouseIndexCotroller struct {
	beego.Controller
}

func (this *HouseIndexCotroller) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *HouseIndexCotroller) GetHouseIndex() {
	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_DATAERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
	this.RetData(resp)
}
