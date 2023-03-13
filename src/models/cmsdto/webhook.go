package cmsdto

type WebhookBuilder struct {
	Code    interface{}   `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Data    []WebhookData `json:"data,omitempty"`
}
type CallBackHeaders struct {
	ContentType   string `json:"Content-Type,omitempty"`
	Authorization string `json:"Authorization,omitempty"`
}
type CallBackBody struct {
}
type WebhookData struct {
	Id                  int64              `json:"id,omitempty"`
	TenantToken         string             `json:"tenantToken,omitempty"`
	CreatedBy           int                `json:"createdBy,omitempty"`
	CreatedOn           int64              `json:"createdOn,omitempty"`
	LastUpdatedBy       int                `json:"lastUpdatedBy,omitempty"`
	LastUpdatedOn       int64              `json:"lastUpdatedOn,omitempty"`
	Version             int                `json:"version,omitempty"`
	IsActive            bool               `json:"isActive,omitempty"`
	WebhookTemplate     string             `json:"webhookTemplate,omitempty"`
	CallBackURL         string             `json:"callBackUrl,omitempty"`
	CallBackQuery       string             `json:"callBackQuery,omitempty"`
	CallBackHeaders     map[string]string  `json:"callBackHeaders,omitempty"`
	CallBackBody        CallBackBody       `json:"callBackBody,omitempty"`
	WebhookCallbackType string             `json:"webhookCallbackType,omitempty"`
	AuthTokenSource     string             `json:"authTokenSource,omitempty"`
	BusinessID          int64              `json:"businessId,omitempty"`
	Purpose             []WebhookStatusMap `json:"purpose,omitempty"`
}

type WebhookStatusMap struct {
	Code           string `json:"code,omitempty" bson:"code"`
	BusinessStatus string `json:"businessStatus,omitempty" bson:"businessStatus"`
}
