package models

type ReturnOrderRequest struct {
	To         string  `json:"to" bson:"to"`
	TrackingId string  `json:"trackingId" bson:"trackingId"`
	DriverId   string  `json:"driverMobileNo" bson:"driverMobileNo"`
	Lat        float64 `json:"lat" bson:"lat"`
	Lng        float64 `json:"lng" bson:"lng"`
	Ops        string  `json:"ops" bson:"ops"`
}
