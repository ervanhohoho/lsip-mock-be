package accessor

import (
	"fmt"
	"github.com/ervanhohoho/lsip-mock-be/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type accessor struct {
	db *gorm.DB
}

func Initialize(host string, dbname string, user string, pw string) accessor {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pw, host, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		fmt.Println("Error Trying to connect to mysql")
	}
	return accessor{
		db: db,
	}
}
func (a *accessor) GetHospitals() []model.Hospital {
	var hospitals []model.Hospital
	result := a.db.Find(&hospitals)
	if result.Error == nil {
		return hospitals
	} else {
		fmt.Printf("Error: %+v", result.Error)
		return []model.Hospital{}
	}
}
