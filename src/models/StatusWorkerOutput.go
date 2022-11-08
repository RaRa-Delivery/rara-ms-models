package models

import "github.com/RaRa-Delivery/rara-ms-models/src/models/order"

type StatusWorkerOutput struct {
	Status      string            `json:"status" bson:"status"`
	Message     string            `json:"message" bson:"message"`
	TrackingId  string            `json:"trackingId" bson:"trackingId"`
	OrderStatus order.OrderStatus `json:"orderStatus" bson:"orderStatus"`
}
