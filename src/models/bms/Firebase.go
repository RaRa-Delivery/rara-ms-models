package bms

type FirebasePayload struct {
	To       string `json:"to"`
	Priority string `json:"priority"`
	Data     struct {
		ReqID            string `json:"reqId"`
		NotificationType string `json:"notificationType"`
		TitleEn          string `json:"titleEn"`
		TitleID          string `json:"titleId"`
		BodyEn           string `json:"bodyEn"`
		BodyID           string `json:"bodyId"`
		Data             string `json:"data"`
	} `json:"data"`
}

type FirebaseData struct {
	BatchID            string   `json:"batchId"`
	Time               int64    `json:"time"`
	ExpireTime         int64    `json:"expireTime"`
	AssignmentID       string   `json:"assignmentId"`
	Orders             []string `json:"orders"`
	Bsht               string   `json:"bsht"`
	IsManualAssignment bool     `json:"isManualAssignment"`
	IsBatchAssigned    bool     `json:"isBatchAssigned"`
}
