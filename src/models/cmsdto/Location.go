package cmsdto

type LocationDto struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Data    LocationData `json:"data"`
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
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID          int    `json:"id"`
		TenantToken string `json:"tenantToken"`
		Version     int    `json:"version"`
		IsActive    bool   `json:"isActive"`
		RegionName  string `json:"regionName"`
	} `json:"data"`
}
