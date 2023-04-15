package cmsdto

type OperationRegionDto struct {
	Code    string              `json:"code"`
	Message string              `json:"message"`
	Data    OperationRegionData `json:"data"`
}
type OperationRegionData struct {
	ID          int    `json:"id"`
	TenantToken string `json:"tenantToken"`
	Version     int    `json:"version"`
	IsActive    bool   `json:"isActive"`
	RegionName  string `json:"regionName"`
}
