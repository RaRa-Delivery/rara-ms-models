package models

type JotformSubmissionResponse struct {
	ResponseCode int     `json:"responseCode" bson:"responseCode"`
	Message      string  `json:"message" bson:"message"`
	Duration     string  `json:"duration" bson:"duration"`
	LimitLeft    int     `json:"limit-left" bson:"limit-left"`
	Content      Content `json:"content" bson:"content"`
}

type Content struct {
	Id        string      `json:"id" bson:"id"`
	Ip        string      `json:"ip" bson:"ip"`
	FormId    string      `json:"form_id" bson:"form_id"`
	CreatedAt string      `json:"created_at" bson:"created_at"`
	Status    string      `json:"status" bson:"status"`
	New       string      `json:"new" bson:"new"`
	Flag      string      `json:"flag" bson:"flag"`
	Notes     string      `json:"notes" bson:"notes"`
	UpdatedAt string      `json:"updated_at" bson:"updated_at"`
	Answers   interface{} `json:"answers" bson:"answers"`
}

type Answer struct {
	Name   string `json:"name" bson:"name"`
	Order  string `json:"order" bson:"order"`
	Text   string `json:"text" bson:"text"`
	Type   string `json:"type" bson:"type"`
	Answer string `json:"answer" bson:"answer"`
}
