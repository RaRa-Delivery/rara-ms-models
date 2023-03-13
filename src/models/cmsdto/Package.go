package cmsdto

type PackageNameDto struct {
	Code    string          `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	Data    PackageNameData `json:"data,omitempty"`
}
type Label struct {
	En string `json:"en,omitempty"`
	ID string `json:"id,omitempty"`
}
type PackageVariables struct {
	ID          int    `json:"id,omitempty"`
	TenantToken string `json:"tenantToken,omitempty"`
	Version     int    `json:"version,omitempty"`
	IsActive    bool   `json:"isActive,omitempty"`
	Reference   string `json:"reference,omitempty"`
	Name        string `json:"name,omitempty"`
	Label       Label  `json:"label,omitempty"`
	Type        string `json:"type,omitempty"`
	IsCurrency  bool   `json:"isCurrency,omitempty"`
	Value       string `json:"value,omitempty"`
}
type PackageDefinitions struct {
	ID               int                `json:"id,omitempty"`
	TenantToken      string             `json:"tenantToken,omitempty"`
	Version          int                `json:"version,omitempty"`
	IsActive         bool               `json:"isActive,omitempty"`
	Metric           string             `json:"metric,omitempty"`
	Reference        string             `json:"reference,omitempty"`
	PackageVariables []PackageVariables `json:"packageVariables,omitempty"`
	Formula          string             `json:"formula,omitempty"`
}
type PackageNameData struct {
	ID                 int                  `json:"id,omitempty"`
	TenantToken        string               `json:"tenantToken,omitempty"`
	Version            int                  `json:"version,omitempty"`
	IsActive           bool                 `json:"isActive,omitempty"`
	PackageName        string               `json:"packageName,omitempty"`
	WeightIndex        int                  `json:"weightIndex,omitempty"`
	PackageStatus      string               `json:"packageStatus,omitempty"`
	PackageDefinitions []PackageDefinitions `json:"packageDefinitions,omitempty"`
	BusinessID         int                  `json:"businessId,omitempty"`
}
