package accessor

import (
	"fmt"
	"github.com/ervanhohoho/lsip-mock-be/model"
	"github.com/ervanhohoho/lsip-mock-be/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Accessor struct {
	db *gorm.DB
	km util.KeyedMutex
}

func Initialize(host string, dbname string, user string, pw string) Accessor {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pw, host, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		fmt.Println("Error Trying to connect to mysql")
	}
	return Accessor{
		db: db,
		km: util.Initialize(),
	}
}
func (a *Accessor) GetHospitals() []model.Hospital {
	var hospitals []model.Hospital
	result := a.db.Find(&hospitals)
	if result.Error == nil {
		return hospitals
	} else {
		fmt.Printf("Error: %+v", result.Error)
		return []model.Hospital{}
	}
}
func (a *Accessor) ReserveHospital(entity model.Hospital) (bool, string) {
	//TODO Implement Locking to reserve hospital
	var dbModel model.Hospital
	lockName := fmt.Sprintf("HOSPITAL:%s", entity.Id)
	success := false
	errMsg := ""
	unlock := a.km.Lock(lockName)
	a.db.Find(&dbModel, entity.Id)
	if dbModel.ReservedQuota < dbModel.Quota {
		dbModel.ReservedQuota = dbModel.ReservedQuota + 1
		result := a.db.Save(dbModel)
		if result.Error != nil {
			success = true
		} else {
			errMsg = result.Error.Error()
		}
	} else {
		errMsg = "Kuota sudah habis"
	}
	unlock()
	return success, errMsg
}

func (a *Accessor) UpdateHospital(entities []model.Hospital) bool {
	//TODO Update Hospital here
	var insertRes, updateRes *gorm.DB
	insert, update := separateInsertAndUpdate(entities)
	insertSuccess := true
	updateSuccess := true
	if len(insert) > 0 {
		insertRes = a.db.Create(&insert)
		if insertRes.Error != nil {
			insertSuccess = false
		}
	}
	if len(update) > 0 {
		for _, element := range update {
			updateRes = a.db.Save(&element)
			if updateRes.Error != nil {
				updateSuccess = false
			}
		}
	}
	if insertSuccess && updateSuccess {
		return true
	} else {
		return false
	}
}
func (a *Accessor) GetAccessTimes() (model.AccessTime, *error) {
	var ret model.AccessTime
	result := a.db.Last(&ret)
	return ret, &result.Error
}
func (a *Accessor) UpdateAccessTimes(time model.AccessTime) (model.AccessTime, *error) {
	var ret model.AccessTime
	a.db.Last(&ret)
	ret.LoginTime = time.LoginTime
	ret.RegisterTime = time.RegisterTime
	result := a.db.Save(&ret)
	return ret, &result.Error
}
func separateInsertAndUpdate(entities []model.Hospital) (insert []model.Hospital, update []model.Hospital) {
	update = []model.Hospital{}
	insert = []model.Hospital{}
	for _, element := range entities {
		if element.Id != 0 {
			update = append(update, element)
		} else {
			insert = append(insert, element)
		}
	}
	return insert, update
}
