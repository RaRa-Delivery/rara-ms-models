package models

import (
	"github.com/RaRa-Delivery/rara-ms-models/src/models/order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BatchSummaryForDriver struct {
	Id            primitive.ObjectID            `json:"_id" bson:"_id"`
	TotalOrders   int                           `json:"orderCount" bson:"orderCount"`
	TotalDistance float64                       `json:"batchDistance" bson:"batchDistance"`
	TotalFee      float64                       `json:"finalEarning" bson:"finalEarning"`
	Orders        []order.OrderSummaryForDriver `json:"batchedOrders" bson:"batchedOrders"`
	Version       int                           `json:"version" bson:"version"`
}
