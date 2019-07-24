package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)
import "github.com/go-xorm/xorm"

type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"`
}

func main() {
	x, err := xorm.NewEngine("mysql", "root:111111@/sys?charset=utf8")
	if err = x.Sync2(new(Account)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}
