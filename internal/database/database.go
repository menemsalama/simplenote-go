package database

import (
	"github.com/jinzhu/gorm"
	// use postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PG postgres db struct
var PG *gorm.DB

// ConnectToPG opens postgres connection
func ConnectToPG(url string) (err error) {
	PG, err = gorm.Open("postgres", url)
	return
}
