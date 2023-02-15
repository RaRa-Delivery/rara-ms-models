package cmsdto

type CmsProvinceDto struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    CmsProvinceData `json:"data"`
}
type CmsProvinceType struct {
	ID           int    `json:"id"`
	TenantToken  string `json:"tenantToken"`
	Version      int    `json:"version"`
	IsActive     bool   `json:"isActive"`
	ProvinceType string `json:"provinceType"`
}
type CmsOperationRegion struct {
	ID          int    `json:"id"`
	TenantToken string `json:"tenantToken"`
	Version     int    `json:"version"`
	IsActive    bool   `json:"isActive"`
	RegionName  string `json:"regionName"`
}
type CmsOperationCity struct {
	ID              int                `json:"id"`
	TenantToken     string             `json:"tenantToken"`
	Version         int                `json:"version"`
	IsActive        bool               `json:"isActive"`
	CityName        string             `json:"cityName"`
	OperationRegion CmsOperationRegion `json:"operationRegion"`
}
type CmsProvinceData struct {
	ID                int                `json:"id"`
	TenantToken       string             `json:"tenantToken"`
	Version           int                `json:"version"`
	IsActive          bool               `json:"isActive"`
	Name              string             `json:"name"`
	ProvinceType      CmsProvinceType    `json:"provinceType"`
	ProvinceTypeID    int                `json:"provinceTypeId"`
	OperationCity     CmsOperationCity   `json:"operationCity"`
	OperationCityID   int                `json:"operationCityId"`
	OperationRegion   CmsOperationRegion `json:"operationRegion"`
	OperationRegionID int                `json:"operationRegionId"`
}
