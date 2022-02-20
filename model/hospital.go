package model

type Hospital struct {
	Id            int    `gorm:"primaryKey:autoIncrement" json:"id"`
	HospitalName  string `json:"hospitalName"`
	Quota         int    `json:"quota"`
	ReservedQuota int    `json:"reservedQuota"`
}
type HospitalListResponse struct {
	Hospitals []Hospital `json:"hospitals"`
}
