package bms

type BatchMilestoneDto struct {
	Milestone      string       `json:"milestone"`
	TrackingID     string       `json:"trackingId"`
	CreatedAt      int64        `json:"createdAt"`
	Actor          BmsActor     `json:"Actor"`
	Pickup         BmsPickup    `json:"pickup"`
	Dropoff        BmsDropoff   `json:"dropoff"`
	DriverLoc      BmsDriverLoc `json:"driverLoc"`
	DriverDistance string       `json:"driverDistance"`
	Notes          string       `json:"notes"`
	Logs           string       `json:"logs"`
}
type BmsActor struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	Role   string `json:"role"`
}
type BmsPickup struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type BmsDropoff struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type BmsDriverLoc struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
