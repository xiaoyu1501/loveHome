package controllers

import (
	"loveHome/models"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
)

type SessionCotroller struct {
	beego.Controller
}

func (this *SessionCotroller) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *SessionCotroller) GetSessionData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	user := models.User{}
	/* 	   	user.Name = "wyj"

	   	   	resp["errno"] = 0
	   	   	resp["errmsg"] = "ok" */
	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

	name := this.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}
}

func (this *SessionCotroller) DeleteSessionData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("name")

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

}
