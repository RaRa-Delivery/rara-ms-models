package cmsdto

type PackageNameDto struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    PackageNameData `json:"data"`
}
type Label struct {
	En string `json:"en"`
	ID string `json:"id"`
}
type PackageVariables struct {
	ID          int    `json:"id"`
	TenantToken string `json:"tenantToken"`
	Version     int    `json:"version"`
	IsActive    bool   `json:"isActive"`
	Reference   string `json:"reference"`
	Name        string `json:"name"`
	Label       Label  `json:"label"`
	Type        string `json:"type"`
	IsCurrency  bool   `json:"isCurrency"`
	Value       string `json:"value"`
}
type PackageDefinitions struct {
	ID               int                `json:"id"`
	TenantToken      string             `json:"tenantToken"`
	Version          int                `json:"version"`
	IsActive         bool               `json:"isActive"`
	Metric           string             `json:"metric"`
	Reference        string             `json:"reference"`
	PackageVariables []PackageVariables `json:"packageVariables"`
	Formula          string             `json:"formula"`
}
type PackageNameData struct {
	ID                 int                  `json:"id"`
	TenantToken        string               `json:"tenantToken"`
	Version            int                  `json:"version"`
	IsActive           bool                 `json:"isActive"`
	PackageName        string               `json:"packageName"`
	WeightIndex        int                  `json:"weightIndex"`
	PackageStatus      string               `json:"packageStatus"`
	PackageDefinitions []PackageDefinitions `json:"packageDefinitions"`
	BusinessID         int                  `json:"businessId"`
}
