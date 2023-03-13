package cmsdto

type PickupByDropoffDto struct {
	Code    string              `json:"code,omitempty"`
	Message string              `json:"message,omitempty"`
	Data    PickupByDropoffData `json:"data,omitempty"`
}
type PickupByDropoffData struct {
	PickupLocationId   int64       `json:"pickupLocationId,omitempty"`
	PickupLocationName string      `json:"pickupLocationName,omitempty"`
	PickupLocationType string      `json:"pickupLocationType,omitempty"`
	Address            string      `json:"address,omitempty"`
	Latitude           float64     `json:"latitude,omitempty"`
	Longitude          float64     `json:"longitude,omitempty"`
	OperationRegionId  int64       `json:"operationRegionId,omitempty"`
	ProvinceId         int64       `json:"provinceId,omitempty"`
	PickupZone         ZoneDetails `json:"pickupZone,omitempty"`
	BusinessAccountId  int64       `json:"businessAccountId,omitempty"`
}

type ZoneDetails struct {
	ID       int    `json:"id,omitempty"`
	ZoneID   string `json:"zoneId,omitempty"`
	ZoneName string `json:"zoneName,omitempty"`
	ZoneType string `json:"zoneType,omitempty"`
}

type DropoffZoneDto struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    ZoneDetails `json:"data,omitempty"`
}
