package order

import "go.mongodb.org/mongo-driver/bson/primitive"

type AcceptanceOrderScreenData struct {
	DriverLoc          Geo                 `json:"driverLoc" bson:"driverLoc"`
	BatchId            string              `json:"batchId" bson:"batchId"`
	BshtTag            string              `json:"bshtTag" bson:"bshtTag"`
	Orders             []BatchForDriverApp `json:"nodes" bson:"nodes"`
	FpuName            string              `json:"fpuName" bson:"fpuName"`
	FpuAddress         string              `json:"fpuAddress" bson:"fpuAddress"`
	FpuReachTime       float64             `json:"fpuReachTime" bson:"fpuReachTime"`
	FpuReachDistance   float64             `json:"fuReachDistance" bson:"fpuReachDistance"`
	TotalBatchOrders   int                 `json:"totalBatchOrders" bson:"totalBatchOrders"`
	TotalBatchDistance float64             `json:"totalBatchDistance" bson:"totalBatchDistance"`
	TotalBatchFee      float64             `json:"totalBatchFee" bson:"totalBatchFee"`
	EarningType        string              `json:"earningType" bson:"earningType"`
	Earning            Earning             `json:"earning" bson:"earning"`
	StartTime          int64               `json:"startTime" bson:"startTime"`
	Duration           int64               `json:"duration" bson:"duration"`
	TimeRemaining      int64               `json:"timeRemaining" bson:"timeRemaining"`
}

type BatchForDriverApp struct {
	Id           string             `json:"id"`
	Type         string             `json:"type"`
	Address      LocationDetails    `json:"address"`
	Sla          float64            `json:"nodeSla" bson:"nodeSla"`
	Prioritize   bool               `json:"prioritize" bson:"prioritize"`
	BatchId      primitive.ObjectID `json:"batchId" bson:"batchId"`
	NodeIncharge struct {
		Name  string `json:"name" bson:"name"`
		Phone string `json:"phone" bson:"phone"`
		Email string `json:"email" bson:"email"`
	} `json:"nodeIncharge" bson:"nodeIncharge"`
	Orders []OrdersForDriverApp `json:"orders"`
}

type OrdersForDriverApp struct {
	TrackingId    string                   `json:"trackingId"`
	Sla           float64                  `json:"sla" bson:"sla"`
	Address       DriverAppLocationDetails `json:"address"`
	Status        string                   `json:"status"`
	PuNote        string                   `json:"puNote" bson:"puNote"`
	DoNote        string                   `json:"doNote" bson:"doNote"`
	BshtTag       string                   `json:"bshtTag"`
	Otp           string                   `json:"otp"`
	Attempt       int                      `json:"attempt" bson:"attempt"`
	OrderIncharge struct {
		Name  string `json:"name" bson:"name"`
		Phone string `json:"phone" bson:"phone"`
		Email string `json:"email" bson:"email"`
	} `json:"orderIncharge" bson:"orderIncharge"`
	PaymentDetails struct {
		PaymentMethod string  `json:"paymentMethod" bson:"paymentMethod"`
		Price         float64 `json:"price" bson:"price"`
	} `json:"paymentDetails" bson:"paymentDetails"`
}

type DriverAppLocationDetails struct {
	LocationName string            `json:"locationName" bson:"locationName"`
	Address      string            `json:"address" bson:"address"`
	SubDistrict  string            `json:"subDistrict" bson:"subDistrict"`
	District     string            `json:"district" bson:"district"`
	City         string            `json:"city" bson:"city"`
	PostalCode   string            `json:"postalCode" bson:"postalCode"`
	GeoPoint     DriverAppGeoPoint `json:"geoPoint" bson:"geoPoint"`
	Type         string            `json:"type" bson:"type"`
	Province     string            `json:"province" bson:"province"`
}

type DriverAppGeoPoint struct {
	Lat     float64 `json:"lat" bson:"lat"`
	Lng     float64 `json:"lng" bson:"lng"`
	PinCode string  `json:"pinCode" bson:"pinCode"`
	GeoHash string  `json:"geoHash" bson:"geoHash"`
}

type DriverEarnings struct {
	LastEarning    Earning `json:"lastEarnings" bson:"lastEarnings"`
	TodayEarning   Earning `json:"todayEarnings" bson:"todayEarnings"`
	HighestEarning Earning `json:"highestEarnings" bson:"highestEarnings"`
}

type Earning struct {
	Amount float64 `json:"amount" bson:"amount"`
	Date   string  `json:"date" bson:"date"`
}
