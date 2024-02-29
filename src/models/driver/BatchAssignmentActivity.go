package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PnAssignment struct {
	Id                    primitive.ObjectID `json:"_id" bson:"_id"`
	DriverId              string             `json:"driverId" bson:"driverId"`
	BatchId               string             `json:"batchId" bson:"batchId"`
	AssignmentId          string             `json:"assignmentId" bson:"assignmentId"`
	PnSent                int64              `json:"pnSent" bson:"pnSent"`
	PnReceive             int64              `json:"pnReceive" bson:"pnReceive"`
	PnClick               int64              `json:"pnClick" bson:"pnClick"`
	PnReject              int64              `json:"pnReject" bson:"pnReject"`
	RejectReason          string             `json:"rejectReason" bson:"rejectReason"`
	AssignmentView        int64              `json:"assignmentView" bson:"assignmentView"`
	DriverStatus          string             `json:"driverStatus" bson:"driverStatus"`
	DriverStatusTime      string             `json:"driverStatusTime" bson:"driverStatusTime"`
	CreatedAt             int64              `json:"createdAt" bson:"createdAt"`
	BatchStartingTime     int64              `json:"batchStartingTime" bson:"batchStartingTime"`
	TotalDuration         int64              `json:"totalDuration" bson:"totalDuration"`
	DriverInitialDistance float64            `json:"driverInitialDistance" bson:"driverInitialDistance"`
	AcceptedAt            int64              `json:"acceptedAt" bson:"acceptedAt"`
	CompletedAt           int64              `json:"completedAt" bson:"completedAt"`
	StartPickupAt         int64              `json:"startPickupAt" bson:"startPickupAt"`
	OperationRegion       int64              `json:"operationRegion" bson:"operationRegion"`
	PickupZone            PnZone             `json:"pickupZone" bson:"pickupZone"`
	DropoffZone           PnZone             `json:"dropoffZone" bson:"dropoffZone"`
	BatchAddedAt          int64              `json:"batchAddedAt" bson:"batchAddedAt"`
	BussinessIds          []int64            `json:"bussinessIds" bson:"bussinessIds"`
	BussinessAccounts     []int64            `json:"bussinessAccounts" bson:"bussinessAccounts"`
	Bsht                  int64              `json:"bsht" bson:"bsht"`
	OrderCount            int64              `json:"orderCount" bson:"orderCount"`
	Active                bool               `json:"active" bson:"active" default:"true"`
	BatchSeen             int64              `json:"batchSeen" bson:"batchSeen"`
	PickLat               float64            `json:"pickLat" bson:"pickLat"`
	PickLng               float64            `json:"pickLng" bson:"pickLng"`
	Version               int                `json:"version" bson:"version"`
}

type BatchAssignmentActivity struct {
	PnAssignment
}

type PnZone struct {
	Id   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
