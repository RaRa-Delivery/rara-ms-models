package models

import models "github.com/RaRa-Delivery/rara-ms-models/src/models/dos"

type DriverBatch struct {
	DriverMobile string         `json:"driverMobile" bson:"driverMobile"`
	BatchId      string         `json:"batchId" bson:"batchId"`
	Creator      models.Creator `json:"creator" bson:"creator"`
}
