package cmsdto

type ZoneDataRequired struct {
	ID         int64  `json:"id"`
	ZoneName   string `json:"zoneName"`
	RegionId   int64  `json:"regionId"`
	RegionName string `json:"regionName"`
}
