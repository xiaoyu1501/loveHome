package controllers

import (
	"encoding/json"
	"loveHome/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserCotroller struct {
	beego.Controller
}

func (this *UserCotroller) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UserCotroller) Reg() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	//获取前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)

	/*mobile: "111"
	     password: "111"
		 sms_code: "111"


	beego.Info(`resp["mobile"] =`, resp["mobile"])
	beego.Info(`resp["password"] =`, resp["mobile"])
	beego.Info(`resp["sms_code"] =`, resp["mobile"])
	*/

	//插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Password_hash = resp["password"].(string)
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)

	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	beego.Info("reg success ,id =", id)
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	this.SetSession("name", user.Name)
}
