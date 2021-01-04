package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBConn() *gorm.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "simplechat"
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
