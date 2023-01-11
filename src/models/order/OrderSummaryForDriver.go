package order

type OrderSummaryForDriver struct {
	OrderId struct {
		Id string `json:"id" bson:"id"`
	} `json:"orderId" bson:"orderId"`
	Geo    Geo `json:"geo" bson:"geo"`
	Origin struct {
		DeliveryType string `json:"deliveryType" bson:"deliveryType"`
	} `json:"origin" bson:"origin"`
	Sla       int `json:"sla" bson:"sla"`
	PickupSla int `json:"pickupSla" bson:"pickupSla"`
	Details   struct {
		OrderDetails struct {
			DeliveryDetails struct {
				Tag string `json:"specialHandling" bson:"specialHandling"`
			} `json:"orderDeliveryDetails" bson:"orderDeliveryDetails"`
		} `json:"orderDetails" bson:"orderDetails"`
		PickupDetails struct {
			PickupIncharge struct {
				Name  string `json:"name" bson:"name"`
				Phone string `json:"phone" bson:"phone"`
				Email string `json:"email" bson:"email"`
			} `json:"pickupIncharge" bson:"pickupIncharge"`
			Location struct {
				Name        string `json:"locationName" bson:"locationName"`
				Address     string `json:"address" bson:"address"`
				SubDistrict string `json:"subDistrict" bson:"subDistrict"`
				District    string `json:"district" bson:"district"`
				City        string `json:"city" bson:"city"`
				PostalCode  string `json:"postalCode" bson:"postalCode"`
				Type        string `json:"type" bson:"type"`
				Province    string `json:"province" bson:"province"`
				Geo         Geo    `json:"geoPoint" bson:"geoPoint"`
			} `json:"locationDetails" bson:"locationDetails"`
			ExpectedPuDateAndTime string `json:"expectedPuDateTime" bson:"expectedPuDateTime"`
			Slot                  string `json:"slot" bson:"slot"`
			Note                  string `json:"note" bson:"note"`
		} `json:"pickupDetails" bson:"pickupDetails"`
		DropoffDetails struct {
			Notes string `json:"dlNote" bson:"dlNote"`
		} `json:"dropoffDetails" bson:"dropoffDetails"`
		HubDetails struct {
			Name    string `json:"name" bson:"name"`
			Address string `json:"address" bson:"address"`
		} `json:"hubDetails" bson:"hubDetails"`
		PaymentDetails struct {
			PaymentMethod string  `json:"paymentMethod" bson:"paymentMethod"`
			Price         float64 `json:"price" bson:"price"`
		} `json:"paymentDetails" bson:"paymentDetails"`
	} `json:"newOrderObject" bson:"newOrderObject"`
}
