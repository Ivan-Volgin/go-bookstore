package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "")   // в mysql оказывается нельзя зарегистрироваться, поэтому работоспособность 
	if err != nil {					   // никак не проверить. Позже можно попытаться прикрутить PostgresSQL
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}