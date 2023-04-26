package routers

import (
	"loveHome/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaCotroller{}, "get:GetArea")
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexCotroller{}, "get:GetHouseIndex")
	//api/v1.0/session
	beego.Router("/api/v1.0/session", &controllers.SessionCotroller{}, "get:GetSessionData;delete:DeleteSessionData")
	//api/v.0/users
	beego.Router("/api/v1.0/users", &controllers.UserCotroller{}, "post:Reg")
}
