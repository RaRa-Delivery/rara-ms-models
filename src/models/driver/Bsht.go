package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BSHT struct {
	TagId  primitive.ObjectID `json:"_id" bson:"_id"`
	Label  string             `json:"label" bson:"label"`
	Code   string             `json:"code" bson:"code"`
	Status string             `json:"status" bson:"status"`
}
