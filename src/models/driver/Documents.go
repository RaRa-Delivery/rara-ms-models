package models

type Documents struct {
	Profile          Profile          `json:"profile" bson:"profile"`
	IdCard           IdCard           `json:"idCard" bson:"idCard"`
	License          License          `json:"license" bson:"license"`
	VehicleDocuments VehicleDocuments `json:"vehicleDocuments" bson:"vehicleDocuments"`
	BankDetails      BankDetails      `json:"bankDetails" bson:"bankDetails"`
}
