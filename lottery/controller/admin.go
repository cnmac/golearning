package controller

import (
	services "github.com/cnmac/golearning/lottery/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserDay services.UserdayService
	ServiceBlackip services.BlackipService
}

func (c AdminController) Get() mvc.Result {
	return mvc.View{
		Name:   "admin/index.html",
		Layout: "admin/layout.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "",
		},
		Code: 0,
		Err:  nil,
	}
}
