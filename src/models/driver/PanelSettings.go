package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PanelSettings struct {
	Id     primitive.ObjectID `json:"_id" bson:"_id"`
	Group  string             `json:"group" bson:"group"`
	Title  string             `json:"title" bson:"title"`
	Key    string             `json:"key" bson:"key"`
	Value  interface{}        `json:"value" bson:"value"`
	BeKey  string             `json:"be_key,omitempty" bson:"be_key,omitempty"`
	CityId int64              `json:"cityId,omitempty" bson:"cityId,omitempty"`
}
