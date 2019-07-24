package main

import _ "github.com/go-sql-driver/mysql"
import  "github.com/go-xorm/xorm"




type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"`
}

func main() {
	x,err:=.NewEngine("mysql", "root:111111@/sys?charset=utf8")
}

