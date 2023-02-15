package cmsdto

type BshtDto struct {
	Code    interface{}       `json:"code"`
	Message string            `json:"message"`
	Data    []CommonLabelData `json:"data"`
}

type PackDto struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID            int    `json:"id"`
		TenantToken   string `json:"tenantToken"`
		CreatedBy     int    `json:"createdBy"`
		CreatedOn     int64  `json:"createdOn"`
		LastUpdatedBy int    `json:"lastUpdatedBy"`
		LastUpdatedOn int64  `json:"lastUpdatedOn"`
		Version       int    `json:"version"`
		IsActive      bool   `json:"isActive"`
		Name          string `json:"name"`
		IsCombinable  bool   `json:"isCombinable"`
		Definition    string `json:"definition"`
	} `json:"data"`
}

type CommonLabelData struct {
	ID               int    `json:"id"`
	TenantToken      string `json:"tenantToken"`
	CreatedBy        int    `json:"createdBy"`
	CreatedOn        int64  `json:"createdOn"`
	LastUpdatedBy    int    `json:"lastUpdatedBy"`
	LastUpdatedOn    int64  `json:"lastUpdatedOn"`
	Version          int    `json:"version"`
	IsActive         bool   `json:"isActive"`
	Label            string `json:"label"`
	CategoryName     string `json:"categoryName"`
	CategoryKey      string `json:"categoryKey"`
	OperationRegions []int  `json:"operationRegions"`
	LabelStatus      string `json:"labelStatus"`
}

type BshtPickupRequest struct {
	PageNo             int     `json:"pageNo"`
	PageSize           int     `json:"pageSize"`
	SortBy             string  `json:"sortBy"`
	SortDir            string  `json:"sortDir"`
	PickupLocationName string  `json:"pickupLocationName"`
	PickupLocationType string  `json:"pickupLocationType"`
	OperationRegionIds []int64 `json:"operationRegionIds"`
	BusinessAccountIds []int64 `json:"businessAccountIds"`
	IsActive           bool    `json:"isActive"`
	BshtTagsIds        []int64 `json:"bshtTagsIds"`
}

type BshtPickup struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Content []BshtPickupData `json:"content"`
	} `json:"data"`
	Pageable struct {
		Sort struct {
			Empty    bool `json:"empty"`
			Sorted   bool `json:"sorted"`
			Unsorted bool `json:"unsorted"`
		} `json:"sort"`
		Offset     int  `json:"offset"`
		PageNumber int  `json:"pageNumber"`
		PageSize   int  `json:"pageSize"`
		Paged      bool `json:"paged"`
		Unpaged    bool `json:"unpaged"`
	} `json:"pageable"`
	TotalPages    int  `json:"totalPages"`
	TotalElements int  `json:"totalElements"`
	Last          bool `json:"last"`
	Size          int  `json:"size"`
	Number        int  `json:"number"`
	Sort          struct {
		Empty    bool `json:"empty"`
		Sorted   bool `json:"sorted"`
		Unsorted bool `json:"unsorted"`
	} `json:"sort"`
	NumberOfElements int  `json:"numberOfElements"`
	First            bool `json:"first"`
	Empty            bool `json:"empty"`
}

type BshtPickupData struct {
	ID                 int64  `json:"id"`
	TenantToken        string `json:"tenantToken"`
	Version            int    `json:"version"`
	IsActive           bool   `json:"isActive"`
	PickupLocationName string `json:"pickupLocationName"`
	OperationRegionID  struct {
		ID          int64  `json:"id"`
		TenantToken string `json:"tenantToken"`
		Version     int    `json:"version"`
		IsActive    bool   `json:"isActive"`
		RegionName  string `json:"regionName"`
	} `json:"operationRegionId"`
	Address                       string  `json:"address"`
	Latitude                      float64 `json:"latitude"`
	Longitude                     float64 `json:"longitude"`
	PickupLocationType            string  `json:"pickupLocationType"`
	BusinessAccountConfigurations []struct {
		ID                 int64   `json:"id"`
		TenantToken        string  `json:"tenantToken"`
		Version            int     `json:"version"`
		IsActive           bool    `json:"isActive"`
		BusinessAccountID  int64   `json:"businessAccountId"`
		RoutingFrequency   int     `json:"routingFrequency"`
		PickupStartTime    int64   `json:"pickupStartTime"`
		PickupEndTime      int64   `json:"pickupEndTime"`
		PickupSLA          int64   `json:"pickupSla"`
		CapacityIndex      int     `json:"capacityIndex"`
		BshtTagsIds        []int64 `json:"bshtTagsIds"`
		IsLineHaulRequired bool    `json:"isLineHaulRequired"`
		DropOffZones       []int64 `json:"dropOffZones"`
		PickupZoneID       struct {
			ID              int64  `json:"id"`
			TenantToken     string `json:"tenantToken"`
			Version         int    `json:"version"`
			IsActive        bool   `json:"isActive"`
			ZoneID          string `json:"zoneId"`
			ZoneName        string `json:"zoneName"`
			ZoneStatus      string `json:"zoneStatus"`
			ZoneType        string `json:"zoneType"`
			HasOverLap      bool   `json:"hasOverLap"`
			OperationRegion struct {
				ID          int64  `json:"id"`
				TenantToken string `json:"tenantToken"`
				Version     int    `json:"version"`
				IsActive    bool   `json:"isActive"`
				RegionName  string `json:"regionName"`
			} `json:"operationRegion"`
			OperationCities []struct {
				ID          int64  `json:"id"`
				TenantToken string `json:"tenantToken"`
				Version     int    `json:"version"`
				IsActive    bool   `json:"isActive"`
				CityName    string `json:"cityName"`
			} `json:"operationCities"`
			Provinces       []interface{} `json:"provinces"`
			ProvinceBshtMap struct {
				Num34 int `json:"34"`
			} `json:"provinceBshtMap"`
			OperationRegionID int64 `json:"operationRegionId"`
		} `json:"pickupZoneId"`
	} `json:"businessAccountConfigurations"`
}

type CmsBsht struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID          int    `json:"id"`
		TenantToken string `json:"tenantToken"`
		Version     int    `json:"version"`
		IsActive    bool   `json:"isActive"`
		Label       string `json:"label"`
		IsLinked    bool   `json:"isLinked"`
	} `json:"data"`
}
