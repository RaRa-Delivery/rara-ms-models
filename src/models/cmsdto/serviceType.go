package cmsdto

type ServiceTypeObject struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    ServiceTypeData `json:"data"`
}
type ServiceType struct {
	ID            int    `json:"id"`
	TenantToken   string `json:"tenantToken"`
	CreatedBy     int    `json:"createdBy"`
	CreatedOn     int64  `json:"createdOn"`
	LastUpdatedBy int    `json:"lastUpdatedBy"`
	LastUpdatedOn int64  `json:"lastUpdatedOn"`
	Version       int    `json:"version"`
	IsActive      bool   `json:"isActive"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}
type DurationUnitLabel struct {
	ID               int     `json:"id"`
	TenantToken      string  `json:"tenantToken"`
	CreatedBy        int     `json:"createdBy"`
	CreatedOn        int64   `json:"createdOn"`
	LastUpdatedBy    int     `json:"lastUpdatedBy"`
	LastUpdatedOn    int64   `json:"lastUpdatedOn"`
	Version          int     `json:"version"`
	IsActive         bool    `json:"isActive"`
	Label            string  `json:"label"`
	CategoryName     string  `json:"categoryName"`
	CategoryKey      string  `json:"categoryKey"`
	OperationRegions []int64 `json:"operationRegions"`
	LabelStatus      string  `json:"labelStatus"`
}
type SLAType struct {
	ID                int               `json:"id"`
	TenantToken       string            `json:"tenantToken"`
	CreatedBy         int               `json:"createdBy"`
	CreatedOn         int64             `json:"createdOn"`
	LastUpdatedBy     int               `json:"lastUpdatedBy"`
	LastUpdatedOn     int64             `json:"lastUpdatedOn"`
	Version           int               `json:"version"`
	IsActive          bool              `json:"isActive"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Duration          int               `json:"duration"`
	DurationUnitLabel DurationUnitLabel `json:"durationUnitLabel"`
}
type OrderSLAStartLabel struct {
	ID               int     `json:"id"`
	TenantToken      string  `json:"tenantToken"`
	CreatedBy        int     `json:"createdBy"`
	CreatedOn        int64   `json:"createdOn"`
	LastUpdatedBy    int     `json:"lastUpdatedBy"`
	LastUpdatedOn    int64   `json:"lastUpdatedOn"`
	Version          int     `json:"version"`
	IsActive         bool    `json:"isActive"`
	Label            string  `json:"label"`
	CategoryName     string  `json:"categoryName"`
	CategoryKey      string  `json:"categoryKey"`
	OperationRegions []int64 `json:"operationRegions"`
	LabelStatus      string  `json:"labelStatus"`
}
type LinehaulLabel struct {
	ID               int     `json:"id"`
	TenantToken      string  `json:"tenantToken"`
	CreatedBy        int     `json:"createdBy"`
	CreatedOn        int64   `json:"createdOn"`
	LastUpdatedBy    int     `json:"lastUpdatedBy"`
	LastUpdatedOn    int64   `json:"lastUpdatedOn"`
	Version          int     `json:"version"`
	IsActive         bool    `json:"isActive"`
	Label            string  `json:"label"`
	CategoryName     string  `json:"categoryName"`
	CategoryKey      string  `json:"categoryKey"`
	OperationRegions []int64 `json:"operationRegions"`
	LabelStatus      string  `json:"labelStatus"`
}
type ServiceTypeData struct {
	ID                 int                `json:"id"`
	TenantToken        string             `json:"tenantToken"`
	CreatedBy          int                `json:"createdBy"`
	CreatedOn          int64              `json:"createdOn"`
	LastUpdatedBy      int                `json:"lastUpdatedBy"`
	LastUpdatedOn      int64              `json:"lastUpdatedOn"`
	Version            int                `json:"version"`
	IsActive           bool               `json:"isActive"`
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	OperationStartTime int64              `json:"operationStartTime"`
	OperationEndTime   int64              `json:"operationEndTime"`
	OrderCutoffTime    int                `json:"orderCutoffTime"`
	ServiceKey         string             `json:"serviceKey"`
	ServiceType        ServiceType        `json:"serviceType"`
	SLAType            SLAType            `json:"slaType"`
	OrderSLAStartLabel OrderSLAStartLabel `json:"orderSlaStartLabel"`
	PickupSLA          int                `json:"pickupSla"`
	LinehaulLabel      LinehaulLabel      `json:"linehaulLabel"`
	OperationRegions   []int              `json:"operationRegions"`
	RoutingFrequency   int                `json:"routingFrequency"`
	ServiceStatus      string             `json:"serviceStatus"`
}
