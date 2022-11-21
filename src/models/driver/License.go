package models

type License struct {
	DrivingLicensePhoto  string `json:"drivingLicensePhoto" bson:"drivingLicensePhoto"`
	DrivingLicenseNumber string `json:"drivingLicenseNumber" bson:"drivingLicenseNumber"`
	ExpiryDate           string `json:"expiryDate" bson:"expiryDate"`
}
