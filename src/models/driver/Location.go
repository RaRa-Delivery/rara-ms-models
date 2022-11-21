package models

type Location struct {
	PostalCode        string `json:"postalCode" bson:"postalCode"`
	SubDistrict       string `json:"subDistrict" bson:"subDistrict"`
	Districts         string `json:"districts" bson:"districts"`
	KotaBogor         string `json:"kota" bson:"kota"`
	Province          string `json:"province" bson:"province"`
	CityId            string `json:"cityId" bson:"cityId"`
	SearchDistrict    string `json:"searchDistrict" bson:"searchDistrict"`
	SearchSubDistrict string `json:"searchSubDistrict" bson:"searchSubDistrict"`
}
