package models

type IdCard struct {
	IdCardPhoto           string `json:"idCardPhoto" bson:"idCardPhoto"`
	SelfieWithIdCardPhoto string `json:"selfieWithIdCardPhoto" bson:"selfieWithIdCardPhoto"`
	IdNumber              string `json:"idNumber" bson:"idNumber"`
	RabbitMobile          string `json:"rabbitMobile" bson:"rabbitMobile"`
	CreatedAt             int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt             int64  `json:"updatedAt" bson:"updatedAt"`
	Status                string `json:"status" bson:"status"`
}
