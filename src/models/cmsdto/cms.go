package cmsdto

type CmsBasic struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type CmsEntity struct {
	ID      int `json:"id"`
	Version int `json:"version"`
}

type DeliveryFeeSchemes struct {
	ID      int    `json:"id"`
	Value   string `json:"value"`
	Version int    `json:"version"`
}
type ServiceAndPricingConfig struct {
	DeliveryService    CmsEntity            `json:"deliveryService"`
	DeliveryFeeSchemes []DeliveryFeeSchemes `json:"deliveryFeeSchemes"`
}
type Bsht struct {
	IsEnabled bool        `json:"isEnabled"`
	Tags      []CmsEntity `json:"tags"`
}
type SurchargeSchemes struct {
	ID      int    `json:"id"`
	Version int    `json:"version"`
	Value   string `json:"value"`
}
type BusinessAccountProperties struct {
	IsLineHaul       bool               `json:"isLineHaul"`
	Bsht             Bsht               `json:"bsht"`
	SurchargeSchemes []SurchargeSchemes `json:"surchargeSchemes"`
}

type PackageTypeScheme struct {
	ID      int    `json:"id"`
	Value   string `json:"value"`
	Version int    `json:"version"`
}
type DeliveryPackaging struct {
	PackageName       CmsEntity         `json:"packageName"`
	WeightIndex       string            `json:"weightIndex"`
	PackageTypeScheme PackageTypeScheme `json:"packageTypeScheme"`
}

type BillingInfo struct {
	BillingAddress string    `json:"billingAddress"`
	InvoiceDate    string    `json:"invoiceDate"`
	TermsOfPayment CmsEntity `json:"termsOfPayment"`
}
type Data struct {
	ID                        int                       `json:"id"`
	TenantToken               string                    `json:"tenantToken"`
	CreatedBy                 int                       `json:"createdBy"`
	CreatedOn                 int                       `json:"createdOn"`
	LastUpdatedBy             int                       `json:"lastUpdatedBy"`
	LastUpdatedOn             int64                     `json:"lastUpdatedOn"`
	Version                   int                       `json:"version"`
	IsActive                  bool                      `json:"isActive"`
	BusinessID                int                       `json:"businessId"`
	AccountName               string                    `json:"accountName"`
	PicName                   string                    `json:"picName"`
	PicPhoneNumber            string                    `json:"picPhoneNumber"`
	PicEmail                  interface{}               `json:"picEmail"`
	OperationRegionId         int                       `json:"operationRegionId"`
	Notes                     interface{}               `json:"notes"`
	AccountStatus             string                    `json:"accountStatus"`
	SpecialHandlingRequired   bool                      `json:"specialHandlingRequired"`
	LinehaulRequired          bool                      `json:"linehaulRequired"`
	ServiceAndPricingConfig   []ServiceAndPricingConfig `json:"serviceAndPricingConfig"`
	BusinessAccountProperties BusinessAccountProperties `json:"businessAccountProperties"`
	DeliveryPackaging         []DeliveryPackaging       `json:"deliveryPackaging"`
	BillingInfo               BillingInfo               `json:"billingInfo"`
	RegistrationStep          interface{}               `json:"registrationStep"`
	SeqID                     interface{}               `json:"seqId"`
}
