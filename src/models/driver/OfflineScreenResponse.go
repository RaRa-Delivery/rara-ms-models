package models

type OfflineScreenResponse struct {
	Header      string `json:"header" bson:"header"`
	Text        string `json:"text" bson:"text"`
	LeftEarning struct {
		EarningType string  `json:"earningType" bson:"earningType"`
		Earning     Earning `json:"earning" bson:"earning"`
	} `json:"leftEarning" bson:"leftEarning"`
	RightEarning struct {
		EarningType string  `json:"earningType" bson:"earningType"`
		Earning     Earning `json:"earning" bson:"earning"`
	} `json:"rightEarning" bson:"rightEarning"`
}
