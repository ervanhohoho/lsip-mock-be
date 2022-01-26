package model

import "time"

type AccessTime struct {
	Id           int       `gorm:"primaryKey:autoIncrement"`
	LoginTime    time.Time `sql:"type:timestamp without time zone"`
	RegisterTime time.Time `sql:"type:timestamp without time zone"`
}
