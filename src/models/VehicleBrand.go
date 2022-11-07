package models

type VehicleBrand struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
