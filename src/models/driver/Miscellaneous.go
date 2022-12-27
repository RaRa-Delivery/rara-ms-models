package models

type Miscellaneous struct {
	LastBatch    string `json:"lastBatch" bson:"lastBatch"`
	ActivatedVia string `json:"activatedVia" bson:"activatedVia"`
}
