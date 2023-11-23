package models

type DriverDetails struct {
	Firstname   string    `json:"firstName" bson:"firstName"`
	Lastname    string    `json:"lastName" bson:"lastName"`
	Emailid     string    `json:"emailId" bson:"emailId"`
	Mobileno    string    `json:"mobileNo" bson:"mobileNo"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
