package models

type Profile struct {
	ImageUrl         string           `json:"imageUrl" bson:"imageUrl"`
	Kecamatan        string           `json:"kecamatan" bson:"kecamatan"`
	EmergencyContact EmergencyContact `json:"emergencyContact" bson:"emergencyContact"`
}
