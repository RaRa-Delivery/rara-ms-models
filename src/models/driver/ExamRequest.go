package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exam struct {
	Id                    primitive.ObjectID `json:"_id" bson:"_id"`
	Position              int                `json:"position" bson:"position"`
	ExamName              string             `json:"examName" bson:"examName"`
	Link                  string             `json:"link" bson:"link"`
	Score                 int                `json:"score" bson:"score"`
	MaxAttempts           int                `json:"maxAttempts" bson:"maxAttempts"`
	ExamDuration          int                `json:"examDuration" bson:"examDuration"`
	AllowTraningAfterDays int                `json:"allowTraningAfterDays" bson:"allowTraningAfterDays"`
	Status                string             `json:"status" bson:"status"`
	BusinessType          string             `json:"businessType" bson:"businessType"`
	CreatedBy             int                `json:"createdBy" bson:"createdBy"`
	UpdatedBy             int                `json:"updatedBy" bson:"updatedBy"`
}

type ExamRequest struct {
	Position              int    `json:"position" bson:"position"`
	ExamName              string `json:"examName" bson:"examName"`
	Link                  string `json:"link" bson:"link"`
	Score                 int    `json:"score" bson:"score"`
	MaxAttempts           int    `json:"maxAttempts" bson:"maxAttempts"`
	ExamDuration          int    `json:"examDuration" bson:"examDuration"`
	AllowTraningAfterDays int    `json:"allowTraningAfterDays" bson:"allowTraningAfterDays"`
	Status                string `json:"status" bson:"status"`
	BusinessType          string `json:"businessType" bson:"businessType"`
}
