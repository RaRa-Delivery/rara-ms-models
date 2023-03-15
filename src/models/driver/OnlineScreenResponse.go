package models

import "github.com/RaRa-Delivery/rara-ms-models/src/models/order"

type OnlineScreenResponse struct {
	TodaysEarning   string                    `json:"earningType" bson:"earningType"`
	Earning         Earning                   `json:"earning" bson:"earning"`
	Orders          []order.BatchForDriverApp `json:"nodes" bson:"nodes"`
	Kecamatan       []string                  `json:"kecamatan" bson:"kecamatan"`
	IsBatchAssigned bool                      `json:"isBatchAssigned" bson:"isBatchAssigned"`
	IsBatchAccepted bool                      `json:"isBatchAccepted" bson:"isBatchAccepted"`
	IsFirstStartPickup bool `json:"isFirstStartPickup" bson:"isFirstStartPickup"`
	StartTime       int64                     `json:"startTime" bson:"startTime"`
	Duration        int64                     `json:"duration" bson:"duration"`
	TimeRemaining   int64                     `json:"timeRemaining" bson:"timeRemaining"`
}
