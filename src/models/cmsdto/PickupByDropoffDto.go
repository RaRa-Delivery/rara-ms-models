package cmsdto

type PickupByDropoffDto struct {
	Code    string              `json:"code"`
	Message string              `json:"message"`
	Data    PickupByDropoffData `json:"data"`
}
type PickupByDropoffData struct {
	PickupLocationId   int64       `json:"pickupLocationId"`
	PickupLocationName string      `json:"pickupLocationName"`
	PickupLocationType string      `json:"pickupLocationType"`
	Address            string      `json:"address"`
	Latitude           float64     `json:"latitude"`
	Longitude          float64     `json:"longitude"`
	OperationRegionId  int64       `json:"operationRegionId"`
	ProvinceId         int64       `json:"provinceId"`
	PickupZone         ZoneDetails `json:"pickupZone"`
	BusinessAccountId  int64       `json:"businessAccountId"`
}

type ZoneDetails struct {
	ID       int    `json:"id"`
	ZoneID   string `json:"zoneId"`
	ZoneName string `json:"zoneName"`
	ZoneType string `json:"zoneType"`
}

type DropoffZoneDto struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    ZoneDetails `json:"data"`
}
