package cmsdto

type WebhookBuilder struct {
	Code    interface{}   `json:"code"`
	Message string        `json:"message"`
	Data    []WebhookData `json:"data"`
}
type CallBackHeaders struct {
	ContentType   string `json:"Content-Type"`
	Authorization string `json:"Authorization"`
}
type CallBackBody struct {
}
type WebhookData struct {
	Id                  int64              `json:"id"`
	TenantToken         string             `json:"tenantToken"`
	CreatedBy           int                `json:"createdBy"`
	CreatedOn           int64              `json:"createdOn"`
	LastUpdatedBy       int                `json:"lastUpdatedBy"`
	LastUpdatedOn       int64              `json:"lastUpdatedOn"`
	Version             int                `json:"version"`
	IsActive            bool               `json:"isActive"`
	WebhookTemplate     string             `json:"webhookTemplate"`
	CallBackURL         string             `json:"callBackUrl"`
	CallBackQuery       string             `json:"callBackQuery"`
	CallBackHeaders     map[string]string  `json:"callBackHeaders"`
	CallBackBody        CallBackBody       `json:"callBackBody"`
	WebhookCallbackType string             `json:"webhookCallbackType"`
	AuthTokenSource     string             `json:"authTokenSource"`
	BusinessID          int64              `json:"businessId"`
	Purpose             []WebhookStatusMap `json:"purpose"`
}

type WebhookStatusMap struct {
	Code           string `json:"code"`
	BusinessStatus string `json:"businessStatus"`
}
