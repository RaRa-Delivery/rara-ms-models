package models

type DriverEarnings struct {
	LastEarning    Earning `json:"lastEarnings" bson:"lastEarnings"`
	TodayEarning   Earning `json:"todayEarnings" bson:"todayEarnings"`
	HighestEarning Earning `json:"highestEarnings" bson:"highestEarnings"`
}

type Earning struct {
	Amount float64 `json:"amount" bson:"amount"`
	Date   string  `json:"date" bson:"date"`
}
