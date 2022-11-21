package models

type Preference struct {
	Priority       int      `json:"priority" bson:"priority"`
	Status         string   `json:"status" bson:"status"`
	BshtTag        BSHT     `json:"bshtTag" bson:"bshtTag"`
	PickupLocation Location `json:"pickupLocation" bson:"pickupLocation"`
}
