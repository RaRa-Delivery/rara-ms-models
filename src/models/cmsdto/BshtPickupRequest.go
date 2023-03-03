package cmsdto

type BshtPickupRequestNew struct {
	PageNo             int     `json:"pageNo"`
	PageSize           int     `json:"pageSize"`
	SortBy             string  `json:"sortBy"`
	SortDir            string  `json:"sortDir"`
	PickupLocationName string  `json:"pickupLocationName"`
	PickupLocationType string  `json:"pickupLocationType"`
	OperationRegionIds []int64 `json:"operationRegionIds"`
	BusinessAccountIds []int64 `json:"businessAccountIds"`
	IsActive           bool    `json:"isActive"`
	BshtTagsIds        []int64 `json:"bshtTagsIds"`
}
