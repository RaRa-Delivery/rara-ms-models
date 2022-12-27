package models

type VehicleDocuments struct {
	VehicleRegistrationPhoto string `json:"vehicleRegistrationPhoto" bson:"vehicleRegistrationPhoto"`
	VehicleFrontPhoto        string `json:"vehicleFrontPhoto" bson:"vehicleFrontPhoto"`
	VehicleTaxPhoto          string `json:"vehicleTaxPhoto" bson:"vehicleTaxPhoto"`
	VehicleSidePhoto         string `json:"vehicleSidePhoto" bson:"vehicleSidePhoto"`
	VehicleLicensePlate      string `json:"vehicleLicensePlate" bson:"vehicleLicensePlate"`
	RegistrationExpiryDate   string `json:"registrationExpiryDate" bson:"registrationExpiryDate"`
	BrandId                  string `json:"brandId" bson:"brandId"`
	TypeId                   string `json:"typeId" bson:"typeId"`
	RabbitMobile             string `json:"rabbitMobile" bson:"rabbitMobile"`
	CreatedAt                int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt                int64  `json:"updatedAt" bson:"updatedAt"`
	Status                   string `json:"status" bson:"status"`
}
