package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDatabase(username, password, dbname, ip, port string) (*gorm.DB, error) {
	return gorm.Open("mysql", dbURL(username, password, dbname, ip, port))
}

func dbURL(username, password, dbname, ip, port string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, ip, port, dbname)
}
