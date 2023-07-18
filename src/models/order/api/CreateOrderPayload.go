package api

import (
	"encoding/json"
	"time"
)

type CreateOrderPayload struct {
	TrackingID     string         `json:"trackingId"`
	Schedule       Schedule       `json:"schedule"`
	Sender         Sender         `json:"sender"`
	Recipient      Recipient      `json:"recipient"`
	Origin         Origin         `json:"origin"`
	Destination    Destination    `json:"destination"`
	PaymentDetails PaymentDetails `json:"paymentDetails"`
	Packages       []Packages     `json:"packages"`
}
type Schedule struct {
	PickupDateTime time.Time `json:"pickupDateTime"`
}
type Sender struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Instruction string `json:"instruction"`
}
type Recipient struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Instruction string `json:"instruction"`
}
type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Origin struct {
	Id          int64    `json:"id"`
	Address     string   `json:"address"`
	GeoPoint    GeoPoint `json:"geoPoint"`
	SubDistrict string   `json:"subDistrict"`
	District    string   `json:"district"`
	City        string   `json:"city"`
	PostalCode  string   `json:"postalCode"`
	Province    string   `json:"province"`
}
type Destination struct {
	Address     string   `json:"address"`
	GeoPoint    GeoPoint `json:"geoPoint"`
	SubDistrict string   `json:"subDistrict"`
	District    string   `json:"district"`
	City        string   `json:"city"`
	PostalCode  string   `json:"postalCode"`
	Province    string   `json:"province"`
	DropOffNote string   `json:"dropOffNote"`
}
type PaymentDetails struct {
	Method string  `json:"method"`
	Price  float64 `json:"price"`
}
type Dimensions struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}
type Packages struct {
	Name       string     `json:"name"`
	Weight     float64    `json:"weight"`
	Dimensions Dimensions `json:"dimensions"`
}

func (o *CreateOrderPayload) NewCreateOrderPayload(payload string) error {
	return json.Unmarshal([]byte(payload), o)
}
