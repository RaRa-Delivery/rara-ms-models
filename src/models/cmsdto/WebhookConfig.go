package cmsdto

type WebhookConfigDto struct {
	Code    string              `json:"code,omitempty"`
	Message string              `json:"message,omitempty"`
	Data    []WebhookConfigData `json:"data,omitempty"`
}
type OrderStatus struct {
	ID          int    `json:"id,omitempty"`
	TenantToken string `json:"tenantToken,omitempty"`
	Version     int    `json:"version,omitempty"`
	IsActive    bool   `json:"isActive,omitempty"`
	Name        string `json:"name,omitempty"`
	Code        string `json:"code,omitempty"`
}
type WebhookConfigData struct {
	ID                int         `json:"id,omitempty"`
	OrderStatus       OrderStatus `json:"orderStatus,omitempty"`
	BusinessAccountID int         `json:"businessAccountId,omitempty"`
	WebhookStatusName string      `json:"webhookStatusName,omitempty"`
	Status            string      `json:"status,omitempty"`
}

type NotificationConfigDto struct {
	Code    string                   `json:"code,omitempty"`
	Message string                   `json:"message,omitempty"`
	Data    []NotificationConfigData `json:"data,omitempty"`
}

type NotificationConfigData struct {
	ID                          int                         `json:"id,omitempty"`
	OrderStatus                 OrderStatus                 `json:"orderStatus,omitempty"`
	BusinessAccountID           int                         `json:"businessAccountId,omitempty"`
	Status                      string                      `json:"status,omitempty"`
	NotificationTargetType      string                      `json:"notificationTargetType,omitempty"`
	ChatbotNotificationTemplate ChatbotNotificationTemplate `json:"chatbotNotificationTemplate,omitempty"`
}

type Params struct {
	Num1 string `json:"1,omitempty"`
	Num2 string `json:"2,omitempty"`
	Num3 string `json:"3,omitempty"`
	Num4 string `json:"4,omitempty"`
	Num5 string `json:"5,omitempty"`
	Num6 string `json:"6,omitempty"`
}
type ChatbotNotificationTemplate struct {
	TemplateID      string `json:"templateId,omitempty"`
	Params          Params `json:"params,omitempty"`
	OrderStatusCode string `json:"orderStatusCode,omitempty"`
}
