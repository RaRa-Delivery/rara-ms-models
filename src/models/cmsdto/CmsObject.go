package cmsdto

type CmsObject struct {
	BusinessDetails         BusinessDetails        `json:"businessDetails"`
	LinehaulRequired        bool                   `json:"linehaulRequired"`
	SpecialHandlingRequired bool                   `json:"specialHandlingRequired"`
	OperationRegions        CmsRegion              `json:"operationRegions"`
	Bsht                    BshtDetails            `json:"bshtDetails"`
	CancellationConditions  CancellationConditions `json:"cancellationConditions"`
	Webhooks                []WebhookData          `json:"webhooks"`
	DeliveryPackaging       []PackageData          `json:"deliveryPackaging"`
	OrderJourneyDetails     []DeliveryStatusDto    `json:"orderJourneyDetails"`
	CmsChatbotService       CmsChatbotService      `json:"chatbot"`
	DeliveryFee             []DeliveryFeeDto       `json:"deliveryFee"`
	SLAservices             []SLAservice           `json:"slaService"`
}

type BusinessDetails struct {
	BusinessId     int64  `json:"businessId"`
	AccountId      int64  `json:"accountId"`
	AccountName    string `json:"accountName"`
	PicName        string `json:"picName"`
	PicPhoneNumber string `json:"picPhoneNumber"`
}

type CmsRegion struct {
	Id         int64  `json:"id"`
	RegionName string `json:"regionName"`
}

type BshtDetails struct {
	IsEnabled bool  `json:"isEnabled"`
	Tags      []Tag `json:"tags"`
}

type Tag struct {
	Id    int64  `json:"id"`
	Label string `json:"label"`
}

type CancellationConditions struct {
	StatusIndex int      `json:"statusIndex"`
	IsEligible  bool     `json:"isEligible"`
	Type        []string `json:"type"`
	By          []string `json:"by"`
	Condition   string   `json:"condition"`
}

type CmsChatbotService struct {
	Language      string                   `json:"language" bson:"language"`
	Notifications []CmsChatbotNotification `json:"notifications" bson:"notifications"`
	PostBackUrl   string                   `json:"postBackUrl" bson:"postBackUrl"`
}

type CmsChatbotNotification struct {
	NotificationType string             `json:"notificationType" bson:"notificationType"`
	TemplateId       string             `json:"templateId" bson:"templateId"`
	Params           []CmsChatbotParams `json:"params" bson:"params"`
}

type CmsChatbotParams struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}
