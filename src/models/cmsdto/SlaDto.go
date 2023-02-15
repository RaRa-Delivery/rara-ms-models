package cmsdto

type SLAservice struct {
	Type         string   `json:"type" bson:"type"`
	ID           string   `json:"id" bson:"id"`
	Name         string   `json:"name" bson:"name"`
	Instant      string   `json:"instant" bson:"instant"`
	SLAValue     int32    `json:"slaValue" bson:"slaValue"`
	SLAUnit      string   `json:"unit" bson:"unit"`
	StartTime    string   `json:"startTime" bson:"startTime"`
	EndTime      string   `json:"endTime" bson:"endTime"`
	CutOffTime   string   `json:"cutOffTime" bson:"cutOffTime"`
	MaxDistance  int32    `json:"maxDistance" bson:"maxDistance"`
	DistanceUnit string   `json:"distanceUnit" bson:"distanceUnit"`
	PickupPoints []string `json:"pickupPoints" bson:"pickupPoints"`
	ServiceId    int64    `json:"serviceId" bson:"serviceId"`
}
