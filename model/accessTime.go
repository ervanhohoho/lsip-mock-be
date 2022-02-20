package model

import "time"

type AccessTime struct {
	Id           int       `gorm:"primaryKey:autoIncrement" json:"id"`
	LoginTime    time.Time `sql:"type:timestamp without time zone" json:"loginTime""`
	RegisterTime time.Time `sql:"type:timestamp without time zone" json:"registerTime"`
}
