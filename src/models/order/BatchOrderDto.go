package order

type BatchOrderDto struct {
	BatchDetails        BatchDetails        `json:"batchDetails" bson:"batchDetails"`
	StatusDetails       StatusDetails       `json:"statusDetails" bson:"statusDetails"`
	OrderDistanceMatrix OrderDistanceMatrix `json:"orderDistanceMatrix" bson:"orderDistanceMatrix"`
	Driver              Driver              `json:"driver" bson:"driver"`
}

type WebhookOrderDto struct {
	TrackingId string   `json:"trackingId" bson:"trackingId"`
	Codes      []string `json:"codes" bson:"codes"`
}
