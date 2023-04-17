package models

type DriverLastLocation struct {
	DriverID     string   `json:"driverId"`
	Mobile       string   `json:"mobile"`
	Location     Location `json:"location"`
	RecordedTime string   `json:"recordedTime"`
	Availability bool     `json:"availability"`
	RabbitName   string   `json:"rabbitName"`
	Model        string   `json:"model"`
	City         City     `json:"city"`
}
type DriverLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type DriverCity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
