package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"os"
)

/**
微信摇一摇
/lucky 只有一个抽奖的接口
*/

//奖品类型
const (
	giftTypeCoin      = iota //虚拟币
	giftTypeCoupon           //不同券
	giftTypeCouponFix        //相同的券
	giftTypeRealSmall        //实物小奖
	giftTypeRealLarge        //实物大奖
)

type gift struct {
	id       int      //奖品ID
	name     string   //奖品名称
	pic      string   //奖品图片
	link     string   //奖品链接
	gtype    int      //奖品类型
	data     string   //奖品的数据（特定的配置信息）
	datalist []string //奖品数据集合（不同的优惠券的编码）
	total    int      //总数， 0 不限量
	left     int      //剩余数量
	inuse    bool     //是否使用中
	rate     int      //中奖概率， 万分之N, 0-9999
	rateMin  int      //大于等于中奖编码
	rateMax  int      //小于中奖编码
}

// 最大的中奖号码
const rateMax = 10000

//初始化日志
var logger *log.Logger

func initLog() {
	f, _ := os.Create("/var/log/lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

// 奖品列表
var giftList []*gift

type lotteryController struct {
	Ctx iris.Context
}

//设置奖品列表
func initGift() {
	giftList = make([]*gift, 5)
	g1 := gift{
		id:   1,
		name: "手机大奖",
		pic:  "",
	}
	giftList = append(giftList, &g1)
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	initLog()
	return app
}

func main() {

}
