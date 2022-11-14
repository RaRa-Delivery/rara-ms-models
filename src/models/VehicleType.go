package models

type VehicleType struct {
	Id             int         `json:"id" bson:"id"`
	Name           VehicleName `json:"name" bson:"name"`
	Status         string      `json:"status" bson:"status"`
	City           []int       `json:"city" bson:"city"`
	Vehicle        []int       `json:"vehicles" bson:"vehicles"`
	CreatedAt      int64       `json:"created_at,omitempty" bson:"created_at,omitempty"`
	LastModifiedAt int64       `json:"last_modified_at,omitempty" bson:"last_modified_at,omitempty"`
}
