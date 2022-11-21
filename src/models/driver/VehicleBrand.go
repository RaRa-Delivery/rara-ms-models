package models

type VehicleBrand struct {
	Id        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty" bson:"created_by,omitempty"`
}
