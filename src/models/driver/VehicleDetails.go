package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type VehicleDetails struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	Name    VehicleName        `json:"name" bson:"name"`
	Status  string             `json:"status" bson:"status"`
	Vehicle Vehicle            `json:"vehicle" bson:"vehicle"`
	Brand   string             `json:"brand" bson:"brand"`
}
