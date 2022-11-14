package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Driver struct {
	Id                     primitive.ObjectID      `json:"_id" bson:"_id"`
	DriverDetails          DriverDetails           `json:"driverDetails" bson:"driverDetails"`
	LocationDetails        LocationDetails         `json:"locationDetails" bson:"locationDetails"`
	Status                 string                  `json:"status" bson:"status"`
	Otppassword            string                  `json:"otpPassword" bson:"otpPassword"`
	BusinessType           BusinessType            `json:"businessType" bson:"businessType"`
	VehicleDetails         VehicleDetails          `json:"vehicleDetails" bson:"vehicleDetails"`
	FirebaseAuthKey        string                  `json:"firebaseAuthKey" bson:"firebaseAuthKey"`
	DriverTrainingAttempts []DriverTrainingAttempt `json:"driverTrainingAttempts" bson:"driverTrainingAttempts"`
	RabbitType             string                  `json:"rabbitType" bson:"rabbitType"`
	LastBatch              string                  `json:"lastBatch" bson:"lastBatch"`
	ActiveBatch            string                  `json:"activeBatch" bson:"activeBatch"`
	LastEarnings           float64                 `json:"lastEarnings" bson:"lastEarnings"`
}

type DriverTrainingAttempt struct {
	TrainingResult      string `json:"training_result" bson:"training_result"`
	TrainingAttempt     int    `json:"training_attempt" bson:"training_attempt"`
	Status              string `json:"status" bson:"status"`
	CreatedAtUnix       int64  `json:"created_at_unix" bson:"unix_date_unix"`
	CreatedAt           string `json:"created_at" bson:"created_at"`
	SubmissionId        string `json:"submission_id" bson:"submission_id"`
	NextAttemptDate     string `json:"next_attempt_date" bson:"next_attempt_date"`
	NextAttemptDateUnix int64  `json:"next_attempt_date_unix" bson:"next_attempt_date_unix"`
	DriverExam          Exam   `json:"driver_exam" bson:"driver_exam"`
}

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
