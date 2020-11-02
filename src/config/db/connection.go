package db

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var once sync.Once
var db *gorm.DB

func ConnectDB(openCnx string, driver string) *gorm.DB {

	once.Do(func() {
		var err error
		db, err = gorm.Open(driver, openCnx)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("%v", openCnx)
		}
		db.DB().SetMaxIdleConns(1)
		db.DB().SetMaxOpenConns(10)
		db.DB().SetConnMaxLifetime(time.Microsecond * 1)

		db.LogMode(false)
	})

	return db
}
