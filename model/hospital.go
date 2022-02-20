package model

type Hospital struct {
	Id            int `gorm:"primaryKey:autoIncrement"`
	HospitalName  string
	Quota         int
	ReservedQuota int
}
type HospitalListResponse struct {
	Hospitals []Hospital
}
