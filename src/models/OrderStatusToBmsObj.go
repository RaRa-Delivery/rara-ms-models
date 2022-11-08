package models

import "github.com/RaRa-Delivery/rara-ms-models/src/models/order"

type OrderStatusToBmsObj struct {
	TrackingId  string            `json:"trackingId" bson:"trackingId"`
	Status      string            `json:"status" bson:"status"`
	Driver      Driver            `json:"driver" bson:"driver"`
	OrderStatus order.OrderStatus `json:"orderStatus" bson:"orderStatus"`
	Creator     Creator           `json:"creator" bson:"creator"`
	Attempt     int               `json:"attempt" bson:"attempt"`
	BatchId     string            `json:"batchId" bson:"batchId"`
}

type Creator struct {
	Id           string `json:"id" bson:"id"`
	Name         string `json:"name" bson:"name"`
	MobileNumber string `json:"mobileNumber" bson:"mobileNumber"`
	Type         string `json:"type" bson:"type"`
}
