package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BatchAssignmentActivity struct {
	Id               primitive.ObjectID `json:"_id" bson:"_id"`
	DriverId         string             `json:"driverId" bson:"driverId"`
	BatchId          string             `json:"batchId" bson:"batchId"`
	AssignmentId     string             `json:"assignmentId" bson:"assignmentId"`
	PnSent           int64              `json:"pnSent" bson:"pnSent"`
	PnReceive        int64              `json:"pnReceive" bson:"pnReceive"`
	PnClick          int64              `json:"pnClick" bson:"pnClick"`
	PnReject         int64              `json:"pnReject" bson:"pnReject"`
	RejectReason     string             `json:"rejectReason" bson:"rejectReason"`
	AssignmentView   int64              `json:"assignmentView" bson:"assignmentView"`
	DriverStatus     string             `json:"driverStatus" bson:"driverStatus"`
	DriverStatusTime string             `json:"driverStatusTime" bson:"driverStatusTime"`
	CreatedAt        int64              `json:"createdAt" bson:"createdAt"`
}
