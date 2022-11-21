package models

type DriverRegistrationRequest struct {
	Firstname   string `json:"firstName" bson:"firstName"`
	Lastname    string `json:"lastName" bson:"lastName"`
	Emailid     string `json:"emailId" bson:"emailId"`
	Mobileno    string `json:"mobileNo" bson:"mobileNo"`
	Kecamatan   string `json:"kecamatan" bson:"kecamatan"`
	CityId      int    `json:"cityId" bson:"cityId"`
	VehicleType int    `json:"vehicleType" bson:"vehicleType"`
}
