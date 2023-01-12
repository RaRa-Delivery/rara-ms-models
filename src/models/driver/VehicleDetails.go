package models

type VehicleDetails struct {
	Id      string      `json:"id" bson:"id"`
	Name    VehicleName `json:"name" bson:"name"`
	Status  string      `json:"status" bson:"status"`
	Vehicle Vehicle     `json:"vehicle" bson:"vehicle"`
	Brand   string      `json:"brand" bson:"brand"`
}
