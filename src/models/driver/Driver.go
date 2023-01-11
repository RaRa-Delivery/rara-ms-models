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
	DriverTrainings        []Exam                  `json:"driverTrainings" bson:"driverTrainings"`
	RabbitType             string                  `json:"rabbitType" bson:"rabbitType"`
	Documents              Documents               `json:"documents" bson:"documents"`
	DocumentsCompleted     DocumentsCompleted      `json:"documentsCompleted" bson:"documentsCompleted"`
	ActiveBatchId          string                  `json:"activeBatch" bson:"activeBatch"`
	DriverEarnings         DriverEarnings          `json:"driverEarnings" bson:"driverEarnings"`
	CreatedAt              int64                   `json:"createdAt" bson:"createdAt"`
	Miscellaneous          Miscellaneous           `json:"miscellaneous" bson:"miscellaneous"`
	StatusDetails          interface{}             `json:"statusDetails" bson:"statusDetails"`
	LastBatch              string                  `json:"lastBatch" bson:"lastBatch"`
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
