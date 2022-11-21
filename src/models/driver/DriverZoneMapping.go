package models

type DzmRequestBody struct {
	BshtTag         BSHT       `json:"bshtTag" bson:"bshtTag"`
	PickupLocations []Location `json:"pickupLocations" bson:"pickupLocations"`
}
