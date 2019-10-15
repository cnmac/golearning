package controller

import (
	"encoding/json"
	"fmt"
	"github.com/cnmac/golearning/web/comm"
	"github.com/cnmac/golearning/web/models"
	services "github.com/cnmac/golearning/web/service"
	"github.com/cnmac/golearning/web/viewmodels"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"time"
)

type AdminGiftController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserDay services.UserdayService
	ServiceBlackip services.BlackipService
}

func (c AdminGiftController) Get() mvc.Result {
	dataList := c.ServiceGift.GetAll(false)
	//total := len(dataList)
	for i, giftInfo := range dataList {
		prizedata := make([][2]int, 0)
		err := json.Unmarshal([]byte(giftInfo.PrizeData), &prizedata)
		if err != nil || len(prizedata) < 1 {
			dataList[i].PrizeData = "[]"
		} else {
			newpd := make([]string, len(prizedata))
			for index, pd := range prizedata {
				ct := comm.FormatFromUnixTime(int64(pd[0]))
				newpd[index] = fmt.Sprintf("[%s] : %d", ct, pd[1])
			}
			str, err := json.Marshal(newpd)
			if err != nil && len(str) > 0 {
				dataList[i].PrizeData = string(str)
			} else {
				dataList[i].PrizeData = "[]"
			}
		}
	}
	return mvc.View{
		Name:   "admin/gift.html",
		Layout: "admin/layout.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "gift",
		},
		Code: 0,
		Err:  nil,
	}
}

func (c AdminGiftController) GetEdit() mvc.Result {
	id := c.Ctx.URLParamIntDefault("id", 0)
	giftInfo := viewmodels.ViewGift{}
	if id > 0 {
		data := c.ServiceGift.Get(id, false)
		giftInfo.Id = data.Id
		giftInfo.Title = data.Title
		giftInfo.PrizeNum = data.PrizeNum
		giftInfo.PrizeTime = data.PrizeTime
		giftInfo.PrizeCode = data.PrizeCode
		giftInfo.Displayorder = data.Displayorder
		giftInfo.Gdata = data.Gdata
		giftInfo.Gtype = data.Gtype
		giftInfo.Img = data.Img
		giftInfo.TimeBegin = comm.FormatFromUnixTime(int64(data.TimeBegin))
		giftInfo.TimeEnd = comm.FormatFromUnixTime(int64(data.TimeEnd))
	}
	return mvc.View{
		Name:   "admin/giftEdit.html",
		Layout: "admin/layout.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "gift",
			"info":    giftInfo,
		},
	}
}

func (c AdminGiftController) PostSave() mvc.Result {
	data := viewmodels.ViewGift{}
	err := c.Ctx.ReadForm(&data)
	if err != nil {
		fmt.Println("admin_gift.PostSave ReadForm error = ", err)
		return mvc.Response{
			Text: fmt.Sprintf("ReadForm转换异常", err),
		}
	}
	gift := models.LtGift{}
	gift.Id = data.Id
	gift.Title = data.Title
	gift.PrizeCode = data.PrizeCode
	gift.PrizeTime = data.PrizeTime
	gift.PrizeNum = data.PrizeNum
	gift.Img = data.Img
	gift.Gtype = data.Gtype
	gift.Gdata = data.Gdata
	gift.Displayorder = data.Displayorder
	t1, err1 := comm.ParseTime(data.TimeBegin)
	t2, err2 := comm.ParseTime(data.TimeEnd)
	if err1 != nil || err2 != nil {
		return mvc.Response{
			Text: fmt.Sprintf("时间格式不正确， err1 = %s, err2 = %s", err1, err2),
		}
	}
	gift.TimeBegin = int(t1.Unix())
	gift.TimeEnd = int(t2.Unix())
	if gift.Id > 0 {
		getGift := c.ServiceGift.Get(gift.Id, false)
		if getGift != nil && getGift.Id > 0 {
			if getGift.PrizeNum != gift.PrizeNum {
				gift.LeftNum = gift.LeftNum - getGift.PrizeNum - gift.PrizeNum
				if gift.LeftNum < 0 || gift.PrizeNum >= 0 {
					gift.LeftNum = 0
				}
			}
			gift.SysUpdated = int(time.Now().Unix())
			c.ServiceGift.Update(&gift, []string{""})
		} else {
			gift.Id = 0
		}
	}
	if gift.Id == 0 {
		gift.LeftNum = gift.PrizeNum
		gift.SysIp = comm.ClientIP(c.Ctx.Request())
		gift.SysCreated = int(time.Now().Unix())
		c.ServiceGift.Create(&gift)
	}
	return mvc.Response{
		Path: "/admin/gift",
	}
}

func (c AdminGiftController) GetDelete() mvc.Result {
	//TODO
	return mvc.Response{
		Path: "/admin/gift",
	}
}

func (c AdminGiftController) GetReset() mvc.Result {
	//TODO
	return mvc.Response{
		Path: "/admin/gift",
	}
}
