package order

type Geo struct {
	Lat     float64 `json:"lat" bson:"lat"`
	Lng     float64 `json:"lng" bson:"lng"`
	PinCode string  `json:"pinCode" bson:"pinCode"`
	GeoHash string  `json:"geoHash" bson:"geoHash"`
}
