package cmsdto

type CmsPickupDto struct {
	Code    string     `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
	Data    PickupData `json:"data,omitempty"`
}
type PickupZoneID struct {
	ID       int    `json:"id,omitempty"`
	ZoneID   string `json:"zoneId,omitempty"`
	ZoneName string `json:"zoneName,omitempty"`
	ZoneType string `json:"zoneType,omitempty"`
}
type DropOffZoneIds struct {
	ID       int    `json:"id,omitempty"`
	ZoneID   string `json:"zoneId,omitempty"`
	ZoneName string `json:"zoneName,omitempty"`
	ZoneType string `json:"zoneType,omitempty"`
}
type MappedAccounts struct {
	BusinessAccountID int              `json:"businessAccountId,omitempty"`
	PickupSLA         int              `json:"pickupSla,omitempty"`
	CapacityIndex     int              `json:"capacityIndex,omitempty"`
	PickupZoneID      PickupZoneID     `json:"pickupZoneId,omitempty"`
	DropOffZoneIds    []DropOffZoneIds `json:"dropOffZoneIds,omitempty"`
}
type PickupData struct {
	PickupLocationID   int              `json:"pickupLocationId,omitempty"`
	PickupLocationName string           `json:"pickupLocationName,omitempty"`
	PickupLocationType string           `json:"pickupLocationType,omitempty"`
	Address            string           `json:"address,omitempty"`
	Latitude           float64          `json:"latitude,omitempty"`
	Longitude          float64          `json:"longitude,omitempty"`
	OperationRegionID  int              `json:"operationRegionId,omitempty"`
	ProvinceID         int              `json:"provinceId,omitempty"`
	MappedAccounts     []MappedAccounts `json:"mappedAccounts,omitempty"`
}
