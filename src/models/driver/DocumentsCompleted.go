package models

type DocumentsCompleted struct {
	Profile          bool `json:"profile" bson:"profile"`
	IdCard           bool `json:"idCard" bson:"idCard"`
	License          bool `json:"license" bson:"license"`
	VehicleDocuments bool `json:"vehicleDocuments" bson:"vehicleDocuments"`
	BankDetails      bool `json:"bankDetails" bson:"bankDetails"`
}
