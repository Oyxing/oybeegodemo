package models

import (
	"fmt"

	"github.com/go-xorm/xorm"

	_ "github.com/mattn/go-sqlite3"
)

var db *xorm.Engine

func init() {
	var err error
	db, err = xorm.NewEngine("sqlite3", "./data/exdemo.db")
	if err != nil {
		fmt.Println(err)
	}
	err = db.Sync(new(User), new(Stats), new(Mode))
	if err != nil {
		return
	}
}

func GetDb() *xorm.Engine {

	return db
}
