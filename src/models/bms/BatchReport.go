package bms

import "go.mongodb.org/mongo-driver/bson/primitive"

type BatchDriver struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	MobileNo string `json:"mobileNo" bson:"mobileNo"`
}

type AssignmentEvents struct {
	Driver    BatchDriver `json:"driver" bson:"driver"`
	CreatedAt int64       `json:"createdAt" bson:"createdAt"`
}

type BatchOperationRegion struct {
	Id   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type BatchReports struct {
	Id primitive.ObjectID `json:"_id" bson:"_id"`
	//During batch creation (BMS service only)
	BatchId            string               `json:"batchId" bson:"batchId"`
	PickupLocation     string               `json:"pickupLocation" bson:"pickupLocation"`
	OperationRegion    BatchOperationRegion `json:"operationRegion" bson:"operationRegion"`
	PickupZone         string               `json:"pickupZone" bson:"pickupZone"`
	DropoffZone        string               `json:"dropoffZone" bson:"dropoffZone"`
	BshtTag            string               `json:"bshtTag" bson:"bshtTag"`
	CreatedAt          int64                `json:"createdAt" bson:"createdAt"`
	ExpectedDeliveryAt int64                `json:"expectedDeliveryAt" bson:"expectedDeliveryAt"`

	//Druring batch assignment (DOAS service only)
	AssignmentAttempts   int                `json:"assignmentAttempts" bson:"assignmentAttempts"`
	AssignmentAttemptsBy []AssignmentEvents `json:"assignmentAttemptsBy" bson:"assignmentAttemptsBy"`
	AssignmentSeen       int                `json:"assignmentSeen" bson:"assignmentSeen"`
	AssignmentSeenBy     []AssignmentEvents `json:"assignmentSeenBy" bson:"assignmentSeenBy"`
	AssignmentMissed     int                `json:"assignmentMissed" bson:"assignmentMissed"`
	AssignmentMissedBy   []AssignmentEvents `json:"assignmentMissedBy" bson:"assignmentMissedBy"`
	AssignmentAccepted   int                `json:"assignmentAccepted" bson:"assignmentAccepted"`
	AssignmentAcceptedBy []AssignmentEvents `json:"assignmentAcceptedBy" bson:"assignmentAcceptedBy"`
	AssignmentRejected   int                `json:"assignmentRejected" bson:"assignmentRejected"`
	AssignmentRejectedBy []AssignmentEvents `json:"assignmentRejectedBy" bson:"assignmentRejectedBy"`
	AssignmentIgnored    int                `json:"assignmentIgnored" bson:"assignmentIgnored"`
	AssignmentIgnoredBy  []AssignmentEvents `json:"assignmentIgnoredBy" bson:"assignmentIgnoredBy"`

	//During status update (DOS service only)
	Status      string      `json:"status" bson:"status"`           //It can be changed on every delivere status update
	Driver      BatchDriver `json:"driver" bson:"driver"`           //It can be changed on batch assigned/accepted/rejected
	AcceptedAt  int64       `json:"acceptedAt" bson:"acceptedAt"`   //It can be changed on batch accepted
	CompletedAt int64       `json:"completedAt" bson:"completedAt"` //It can be changed on batch completion

	OrderCount          int `json:"orderCount" bson:"orderCount"`                   //It can be changed on batch creation/add order/cancel order/pickup failed
	DeliveredOrderCount int `json:"deliveredOrderCount" bson:"deliveredOrderCount"` //It can be changed on order delivery of the batch

	UpdatedAt int64 `json:"updatedAt" bson:"updatedAt"` //With every events, it will be updated

}
