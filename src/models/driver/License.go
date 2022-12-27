package models

type License struct {
	DrivingLicensePhoto  string `json:"drivingLicensePhoto" bson:"drivingLicensePhoto"`
	DrivingLicenseNumber string `json:"drivingLicenseNumber" bson:"drivingLicenseNumber"`
	ExpiryDate           string `json:"expiryDate" bson:"expiryDate"`
	RabbitMobile         string `json:"rabbitMobile" bson:"rabbitMobile"`
	CreatedAt            int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt            int64  `json:"updatedAt" bson:"updatedAt"`
	Status               string `json:"status" bson:"status"`
}
