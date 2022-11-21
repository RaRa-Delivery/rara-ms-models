package models

type EmergencyContact struct {
	Name         string `json:"name" bson:"name"`
	Relationship string `json:"relationship" bson:"relationship"`
	MobileNo     string `json:"mobileNo" bson:"mobileNo"`
}
