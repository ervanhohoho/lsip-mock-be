package accessor

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Initialize(host string, dbname string, user string, pw string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pw, host, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error Trying to connect to mysql")
	}
	return db
}
func Test() {
	fmt.Println("Test")
}
