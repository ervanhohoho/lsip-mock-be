package payload

import "github.com/ervanhohoho/lsip-mock-be/model"

type UpdateHospitalPayload struct {
	Entities []model.Hospital
}
