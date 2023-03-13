package cmsdto

type CmsProvinceDto struct {
	Code    string          `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	Data    CmsProvinceData `json:"data,omitempty"`
}
type CmsProvinceType struct {
	ID           int    `json:"id,omitempty"`
	TenantToken  string `json:"tenantToken,omitempty"`
	Version      int    `json:"version,omitempty"`
	IsActive     bool   `json:"isActive,omitempty"`
	ProvinceType string `json:"provinceType,omitempty"`
}
type CmsOperationRegion struct {
	ID          int    `json:"id,omitempty"`
	TenantToken string `json:"tenantToken,omitempty"`
	Version     int    `json:"version,omitempty"`
	IsActive    bool   `json:"isActive,omitempty"`
	RegionName  string `json:"regionName,omitempty"`
}
type CmsOperationCity struct {
	ID              int                `json:"id,omitempty"`
	TenantToken     string             `json:"tenantToken,omitempty"`
	Version         int                `json:"version,omitempty"`
	IsActive        bool               `json:"isActive,omitempty"`
	CityName        string             `json:"cityName,omitempty"`
	OperationRegion CmsOperationRegion `json:"operationRegion,omitempty"`
}
type CmsProvinceData struct {
	ID                int                `json:"id,omitempty"`
	TenantToken       string             `json:"tenantToken,omitempty"`
	Version           int                `json:"version,omitempty"`
	IsActive          bool               `json:"isActive,omitempty"`
	Name              string             `json:"name,omitempty"`
	ProvinceType      CmsProvinceType    `json:"provinceType,omitempty"`
	ProvinceTypeID    int                `json:"provinceTypeId,omitempty"`
	OperationCity     CmsOperationCity   `json:"operationCity,omitempty"`
	OperationCityID   int                `json:"operationCityId,omitempty"`
	OperationRegion   CmsOperationRegion `json:"operationRegion,omitempty"`
	OperationRegionID int                `json:"operationRegionId,omitempty"`
}
