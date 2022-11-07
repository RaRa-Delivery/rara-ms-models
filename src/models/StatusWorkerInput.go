package models

import "rara-ms-models/src/models/order"

type StatusWorkerInput struct {
	RequestedSt        DeliveryStatusDto    `json:"requestedSt" bson:"requestedSt"`
	TrackingId         string               `json:"trackingId" bson:"trackingId"`
	UpdateOrderRequest UpdateStatusRequest  `json:"updateOrderRequest" bson:"updateOrderRequest"`
	Description        string               `json:"description" bson:"description"`
	Driver             order.Driver         `json:"driver" bson:"driver"`
	OrderObj           order.NewOrderObject `json:"orderObj" bson:"orderObj"`
}

type DeliveryStatusDto struct {
	Index       float32   `json:"index" bson:"index"`
	Status      string    `json:"status" bson:"status"`
	Code        string    `json:"code" bson:"code"`
	Description string    `json:"description" bson:"description"`
	Attempt     int       `json:"attempt" bson:"attempt"`
	PrevIndex   []float32 `json:"prevIndex" bson:"prevIndex"`
	NextIndex   []float32 `json:"nextIndex" bson:"nextIndex"`
	BatchStatus string    `json:"batchStatus" bson:"batchStatus"`
}
