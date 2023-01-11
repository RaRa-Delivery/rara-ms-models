package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type VehicleBrand struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	CreatedAt int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
}
