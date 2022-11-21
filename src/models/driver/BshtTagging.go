package models

type BshtTagging struct {
	DriverId string `json:"driverId" bson:"driverId"`
	TagId    int    `json:"tagId" bson:"tagId"`
}
