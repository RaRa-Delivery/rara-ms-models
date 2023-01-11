package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vehicles struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	Name    VehicleName        `json:"name" bson:"name"`
	Status  string             `json:"status" bson:"status"`
	City    []City             `json:"city" bson:"city"`
	Vehicle []Vehicle          `json:"vehicles" bson:"vehicles"`
}
