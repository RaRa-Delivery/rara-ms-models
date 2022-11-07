package models

type DriverBatch struct {
	DriverMobile string  `json:"driverMobile" bson:"driverMobile"`
	BatchId      string  `json:"batchId" bson:"batchId"`
	Creator      Creator `json:"creator" bson:"creator"`
}
