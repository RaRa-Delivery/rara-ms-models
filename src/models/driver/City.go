package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type City struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	CityId     int                `json:"cityId,omitempty" bson:"cityId,omitempty"`
	CityName   string             `json:"cityName" bson:"cityName"`
	CitySlug   string             `json:"citySlug" bson:"citySlug"`
	CityStatus string             `json:"cityStatus" bson:"cityStatus"`
	Display    string             `json:"display" bson:"display"`
	Version    int                `json:"version" bson:"version"`
}
