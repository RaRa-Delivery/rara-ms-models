package cmsdto

type CmsCityObj struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    CmsCityData `json:"data"`
}
type CmsCityData struct {
	ID          int64  `json:"id"`
	TenantToken string `json:"tenantToken"`
	Version     int    `json:"version"`
	IsActive    bool   `json:"isActive"`
	RegionName  string `json:"regionName"`
}
