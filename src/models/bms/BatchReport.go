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
	Id                   primitive.ObjectID   `json:"_id" bson:"_id"`
	BatchId              string               `json:"batchId" bson:"batchId"`
	PickupLocation       string               `json:"pickupLocation" bson:"pickupLocation"`
	Status               string               `json:"status" bson:"status"`
	Driver               BatchDriver          `json:"driver" bson:"driver"`
	OperationRegion      BatchOperationRegion `json:"operationRegion" bson:"operationRegion"`
	PickupZone           string               `json:"pickupZone" bson:"pickupZone"`
	DropoffZone          string               `json:"dropoffZone" bson:"dropoffZone"`
	AssignmentAttempts   int                  `json:"assignmentAttempts" bson:"assignmentAttempts"`
	AssignmentAttemptsBy []AssignmentEvents   `json:"assignmentAttemptsBy" bson:"assignmentAttemptsBy"`
	AssignmentSeen       int                  `json:"assignmentSeen" bson:"assignmentSeen"`
	AssignmentSeenBy     []AssignmentEvents   `json:"assignmentSeenBy" bson:"assignmentSeenBy"`
	AssignmentMissed     int                  `json:"assignmentMissed" bson:"assignmentMissed"`
	AssignmentMissedBy   []AssignmentEvents   `json:"assignmentMissedBy" bson:"assignmentMissedBy"`
	AssignmentAccepted   int                  `json:"assignmentAccepted" bson:"assignmentAccepted"`
	AssignmentAcceptedBy []AssignmentEvents   `json:"assignmentAcceptedBy" bson:"assignmentAcceptedBy"`
	AssignmentRejected   int                  `json:"assignmentRejected" bson:"assignmentRejected"`
	AssignmentRejectedBy []AssignmentEvents   `json:"assignmentRejectedBy" bson:"assignmentRejectedBy"`
	AssignmentIgnored    int                  `json:"assignmentIgnored" bson:"assignmentIgnored"`
	AssignmentIgnoredBy  []AssignmentEvents   `json:"assignmentIgnoredBy" bson:"assignmentIgnoredBy"`
	BshtTag              string               `json:"bshtTag" bson:"bshtTag"`

	CreatedAt          int64 `json:"createdAt" bson:"createdAt"`
	AcceptedAt         int64 `json:"acceptedAt" bson:"acceptedAt"`
	DeliveredAt        int64 `json:"deliveredAt" bson:"deliveredAt"`
	ExpectedDeliveryAt int64 `json:"expectedDeliveryAt" bson:"expectedDeliveryAt"`
	UpdatedAt          int64 `json:"updatedAt" bson:"updatedAt"`
}
