package cmsdto

type CommonLabel struct {
	Code    interface{}       `json:"code,omitempty"`
	Message string            `json:"message,omitempty"`
	Data    []CommonLableData `json:"data,omitempty"`
}
type CommonLableData struct {
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
