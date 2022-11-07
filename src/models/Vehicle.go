package models

type Vehicle struct {
	Id              int         `json:"id" bson:"id"`
	City            []int       `json:"city" bson:"city"`
	Mileage         int         `json:"mileage" bson:"mileage"`
	Status          string      `json:"status" bson:"status"`
	IsShow          int         `json:"isShow" bson:"isShow"`
	Name            VehicleName `json:"name" bson:"name"`
	EngineCapacity  int         `json:"engineCapacity" bson:"engineCapacity"`
	SeatingCapacity int         `json:"seatingCapacity" bson:"seatingCapacity"`
	LoadingCapacity int         `json:"loadingCapacity" bson:"loadingCapacity"`
	FuelType        string      `json:"fuelType" bson:"fuelType"`
	FuelCapacity    int         `json:"fuelCapacity" bson:"fuelCapacity"`
	VehicleBrand    int         `json:"brand" bson:"brand"`
}
