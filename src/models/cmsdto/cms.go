package cmsdto

type CmsBasic struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    Data   `json:"data,omitempty"`
}

type CmsEntity struct {
	ID      int `json:"id,omitempty"`
	Version int `json:"version,omitempty"`
}

type DeliveryFeeSchemes struct {
	ID      int    `json:"id,omitempty"`
	Value   string `json:"value,omitempty"`
	Version int    `json:"version,omitempty"`
}
type ServiceAndPricingConfig struct {
	DeliveryService    CmsEntity            `json:"deliveryService,omitempty"`
	DeliveryFeeSchemes []DeliveryFeeSchemes `json:"deliveryFeeSchemes,omitempty"`
}
type Bsht struct {
	IsEnabled bool        `json:"isEnabled,omitempty"`
	Tags      []CmsEntity `json:"tags,omitempty"`
}
type SurchargeSchemes struct {
	ID      int    `json:"id,omitempty"`
	Version int    `json:"version,omitempty"`
	Value   string `json:"value,omitempty"`
}
type BusinessAccountProperties struct {
	IsLineHaul       bool               `json:"isLineHaul,omitempty"`
	Bsht             Bsht               `json:"bsht,omitempty"`
	SurchargeSchemes []SurchargeSchemes `json:"surchargeSchemes,omitempty"`
}

type PackageTypeScheme struct {
	ID      int    `json:"id,omitempty"`
	Value   string `json:"value,omitempty"`
	Version int    `json:"version,omitempty"`
}
type DeliveryPackaging struct {
	PackageName       CmsEntity         `json:"packageName,omitempty"`
	WeightIndex       string            `json:"weightIndex,omitempty"`
	PackageTypeScheme PackageTypeScheme `json:"packageTypeScheme,omitempty"`
}

type BillingInfo struct {
	BillingAddress string    `json:"billingAddress,omitempty"`
	InvoiceDate    string    `json:"invoiceDate,omitempty"`
	TermsOfPayment CmsEntity `json:"termsOfPayment,omitempty"`
}
type Data struct {
	ID                        int                       `json:"id,omitempty"`
	TenantToken               string                    `json:"tenantToken,omitempty"`
	CreatedBy                 int                       `json:"createdBy,omitempty"`
	CreatedOn                 int                       `json:"createdOn,omitempty"`
	LastUpdatedBy             int                       `json:"lastUpdatedBy,omitempty"`
	LastUpdatedOn             int64                     `json:"lastUpdatedOn,omitempty"`
	Version                   int                       `json:"version,omitempty"`
	IsActive                  bool                      `json:"isActive,omitempty"`
	BusinessID                int                       `json:"businessId,omitempty"`
	AccountName               string                    `json:"accountName,omitempty"`
	PicName                   string                    `json:"picName,omitempty"`
	PicPhoneNumber            string                    `json:"picPhoneNumber,omitempty"`
	PicEmail                  interface{}               `json:"picEmail,omitempty"`
	OperationRegionId         int                       `json:"operationRegionId,omitempty"`
	Notes                     interface{}               `json:"notes,omitempty"`
	AccountStatus             string                    `json:"accountStatus,omitempty"`
	SpecialHandlingRequired   bool                      `json:"specialHandlingRequired,omitempty"`
	LinehaulRequired          bool                      `json:"linehaulRequired,omitempty"`
	ServiceAndPricingConfig   []ServiceAndPricingConfig `json:"serviceAndPricingConfig,omitempty"`
	BusinessAccountProperties BusinessAccountProperties `json:"businessAccountProperties,omitempty"`
	DeliveryPackaging         []DeliveryPackaging       `json:"deliveryPackaging,omitempty"`
	BillingInfo               BillingInfo               `json:"billingInfo,omitempty"`
	RegistrationStep          interface{}               `json:"registrationStep,omitempty"`
	SeqID                     interface{}               `json:"seqId,omitempty"`
}
