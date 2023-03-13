package cmsdto

type LocationDto struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Data    LocationData `json:"data,omitempty"`
}
type LocationData struct {
	ID            int    `json:"id"`
	TenantToken   string `json:"tenantToken"`
	CreatedBy     int    `json:"createdBy"`
	CreatedOn     int64  `json:"createdOn"`
	LastUpdatedBy int    `json:"lastUpdatedBy"`
	LastUpdatedOn int64  `json:"lastUpdatedOn"`
	Version       int    `json:"version"`
	IsActive      bool   `json:"isActive"`
	Zipcode       string `json:"zipcode"`
	Village       string `json:"village"`
	SubDistrict   string `json:"subDistrict"`
	District      string `json:"district"`
	Province      string `json:"province"`
	ZipcodeType   string `json:"zipcodeType"`
}

type OperationRegionByPostalCode struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		ID          int    `json:"id,omitempty"`
		TenantToken string `json:"tenantToken,omitempty"`
		Version     int    `json:"version,omitempty"`
		IsActive    bool   `json:"isActive,omitempty"`
		RegionName  string `json:"regionName,omitempty"`
	} `json:"data,omitempty"`
}
