package cmsdto

type ZoneDto struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Data    ZoneData `json:"data"`
}

type ZoneData struct {
	ID              int64  `json:"id"`
	TenantToken     string `json:"tenantToken"`
	CreatedBy       int    `json:"createdBy"`
	CreatedOn       int64  `json:"createdOn"`
	LastUpdatedBy   int    `json:"lastUpdatedBy"`
	LastUpdatedOn   int64  `json:"lastUpdatedOn"`
	Version         int    `json:"version"`
	IsActive        bool   `json:"isActive"`
	ZoneID          string `json:"zoneId"`
	ZoneName        string `json:"zoneName"`
	ZoneStatus      string `json:"zoneStatus"`
	ZoneType        string `json:"zoneType"`
	HasOverLap      bool   `json:"hasOverLap"`
	OperationRegion struct {
		ID            int64  `json:"id"`
		TenantToken   string `json:"tenantToken"`
		CreatedBy     int    `json:"createdBy"`
		CreatedOn     int64  `json:"createdOn"`
		LastUpdatedBy int    `json:"lastUpdatedBy"`
		LastUpdatedOn int64  `json:"lastUpdatedOn"`
		Version       int    `json:"version"`
		IsActive      bool   `json:"isActive"`
		RegionName    string `json:"regionName"`
	} `json:"operationRegion"`
	OperationRegionID int `json:"operationRegionId"`
	ZoneData          struct {
		Kecamatan []int `json:"Kecamatan"`
		City      []int `json:"city"`
	} `json:"zoneData"`
}
