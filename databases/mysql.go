package databases

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	once = sync.Once{}
	DB   *gorm.DB
)

func GetCon(tenant string) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), tenant)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = db

	return DB
}
