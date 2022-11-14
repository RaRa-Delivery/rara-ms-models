package models

type VehicleDetails struct {
	Id      int         `json:"id" bson:"id"`
	Name    VehicleName `json:"name" bson:"name"`
	Status  string      `json:"status" bson:"status"`
	Vehicle Vehicle     `json:"vehicles" bson:"vehicles"`
}
