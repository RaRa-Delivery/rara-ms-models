package models

type PaginatedData struct {
	Content          interface{} `json:"content" bson:"content"`
	NumberOfElements int         `json:"numberOfElements" bson:"numberOfElements"`
	TotalElements    int         `json:"totalElements" bson:"totalElements"`
	TotalPages       int         `json"totalPages" bson:"totalPages"`
	CurrentPage      int64       `json:"currentPage" bson:currentPage"`
}
