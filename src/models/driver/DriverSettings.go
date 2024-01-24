package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DriverSetting struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Group     string             `json:"group" bson:"group"`
	MetaLabel string             `json:"metaLabel" bson:"metaLabel"`
	MetaKey   string             `json:"metaKey" bson:"metaKey"`
	MetaValue string             `json:"metaValue" bson:"metaValue"`
	Version   int                `json:"version" bson:"version"`
}
