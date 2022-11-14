package models

type LocationDetails struct {
	Country     string  `json:"country" bson:"country"`
	CountryCode float64 `json:"countryCode" bson:"countryCode"`
	Kecamatan   string  `json:"kecamatan" bson:"kecamatan"`
	City        City    `json:"city" bson:"city"`
}
