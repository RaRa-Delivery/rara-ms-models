package models

type IdCard struct {
	IdCardPhoto           string `json:"idCardPhoto" bson:"idCardPhoto"`
	SelfieWithIdCardPhoto string `json:"selfieWithIdCardPhoto" bson:"selfieWithIdCardPhoto"`
	IdNumber              string `json:"idNumber" bson:"idNumber"`
}
