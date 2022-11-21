package models

type VehicleDocuments struct {
	VehicleRegistrationPhoto string `json:"vehicleRegistrationPhoto" bson:"vehicleRegistrationPhoto"`
	VehicleFrontPhoto        string `json:"vehicleFrontPhoto" bson:"vehicleFrontPhoto"`
	VehicleTaxPhoto          string `json:"vehicleTaxPhoto" bson:"vehicleTaxPhoto"`
	VehicleSidePhoto         string `json:"vehicleSidePhoto" bson:"vehicleSidePhoto"`
	VehicleLicensePlate      string `json:"vehicleLicensePlate" bson:"vehicleLicensePlate"`
	RegistrationExpiryDate   string `json:"registrationExpiryDate" bson:"registrationExpiryDate"`
	BrandId                  int    `json:"brandId" bson:"brandId"`
	TypeId                   int    `json:"typeId" bson:"typeId"`
}
