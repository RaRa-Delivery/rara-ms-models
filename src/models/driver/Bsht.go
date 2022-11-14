package models

type BSHT struct {
	TagId  int    `json:"tagId" bson:"tagId"`
	Label  string `json:"label" bson:"label"`
	Code   string `json:"code" bson:"code"`
	Status string `json:"status" bson:"status"`
}
