package models

type Vehicles struct {
	Id      int         `json:"id" bson:"id"`
	Name    VehicleName `json:"name" bson:"name"`
	Status  string      `json:"status" bson:"status"`
	Vehicle []Vehicle   `json:"vehicles" bson:"vehicles"`
	City    []City      `json:"city" bson:"city"`
}
