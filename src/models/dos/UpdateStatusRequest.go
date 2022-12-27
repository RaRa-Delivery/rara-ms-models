package models

type UpdateStatusRequest struct {
	To                string   `json:"to" bson:"to"`
	BatchId           string   `json:"batchId" bson:"batchId"`
	FirstStartPickup  bool     `json:"firstStartPickup" bson:"firstStartPickup"`
	TrackingIds       []string `json:"trackingIds" bson:"trackingIds"`
	DriverMobile      string   `json:"driverMobileNo" bson:"driverMobileNo"`
	Lat               float64  `json:"lat" bson:"lat"`
	Lng               float64  `json:"lng" bson:"lng"`
	Ops               string   `json:"ops" bson:"ops"`
	OnlineStatus      string   `json:"onlineStatus" bson:"onlineStatus"`
	Comment           string   `json:"comment" bson:"comment"`
	RecipientName     string   `json:"recipientName" bson:"recipientName"`
	RecipientRelation string   `json:"recipientRelation" bson:"recipientRelation"`
	Images            []string `json:"images" bson:"images"`
}
