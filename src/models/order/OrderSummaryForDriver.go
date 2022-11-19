package order

type OrderSummaryForDriver struct {
	OrderId struct {
		Id string `json:"id" bson:"id"`
	} `json:"orderId" bson:"orderId"`
	Geo     Geo `json:"geo" bson:"geo"`
	Details struct {
		OrderDetails struct {
			DeliveryDetails struct {
				Tag string `json:"specialHandling" bson:"specialHandling"`
			} `json:"orderDeliveryDetails" bson:"orderDeliveryDetails"`
		} `json:"orderDetails" bson:"orderDetails"`
		PickupDetails struct {
			Location struct {
				Name    string `json:"locationName" bson:"locationName"`
				Address string `json:"address" bson:"address"`
			} `json:"locationDetails" bson:"locationDetails"`
		} `json:"pickupDetails" bson:"pickupDetails"`
		HubDetails struct {
			Name    string `json:"name" bson:"name"`
			Address string `json:"address" bson:"address"`
		} `json:"hubDetails" bson:"hubDetails"`
	} `json:"newOrderObject" bson:"newOrderObject"`
}
