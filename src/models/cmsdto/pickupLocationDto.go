package cmsdto

type CmsPickupDto struct {
	Code    string     `json:"code"`
	Message string     `json:"message"`
	Data    PickupData `json:"data"`
}
type PickupZoneID struct {
	ID       int    `json:"id"`
	ZoneID   string `json:"zoneId"`
	ZoneName string `json:"zoneName"`
	ZoneType string `json:"zoneType"`
}
type DropOffZoneIds struct {
	ID       int    `json:"id"`
	ZoneID   string `json:"zoneId"`
	ZoneName string `json:"zoneName"`
	ZoneType string `json:"zoneType"`
}
type MappedAccounts struct {
	BusinessAccountID int              `json:"businessAccountId"`
	PickupSLA         int              `json:"pickupSla"`
	CapacityIndex     int              `json:"capacityIndex"`
	PickupZoneID      PickupZoneID     `json:"pickupZoneId"`
	DropOffZoneIds    []DropOffZoneIds `json:"dropOffZoneIds"`
}
type PickupData struct {
	PickupLocationID   int              `json:"pickupLocationId"`
	PickupLocationName string           `json:"pickupLocationName"`
	PickupLocationType string           `json:"pickupLocationType"`
	Address            string           `json:"address"`
	Latitude           float64          `json:"latitude"`
	Longitude          float64          `json:"longitude"`
	OperationRegionID  int              `json:"operationRegionId"`
	ProvinceID         int              `json:"provinceId"`
	MappedAccounts     []MappedAccounts `json:"mappedAccounts"`
}
