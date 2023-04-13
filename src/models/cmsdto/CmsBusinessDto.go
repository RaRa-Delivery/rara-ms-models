package cmsdto

type CmsBusinessDto struct {
	Code    string       `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
	Data    BusinessData `json:"data,omitempty"`
}
type BusinessData struct {
	ID                   int    `json:"id,omitempty"`
	TenantToken          string `json:"tenantToken,omitempty"`
	CreatedBy            int    `json:"createdBy,omitempty"`
	CreatedOn            int64  `json:"createdOn,omitempty"`
	LastUpdatedBy        int    `json:"lastUpdatedBy,omitempty"`
	LastUpdatedOn        int64  `json:"lastUpdatedOn,omitempty"`
	Version              int    `json:"version,omitempty"`
	IsActive             bool   `json:"isActive,omitempty"`
	OfficialBusinessName string `json:"officialBusinessName,omitempty"`
	BrandName            string `json:"brandName,omitempty"`
	RegisteredAddress    string `json:"registeredAddress,omitempty"`
	RegisteredPhone      string `json:"registeredPhone,omitempty"`
	RegistrationID       string `json:"registrationId,omitempty"`
	NpwpID               string `json:"npwpId,omitempty"`
	DirectorName         string `json:"directorName,omitempty"`
	DirectorKtpID        string `json:"directorKtpId,omitempty"`
	OwnerIamID           int    `json:"ownerIamId,omitempty"`
	LogoFilePath         string `json:"logoFilePath,omitempty"`
	Status               bool   `json:"status,omitempty"`
	BusinessAccountCount int    `json:"businessAccountCount,omitempty"`
	Notes                string `json:"notes,omitempty"`
}
