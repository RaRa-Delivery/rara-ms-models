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
	CreatedBy             string             `json:"createdBy" bson:"createdBy"`
	UpdatedBy             string             `json:"updatedBy" bson:"updatedBy"`

	CreatedAt int64 `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64 `json:"updatedAt" bson:"updatedAt"`
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

type DriverExam struct {
	ID                     string      `json:"id"`
	NumberOfOrders         int         `json:"numberOfOrders"`
	ExamName               string      `json:"examName"`
	JotFormLink            string      `json:"jotFormLink"`
	MinScore               int         `json:"minScore"`
	MaxAttempts            int         `json:"maxAttempts"`
	ExamDuration           int         `json:"examDuration"`
	AllowTrainingAfterDays int         `json:"allowTrainingAfterDays"`
	Status                 string      `json:"status"`
	Driver                 interface{} `json:"driver"`
	Allowed                interface{} `json:"allowed"`
}
type DriverData struct {
	Id                  string      `json:"id"`
	MobileNumber        string      `json:"mobileNumber"`
	TrainingResult      string      `json:"trainingResult"`
	TrainingAttempt     int         `json:"trainingAttempt"`
	Status              string      `json:"status"`
	UnixDate            int         `json:"unixDate"`
	CreatedAt           int         `json:"createdAt"`
	SubmitssionID       string      `json:"submitssionId"`
	NextAttemptDate     interface{} `json:"nextAttemptDate"`
	NextAttemptDateUnix interface{} `json:"nextAttemptDateUnix"`
	DriverExam          DriverExam  `json:"driverExam"`
}
type DriverExamResponse struct {
	ExamId                 string     `json:"examId"`
	NumberOfOrders         int        `json:"numberOfOrders"`
	ExamName               string     `json:"examName"`
	JotFormLink            string     `json:"jotFormLink"`
	MinScore               int        `json:"minScore"`
	MaxAttempts            int        `json:"maxAttempts"`
	ExamDuration           int        `json:"examDuration"`
	AllowTrainingAfterDays int        `json:"allowTrainingAfterDays"`
	Status                 string     `json:"status"`
	Driver                 DriverData `json:"driver"`
	Allowed                bool       `json:"allowed"`
}
