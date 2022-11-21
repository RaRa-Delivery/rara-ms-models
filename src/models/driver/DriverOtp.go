package models

type DriverOtp struct {
	Mobile string `json:"mobile" bson:"mobile"`
	Otp    string `json:"otp" bson:"otp"`
}
