package bms

import (
	"github.com/RaRa-Delivery/rara-ms-models/src/models/order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type BatchDriver struct {
// 	ID       string `json:"id" bson:"id"`
// 	Name     string `json:"name" bson:"name"`
// 	MobileNo string `json:"mobileNo" bson:"mobileNo"`
// }

type EnumSizeType string
type EnumWeightType string
type EnumBatchType string
type EnumBatchStatus string
type EnumRemoteServiceEntityType string
type EnumRemoteServiceType string
type EnumOriginType string
type EnumOrderType string
type EnumOrderStatus string
type EnumDeliveryType string

type EnumOrderStatusCode string
type EnumMilestoneStatus string

const EnumDeliveryType_DROP EnumDeliveryType = "DROP"
const EnumDeliveryType_PICK EnumDeliveryType = "PICK"

const EnumOrderType_MARKETPLACE EnumOrderType = "MARKETPLACE"
const EnumOrderType_HUB EnumOrderType = "HUB"

const EnumBatchType_MARKETPLACE EnumBatchType = "MARKETPLACE"
const EnumBatchType_HUB EnumBatchType = "HUB"

const EnumOriginType_MARKETPLACE EnumOriginType = "MARKETPLACE"
const EnumOriginType_HUB EnumOriginType = "HUB"

const EnumOriginType_PICK EnumOriginType = "PIC"
const EnumOriginType_DROP EnumOriginType = "DROP"

const EnumRemoteServiceType_OMS EnumRemoteServiceType = "OMS"
const EnumRemoteServiceType_UMS EnumRemoteServiceType = "UMS"

const EnumRemoteServiceEntityType_ORDER EnumRemoteServiceEntityType = "ORDER"
const EnumRemoteServiceEntityType_HUB EnumRemoteServiceEntityType = "HUB"
const EnumRemoteServiceEntityType_USER EnumRemoteServiceEntityType = "USER"
const EnumRemoteServiceEntityType_CITY EnumRemoteServiceEntityType = "CITY"

const EnumSizeType_RARA_STD EnumSizeType = "RARA_STD"
const EnumSizeType_INCHES EnumSizeType = "INCHES"
const EnumSizeType_CM EnumSizeType = "CM"

const EnumWeightType_RARA_STD EnumWeightType = "RARA_STD"
const EnumWeightType_KG EnumWeightType = "KG"
const EnumWeightType_LBS EnumWeightType = "LBS"

const EnumBatchStatus_NEW EnumBatchStatus = "NEW"
const EnumBatchStatus_START_PICKUP EnumBatchStatus = "START_PICKUP"
const EnumBatchStatus_IN_TRANSIT EnumBatchStatus = "IN_TRANSIT"
const EnumBatchStatus_COMPLETE EnumBatchStatus = "COMPLETE"
const EnumBatchStatus_ACCEPTED EnumBatchStatus = "ACCEPTED"
const EnumBatchStatus_PARTIALLY_DELIVERED EnumBatchStatus = "PARTIALLY_DELIVERED"
const EnumBatchStatus_DELIVERED EnumBatchStatus = "DELIVERED"
const EnumBatchStatus_CANCELLED EnumBatchStatus = "CANCELLED"
const EnumBatchStatus_COMPLETED EnumBatchStatus = "COMPLETED"

const EnumOrderStatus_NEW EnumOrderStatus = "NEW"
const EnumOrderStatus_CANCELLED EnumOrderStatus = "CANCELLED"
const EnumOrderStatus_BATCHED EnumOrderStatus = "BATCHED"
const EnumOrderStatus_DELIVERED EnumOrderStatus = "DELIVERED"
const EnumOrderStatus_IN_TRANSIT EnumOrderStatus = "IN_TRANSIT"

const EnumOrderStatus_VALIDATION_COMPLETED EnumOrderStatus = "Validation completed"
const EnumOrderStatus_DELIVERED_TO_HUB EnumOrderStatus = "Delivered to Hub"
const EnumOrderStatus_BATCH_CREATED EnumOrderStatus = "Batch created"
const EnumOrderStatus_BATCH_ACCEPTED EnumOrderStatus = "Accepted"

const EnumOrderStatusCode_OC EnumOrderStatusCode = "OC"
const EnumOrderStatusCode_VC EnumOrderStatusCode = "VC"
const EnumOrderStatusCode_DTH EnumOrderStatusCode = "DTH"
const EnumOrderStatusCode_BC EnumOrderStatusCode = "BC"
const EnumOrderStatusCode_BA EnumOrderStatusCode = "BA"
const EnumOrderStatusCode_PS EnumOrderStatusCode = "PS"

func ResolveEnumBatchType(str string) EnumBatchStatus {
	switch str {
	case "COMPLETE":
		return EnumBatchStatus_COMPLETE
	case "IN_TRANSIT":
		return EnumBatchStatus_IN_TRANSIT
	case "START_PICKUP":
		return EnumBatchStatus_START_PICKUP
	case "ACCEPTED":
		return EnumBatchStatus_ACCEPTED
	case "PARTIALLY_DELIVERED":
		return EnumBatchStatus_PARTIALLY_DELIVERED
	case "NEW":
		return EnumBatchStatus_NEW
	case "DELIVERED":
		return EnumBatchStatus_DELIVERED
	case "CANCELLED":
		return EnumBatchStatus_CANCELLED
	default:
		return EnumBatchStatus_NEW
	}
}

const EnumOrderStatus_BC EnumOrderStatus = "BC"
const EnumOrderStatus_BA EnumOrderStatus = "BA"
const EnumOrderStatus_PS EnumOrderStatus = "PS"
const EnumOrderStatus_PA EnumOrderStatus = "PA"
const EnumOrderStatus_PP EnumOrderStatus = "PP"
const EnumOrderStatus_PF EnumOrderStatus = "PF"
const EnumOrderStatus_SD EnumOrderStatus = "SD"
const EnumOrderStatus_AD EnumOrderStatus = "AD"
const EnumOrderStatus_NS EnumOrderStatus = "NS"
const EnumOrderStatus_DL EnumOrderStatus = "DL"
const EnumOrderStatus_DF EnumOrderStatus = "DF"
const EnumOrderStatus_OC EnumOrderStatus = "OC"

func ResolveEnumOrderType(str string) EnumOrderStatus {
	switch str {
	case "DELIVERED":
		return EnumOrderStatus_DELIVERED
	case "CANCELLED":
		return EnumOrderStatus_CANCELLED
	case "BATCHED":
		return EnumOrderStatus_BATCHED
	case "IN_TRANSIT":
		return EnumOrderStatus_IN_TRANSIT
	case "BC":
		return EnumOrderStatus_BC
	case "BA":
		return EnumOrderStatus_BA
	case "PS":
		return EnumOrderStatus_PS
	case "PA":
		return EnumOrderStatus_PA
	case "PP":
		return EnumOrderStatus_PP
	case "PF":
		return EnumOrderStatus_PF
	case "SD":
		return EnumOrderStatus_SD
	case "AD":
		return EnumOrderStatus_AD
	case "NS":
		return EnumOrderStatus_NS
	case "DL":
		return EnumOrderStatus_DL
	case "DF":
		return EnumOrderStatus_DF
	case "OC":
		return EnumOrderStatus_OC
	default:
		return EnumOrderStatus_NEW
	}
}

// resolve EnumOriginType
func ResolveEnumOriginType(str string) EnumOriginType {
	switch str {
	case "MARKETPLACE":
		return EnumOriginType_MARKETPLACE
	case "HUB":
		return EnumOriginType_HUB
	case "PICK":
		return EnumOriginType_PICK
	case "DROP":
		return EnumOriginType_DROP
	default:
		return EnumOriginType_HUB
	}
}

// DateTime struct
type DateTime struct {
	Timestamp string `json:"timestamp" bson:"timestamp"`
	Iso       string `json:"iso" bson:"iso"`
}

// RemoteId struct
type RemoteId struct {
	Service EnumRemoteServiceType       `json:"service" bson:"service"`
	Entity  EnumRemoteServiceEntityType `json:"entity" bson:"entity"`
	Id      string                      `json:"id" bson:"id"`
}

type RemoteObj struct {
	Service string `json:"service" bson:"service"`
	Entity  string `json:"entity" bson:"entity"`
	Id      string `json:"id" bson:"id"`
}

// Name Struct
type Name struct {
	First  string `json:"first" bson:"first"`
	Middle string `json:"middle" bson:"middle"`
	Last   string `json:"last" bson:"last"`
}

// RemoteUser struct
type RemoteUser struct {
	UserId RemoteId `json:"userId" bson:"userId"`
	Name   Name     `json:"name" bson:"name"`
	Email  string   `json:"email" bson:"email"`
}

// Size struct
type Size struct {
	Units EnumSizeType `json:"units" bson:"units"`
	Value float64      `json:"value" bson:"value"`
}

// Weight struct
type Weight struct {
	Units EnumWeightType `json:"units" bson:"units"`
	Value float64        `json:"value" bson:"value"`
}

type KecematanInfo struct {
	Name string `json:"name" bson:"name"`
	Id   string `json:"id" bson:"id"`
}

// OriginHub struct
type OriginHub struct {
	OriginType    EnumOriginType `json:"originType" bson:"originType"`
	HubId         RemoteId       `json:"hubId" bson:"hubId"`
	DeliveryType  string         `json:"deliveryType" bson:"deliveryType"`
	CapacityIndex string         `json:"capacityIndex" bson:"capacityIndex"`
	WeightIndex   string         `json:"weightIndex" bson:"weightIndex"`
	Geo           GeoPoint       `json:"geo" bson:"geo"`
}

// OrderDetail struct
type OrderDetail struct {
	Status         EnumOrderStatus      `json:"status" bson:"status"`
	OrderType      EnumOrderType        `json:"orderType" bson:"orderType"`
	OrderId        RemoteId             `json:"orderId" bson:"orderId"`
	CityId         RemoteId             `json:"cityId" bson:"cityId"`
	Kecematan      KecematanInfo        `json:"kecematan" bson:"kecematan"`
	Sla            float64              `json:"sla" bson:"sla"`
	PickupSla      float64              `json:"pickupSla" bson:"pickupSla"`
	Size           Size                 `json:"size" bson:"size"`
	Weight         Weight               `json:"weight" bson:"weight"`
	Geo            GeoPoint             `json:"geo" bson:"geo"`
	Origin         OriginHub            `json:"origin" bson:"origin"`
	BatchInfo      BatchInfo            `json:"batchInfo" bson:"batchInfo"`
	OmsOrderObject order.NewOrderObject `json:"newOrderObject" bson:"newOrderObject"`
	CanPrioritize  bool                 `json:"canPrioritize" bson:"canPrioritize"`
	Otp            string               `json:"otp" bson:"otp"`
}

type OrderDetailObj struct {
	Status         string               `json:"status" bson:"status"`
	OrderType      string               `json:"orderType" bson:"orderType"`
	OrderId        RemoteObj            `json:"orderId" bson:"orderId"`
	CityId         RemoteObj            `json:"cityId" bson:"cityId"`
	Kecematan      KecematanInfo        `json:"kecematan" bson:"kecematan"`
	Sla            float64              `json:"sla" bson:"sla"`
	PickupSla      float64              `json:"pickupSla" bson:"pickupSla"`
	Size           Size                 `json:"size" bson:"size"`
	Weight         Weight               `json:"weight" bson:"weight"`
	Geo            GeoPoint             `json:"geo" bson:"geo"`
	Origin         OriginHub            `json:"origin" bson:"origin"`
	BatchInfo      BatchInfo            `json:"batchInfo" bson:"batchInfo"`
	OmsOrderObject order.NewOrderObject `json:"newOrderObject" bson:"newOrderObject"`
	CanPrioritize  bool                 `json:"canPrioritize" bson:"canPrioritize"`
	Otp            string               `json:"otp" bson:"otp"`
}

// priority_completed
type BatchInfo struct {
	BatchId      string  `json:"batchId" bson:"batchId"`
	OrderInBatch int     `json:"orderInBatch" bson:"orderInBatch"`
	Time         float64 `json:"time" bson:"time"`
	Distance     float64 `json:"distance" bson:"distance"`
	//remaining capacity info for every order is important for marketplace algo ,while hub algo is not dependent on this
	RemainingCapacity float64 `json:"remainingCapacity" bson:"remainingCapacity"`
	PriorityCompleted bool    `json:"priorityCompleted" bson:"priorityCompleted"`
}

type BatchGeo struct {
	LatLngFirst    GeoPoint `json:"latLngFirst" bson:"latLngFirst"`
	LatLngMid      GeoPoint `json:"latLngMid" bson:"latLngMid"`
	LatLngLast     GeoPoint `json:"latLngLast" bson:"latLngLast"`
	LatLngCentroid GeoPoint `json:"latLngCentroid" bson:"latLngCentroid"`
}

type GeoPoint struct {
	Lat     float64 `json:"lat" bson:"lat"`
	Lng     float64 `json:"lng" bson:"lng"`
	PinCode string  `json:"pinCode" bson:"pinCode"`
	GeoHash string  `json:"geoHash" bson:"geoHash"`
}

type BatchGen struct {
	Batch       Batch `json:"batch" bson:"batch"`
	Index       int   `json:"index" bson:"index"`
	Initial     int   `json:"initial" bson:"initial"`
	PickupIndex int   `json:"pickupIndex" bson:"pickupIndex"`
}

type BatchSolution struct {
	Nodes              []BatchNode `json:"nodes" bson:"nodes"`
	BatchId            string      `json:"batchId" bson:"batchId"`
	IsBatchingComplete bool        `json:"isBatchingComplete" bson:"isBatcherComplete"`
	PickupDateTimeUnix int64       `json:"pickupDateTimeUnix" bson:"pickupDateTimeUnix"`
	Duration           float64     `json:"duration" bson:"duration"`
	Distance           float64     `json:"distance" bson:"distance"`
	BatchIdentifier    string      `json:"batchIdentifier" bson:"batchIdentifier"`
}

type BatchNode struct {
	Id                 string     `json:"id" bson:"id"`
	TrackingId         string     `json:"trackingId" bson:"trackingId"`
	IsPickup           bool       `json:"isPickup" bson:"isPickup"`
	IsRouted           bool       `json:"isRouted" bson:"isRouted"`
	IsClosestNode      bool       `json:"isClosestNode" bson:"isClosestNode"`
	Demand             float64    `json:"demand" bson:"demand"`
	LatLng             GeoPoint   `json:"latlng" bson:"latlng"`
	Pickup             PickupInfo `json:"pickup" bson:"pickup"`
	Index              int        `json:"index" bson:"index"`
	BatchId            string     `json:"batchId" bson:"batchId"`
	PickupDateTimeUnix int64      `json:"pickupDateTimeUnix" bson:"pickupDateTimeUnix"`
	Duration           float64    `json:"duration" bson:"duration"`
	Distance           float64    `json:"distance" bson:"distance"`
	Sla                int64      `json:"sla" bson:"sla"`
}

type Batch struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt   DateTime           `json:"createdAt" bson:"createdAt"`
	UpdatedAt   DateTime           `json:"updatedAt" bson:"updatedAt"`
	CreatedBy   RemoteUser         `json:"createdBy" bson:"createdBy"`
	UpdatedBy   RemoteUser         `json:"updatedBy" bson:"updatedBy"`
	Status      EnumBatchStatus    `json:"status" bson:"status"`
	BatchType   EnumBatchType      `json:"batchType" bson:"batchType"`
	OrderCount  int                `json:"orderCount" bson:"orderCount"`
	Bucket      string             `json:"bucket" bson:"bucket"`
	PickupZone  order.Zone         `json:"pickupZone" bson:"pickupZone"`
	DropoffZone order.Zone         `json:"dropoffZone" bson:"dropoffZone"`
	BatchGeo    struct {
		LatLngFirst    GeoPoint `json:"latLngFirst" bson:"latLngFirst"`
		LatLngMid      GeoPoint `json:"latLngMid" bson:"latLngMid"`
		LatLngLast     GeoPoint `json:"latLngLast" bson:"latLngLast"`
		LatLngCentroid GeoPoint `json:"latLngCentroid" bson:"latLngCentroid"`
		KecamatanIn    []string `json:"KecamatanIn" bson:"KecamatanIn"`
		PostalCodeIn   []string `json:"postalCodeIn" bson:"postCodeIn"`
		City           string   `json:"city" bson:"city"`
	} `json:"batchGeo" bson:"batchGeo"`
	BatchCapacity      int           `json:"batchCapacity" bson:"batchCapacity"`
	RemainingCapacity  int           `json:"remainingCapacity" bson:"remainingCapacity"`
	BatchDistance      float64       `json:"batchDistance" bson:"batchDistance"`
	BatchDuration      float64       `json:"batchDuration" bson:"batchDuration"`
	Time               float64       `json:"time" bson:"time"`
	BatchedOrders      []OrderDetail `json:"batchedOrders" bson:"batchedOrders"`
	CancelledOrders    []OrderDetail `json:"cancelledOrders" bson:"cancelledOrders"`
	BatchData          Batchdata     `json:"batchData" bson:"batchData"`
	PickupDateTime     string        `json:"pickupDateTime" bson:"pickupDateTime"`
	PickupDateTimeUnix int64         `json:"pickupDateTimeUnix" bson:"pickupDateTimeUnix"`
	BatchIdentifier    string        `json:"batchIdentifier" bson:"batchIdentifier"`
	Driver             []BatchDriver `json:"driver" bson:"driver"`
	ActualEarning      float64       `json:"actualEarning" bson:"actualEarning"`
	FinalEarning       float64       `json:"finalEarning" bson:"finalEarning"`
	BatchId            string        `json:"batchId" bson:"batchId"`
	LockedBy           string        `json:"lockedBy" bson:"lockedBy"`
	BatchingStatus     int32         `json:"batchingStatus" bson:"batchingStatus"`
	ActiveDriver       BatchDriver   `json:"activeDriver" bson:"activeDriver"`
	AvoidPickups       []PickupInfo  `json:"avoidPickups" bson:"avoidPickups"`
	FirstPickup        string        `json:"firstPickup" bson:"firstPickup"`
}

type PickupInfo struct {
	Name    string   `json:"name" bson:"name"`
	Address string   `json:"address" bson:"address"`
	LatLng  GeoPoint `json:"latlng" bson:"latlng"`
}

type BatchJourney struct {
	Status       string `json:"status" bson:"status"`
	ErrorMessage string `json:"errorMessage" bson:"errorMessage"`
	CreatedAt    string `json:"createdAt"  bson:"createdAt"`
	CreatedBy    string `json:"createdBy" bson:"createdBy"`
	OrderID      string `json:"orderId" bson:"orderId"`
	FromBatchID  string `json:"fromBatchId" bson:"fromBatchId"`
	BatchingType string `json:"batchingType" bson:"batchingType"`
	Driver       struct {
		ID       string `json:"id" bson:"id"`
		Name     string `json:"name" bson:"name"`
		MobileNo string `json:"mobileNo" bson:"mobileNo"`
	} `json:"driver" bson:"driver"`
	ActualEarning int `json:"actualEarning"`
	FinalEarning  int `json:"finalEarning"`
}

type BatchObj struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt   DateTime           `json:"createdAt" bson:"createdAt"`
	UpdatedAt   DateTime           `json:"updatedAt" bson:"updatedAt"`
	CreatedBy   RemoteUser         `json:"createdBy" bson:"createdBy"`
	UpdatedBy   RemoteUser         `json:"updatedBy" bson:"updatedBy"`
	Status      EnumBatchStatus    `json:"status" bson:"status"`
	BatchType   EnumBatchType      `json:"batchType" bson:"batchType"`
	OrderCount  int                `json:"orderCount" bson:"orderCount"`
	Bucket      string             `json:"bucket" bson:"bucket"`
	PickupZone  order.Zone         `json:"pickupZone" bson:"pickupZone"`
	DropoffZone order.Zone         `json:"dropoffZone" bson:"dropoffZone"`
	BatchGeo    struct {
		LatLngFirst    GeoPoint `json:"latLngFirst" bson:"latLngFirst"`
		LatLngMid      GeoPoint `json:"latLngMid" bson:"latLngMid"`
		LatLngLast     GeoPoint `json:"latLngLast" bson:"latLngLast"`
		LatLngCentroid GeoPoint `json:"latLngCentroid" bson:"latLngCentroid"`
		KecamatanIn    []string `json:"KecamatanIn" bson:"KecamatanIn"`
		PostalCodeIn   []string `json:"postalCodeIn" bson:"postCodeIn"`
		City           string   `json:"city" bson:"city"`
	} `json:"batchGeo" bson:"batchGeo"`
	BatchCapacity      int              `json:"batchCapacity" bson:"batchCapacity"`
	RemainingCapacity  int              `json:"remainingCapacity" bson:"remainingCapacity"`
	BatchDistance      float64          `json:"batchDistance" bson:"batchDistance"`
	BatchDuration      float64          `json:"batchDuration" bson:"batchDuration"`
	Time               float64          `json:"time" bson:"time"`
	BatchedOrders      []OrderDetailObj `json:"batchedOrders" bson:"batchedOrders"`
	CancelledOrders    []OrderDetailObj `json:"cancelledOrders" bson:"cancelledOrders"`
	BatchData          Batchdata        `json:"batchData" bson:"batchData"`
	PickupDateTime     string           `json:"pickupDateTime" bson:"pickupDateTime"`
	PickupDateTimeUnix int64            `json:"pickupDateTimeUnix" bson:"pickupDateTimeUnix"`
	BatchIdentifier    string           `json:"batchIdentifier" bson:"batchIdentifier"`
	Driver             []BatchDriver    `json:"driver" bson:"driver"`
	ActualEarning      int              `json:"actualEarning"`
	FinalEarning       int              `json:"finalEarning"`
	LockedBy           string           `json:"lockedBy" bson:"lockedBy"`
	BatchingStatus     int32            `json:"batchingStatus" bson:"batchingStatus"`
	ActiveDriver       BatchDriver      `json:"activeDriver" bson:"activeDriver"`
}

type Batchdata struct {
	BatchTrigeringType      string                  `json:"batchTrigeringType" bson:"batchTrigeringType"`
	ZoneFirstPickup         order.Zone              `json:"zoneFirstPickup" bson:"zoneFirstPickup"`
	Customers               []order.PersonalDetails `json:"customers" bson:"customers"`
	SpecialHandlingCutomers []SpecialHandlingList   `json:"specialHandlingCutomers" bson:"specialHandlingCutomers"`
	ZoneList                []order.Zone            `json:"zoneList" bson:"zoneList"`
	OrderCount              int                     `json:"orderCount" bson:"orderCount"`
	NodesCount              int                     `json:"nodesCount" bson:"nodesCount"`
	BatchJourney            []BatchJourney          `json:"batchJourney" bson:"batchJourney"`
	OperationRegion         order.OperationRegion   `json:"operationRegion" bson:"operationRegion"`
	Bsht                    order.BshtTag           `json:"bsht" bson:"bsht"`
}
type SpecialHandlingList struct {
	Tag          string                `json:"tag" bson:"tag"`
	CustomerInfo order.PersonalDetails `json:"customerInfo" bson:"customerInfo"`
}
