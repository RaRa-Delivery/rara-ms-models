package order

type OrderBmsProduceModel struct {
	BatchId  string `json:"batchId" bson:"batchId"`
	DriverId string `json:"driverId" bson:"driverId"`
	OrderId  string `json:"orderId" bson:"orderId"`
}
