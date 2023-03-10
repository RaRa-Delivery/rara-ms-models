package cmsdto

type WebhookConfigDto struct {
	Code    string              `json:"code"`
	Message string              `json:"message"`
	Data    []WebhookConfigData `json:"data"`
}
type OrderStatus struct {
	ID          int    `json:"id"`
	TenantToken string `json:"tenantToken"`
	Version     int    `json:"version"`
	IsActive    bool   `json:"isActive"`
	Name        string `json:"name"`
	Code        string `json:"code"`
}
type WebhookConfigData struct {
	ID                int         `json:"id"`
	OrderStatus       OrderStatus `json:"orderStatus"`
	BusinessAccountID int         `json:"businessAccountId"`
	WebhookStatusName string      `json:"webhookStatusName"`
	Status            string      `json:"status"`
}

type NotificationConfigDto struct {
	Code    string                   `json:"code"`
	Message string                   `json:"message"`
	Data    []NotificationConfigData `json:"data"`
}

type NotificationConfigData struct {
	ID                     int         `json:"id"`
	OrderStatus            OrderStatus `json:"orderStatus"`
	BusinessAccountID      int         `json:"businessAccountId"`
	Status                 string      `json:"status"`
	NotificationTargetType string      `json:"notificationTargetType"`
}
