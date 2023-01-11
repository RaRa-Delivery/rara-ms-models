package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type VehicleType struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	Name           VehicleName        `json:"name" bson:"name"`
	Status         string             `json:"status" bson:"status"`
	City           []int              `json:"city" bson:"city"`
	Vehicle        []string           `json:"vehicles" bson:"vehicles"`
	CreatedAt      int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	LastModifiedAt int64              `json:"last_modified_at,omitempty" bson:"last_modified_at,omitempty"`
}
