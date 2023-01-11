package models

type BshtTagging struct {
	DriverId string `json:"driverId" bson:"driverId"`
	TagId    string `json:"tagId" bson:"tagId"`
}
