package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vehicle struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id"`
	Mileage         int                `json:"mileage" bson:"mileage"`
	Status          string             `json:"status" bson:"status"`
	IsShow          int                `json:"isShow" bson:"isShow"`
	Name            VehicleName        `json:"name" bson:"name"`
	EngineCapacity  int                `json:"engineCapacity" bson:"engineCapacity"`
	SeatingCapacity int                `json:"seatingCapacity" bson:"seatingCapacity"`
	LoadingCapacity int                `json:"loadingCapacity" bson:"loadingCapacity"`
	FuelType        string             `json:"fuelType" bson:"fuelType"`
	FuelCapacity    int                `json:"fuelCapacity" bson:"fuelCapacity"`
	CreatedAt       int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	City            []int              `json:"city" bson:"city"`
}

type AdminVehicle struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id"`
	Mileage         int                `json:"mileage" bson:"mileage"`
	Status          string             `json:"status" bson:"status"`
	IsShow          int                `json:"isShow" bson:"isShow"`
	Name            VehicleName        `json:"name" bson:"name"`
	EngineCapacity  int                `json:"engineCapacity" bson:"engineCapacity"`
	SeatingCapacity int                `json:"seatingCapacity" bson:"seatingCapacity"`
	LoadingCapacity int                `json:"loadingCapacity" bson:"loadingCapacity"`
	FuelType        string             `json:"fuelType" bson:"fuelType"`
	FuelCapacity    int                `json:"fuelCapacity" bson:"fuelCapacity"`
	City            []int              `json:"city" bson:"city"`
	VehicleType     struct {
		Id     primitive.ObjectID `json:"_id" bson:"_id"`
		Status string             `json:"status" bson:"status"`
		Name   VehicleName        `json:"name" bson:"name"`
	} `json:"vehicleType" bson:"vehicleType"`
}
