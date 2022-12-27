package models

type BusinessType struct {
	DriverType string `json:"driverType" bson:"driverType"`
	BSHT       []BSHT `json:"bsht" bson:"bsht"`
}
