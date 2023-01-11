package models

type Profile struct {
	ImageUrl         string `json:"image" bson:"image"`
	Kecamatan        string `json:"kecamatan" bson:"kecamatan"`
	EmergencyName    string `json:"emergencyName" bson:"emergencyName"`
	EmergencyContact string `json:"emergencyContact" bson:"emergencyContact"`
	Relationship     string `json:"relationship" bson:"relationship"`
	RabbitMobile     string `json:"rabbitMobile" bson:"rabbitMobile"`
	CreatedAt        int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt        int64  `json:"updatedAt" bson:"updatedAt"`
	Status           string `json:"status" bson:"status"`
}

type DocumentCompletedModel struct {
	Label       string `json:"label" bson:"label"`
	IsCompleted bool   `json:"isCompleted" bson:"isCompleted"`
}
