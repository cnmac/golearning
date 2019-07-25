package dao

import (
	"github.com/cnmac/golearning/lottery/models"
	"github.com/go-xorm/xorm"
	"log"
)

type GiftDao struct {
	engine *xorm.Engine
}

func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{engine}
}

func (d GiftDao) Get(id int) *models.LtGift {
	gift := &models.LtGift{Id: id}
	ok, e := d.engine.Get(gift)
	if ok && e == nil {
		return gift
	} else {
		gift.Id = 0
		return gift
	}
}

func (d GiftDao) GetAll() []models.LtGift {
	gifts := make([]models.LtGift, 0)
	err := d.engine.Asc("sys_status", "displayorder").Find(&gifts)
	if err != nil {
		log.Println("gift_dao.GetAll error=", err)
		return gifts
	}
	return gifts
}

func (d GiftDao) CountAll() int64 {
	num, e := d.engine.Count(&models.LtGift{})
	if e != nil {
		return 0
	} else {
		return num
	}
}

func (d GiftDao) Delete(id int) error {
	gift := &models.LtGift{Id: id, SysStatus: 1}
	_, e := d.engine.Id(id).Update(gift)
	return e
}

func (d GiftDao) Update(data *models.LtGift, columns []string) error {
	_, e := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return e
}

func (d GiftDao) Create(gift *models.LtGift) error {
	_, e := d.engine.Insert(gift)
	return e
}
