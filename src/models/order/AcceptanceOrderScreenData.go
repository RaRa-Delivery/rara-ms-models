package order

type AcceptanceOrderScreenData struct {
	DriverLoc          Geo     `json:"driverLoc" bson:"driverLoc"`
	FpuLoc             Geo     `json:"fpuLoc" bson:"fpuLoc"`
	FpuName            string  `json:"fpuName" bson:"fpuName"`
	FpuAddress         string  `json:"fpuAddress" bson:"fpuAddress"`
	FpuReachTime       float64 `json:"fpuReachTime" bson:"fpuReachTime"`
	FpuReachDistance   float64 `json:"fuReachDistance" bson:"fpuReachDistance"`
	TotalBatchOrders   int     `json:"totalBatchOrders" bson:"totalBatchOrders"`
	TotalBatchDistance float64 `json:"totalBatchDistance" bson:"totalBatchDistance"`
	TotalBatchFee      float64 `json:"totalBatchFee" bson:"totalBatchFee"`
	DriverLastEarning  float64 `json:"driverLastEarning" bson:"driverLastEarning"`
}
