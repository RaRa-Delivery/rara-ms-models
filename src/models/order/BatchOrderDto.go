package order

type BatchOrderDto struct {
	BatchDetails        BatchDetails        `json:"batchDetails" bson:"batchDetails"`
	StatusDetails       StatusDetails       `json:"statusDetails" bson:"statusDetails"`
	OrderDistanceMatrix OrderDistanceMatrix `json:"orderDistanceMatrix" bson:"orderDistanceMatrix"`
}
