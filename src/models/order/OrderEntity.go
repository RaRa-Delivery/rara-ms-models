package order

import (
	"github.com/RaRa-Delivery/rara-ms-models/src/models/cmsdto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewOrderObject struct {
	Id                  primitive.ObjectID  `json:"_id" bson:"_id"`
	UploadId            string              `json:"uploadId" bson:"uploadId"`
	TenantToken         string              `json:"tenantToken" bson:"tenantToken"`
	BusinessDetails     BusinessDetails     `json:"businessDetails" bson:"businessDetails"`
	OrderDetails        NewOrderDetails     `json:"orderDetails" bson:"orderDetails"`
	PickupDetails       PickupDetails       `json:"pickupDetails" bson:"pickupDetails"`
	DropoffDetails      DropOffDetails      `json:"dropoffDetails" bson:"dropoffDetails"`
	SourcePickupDetails PickupDetails       `json:"sourcePickupDetails" bson:"sourcePickupDetails"`
	HubDetails          Hub                 `json:"hubDetails" bson:"hubDetails"`
	ZoneDetails         ZoneDetails         `json:"zoneDetails" bson:"zoneDetails"`
	PaymentDetails      PaymentDetails      `json:"paymentDetails" bson:"paymentDetails"`
	BatchDetails        BatchDetails        `json:"batchDetails" bson:"batchDetails"`
	Notifications       []Notification      `json:"notification" bson:"notification"`
	Webhooks            []Webhook           `json:"webhooks" bson:"webhooks"`
	StatusDetails       StatusDetails       `json:"statusDetails" bson:"statusDetails"`
	Status              string              `json:"status" bson:"status"`
	Source              string              `json:"source" bson:"source"`
	OrderDistanceMatrix OrderDistanceMatrix `json:"orderDistanceMatrix" bson:"orderDistanceMatrix"`
	Driver              Driver              `json:"driver" bson:"driver"`
	Attempt             int                 `json:"attempt" bson:"attempt"`
	CodAmountDetails    CodAmountDetails    `json:"codAmountDetails" bson:"codAmountDetails"`
	OperationRegion     OperationRegion     `json:"operationRegion" bson:"operationRegion"`
	LockedBy            string              `json:"lockedBy" bson:"lockedBy"`
	Version             int                 `json:"version" bson:"version"`
}

type ApiPayload struct {
	UploadId        string          `json:"uploadId" bson:"uploadId"`
	TenantToken     string          `json:"tenantToken" bson:"tenantToken"`
	BusinessDetails BusinessDetails `json:"businessDetails" bson:"businessDetails"`
	OrderDetails    OrderDetails    `json:"orderDetails" bson:"orderDetails"`
	PickupDetails   PickupDetails   `json:"pickupDetails" bson:"pickupDetails"`
	DropOffDetails  DropOffDetails  `json:"dropoffDetails" bson:"dropoffDetails"`
	// PackageDetails  PackageDetails  `json:"packageDetails" bson:"packageDetails"`
	PaymentDetails PaymentDetails `json:"paymentDetails" bson:"paymentDetails"`
	Pieces         []Piece        `json:"pieces" bson:"pieces"`
	Source         string         `json:"source" bson:"source"`
	Webhooks       []Webhook      `json:"webhooks" bson:"webhooks"`
}

type Piece struct {
	Id              primitive.ObjectID `json:"Id,omitempty" bson:"Id,omitempty"`
	PieceId         string             `json:"pieceId" bson:"pieceId"`
	Weight          float64            `json:"weight" bson:"weight"`
	PieceSize       string             `json:"pieceSize" bson:"pieceSize"`
	Dimensions      OrderDimensions    `json:"dimensions" bson:"dimensions"`
	VolWeight       float64            `json:"volWeight" bson:"volWeight"`
	BillableWeight  float64            `json:"billableWeight" bson:"billableWeight"`
	WeightIndex     float64            `json:"weightIndex" bson:"weightIndex"`
	Price           float64            `json:"price" bson:"price"`
	SpecialHandling string             `json:"specialHandling" bson:"specialHandling"`
	StatusDetails   StatusDetails      `json:"statusDetails" bson:"statusDetails"`
}

type CodAmountDetails struct {
	InitialAmount float64 `json:"initialAmount" bson:"initialAmount"`
	Amount        float64 `json:"amount" bson:"amount"`
}

type OperationRegion struct {
	Id   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type BusinessDetails struct {
	BusinessId    int64  `json:"businessId" bson:"businessId"`
	AccountNumber int64  `json:"accountNumber" bson:"accountNumber"`
	ServiceType   string `json:"serviceType" bson:"serviceType"`
	ServiceID     string `json:"serviceId" bson:"serviceId"`
	ServiceTypeId int64  `json:"serviceTypeId" bson:"serviceTypeId"`
}

type NewOrderDetails struct {
	PieceId       string             `json:"pieceId" bson:"pieceId"`
	Status        string             `json:"status" bson:"status"`
	TrackingId    string             `json:"trackingId" bson:"trackingId"`
	WeightDetails OrderWeightDetails `json:"weightDetails" bson:"weightDetails"`
	ParcelSize    string             `json:"parcelSize" bson:"parcelSize"`
	Dimensions    OrderDimensions    `json:"dimensions" bson:"dimensions"`
	WeightIndex   float64            `json:"weightIndex" bson:"weightIndex"`
	PackageId     int64              `json:"packageId" bson:"packageId"`

	OrderDeliveryDetails NewOrderDeliveryDetails `json:"orderDeliveryDetails" bson:"orderDeliveryDetails"`
}

type OrderDetails struct {
	PieceId              string               `json:"pieceId" bson:"pieceId"`
	Status               string               `json:"status" bson:"status"`
	TrackingId           string               `json:"trackingId" bson:"trackingId"`
	WeightDetails        OrderWeightDetails   `json:"weightDetails" bson:"weightDetails"`
	PackageId            int64                `json:"packageId" bson:"packageId"`
	ParcelSize           string               `json:"parcelSize" bson:"parcelSize"`
	Dimensions           OrderDimensions      `json:"dimensions" bson:"dimensions"`
	WeightIndex          float64              `json:"weightIndex" bson:"weightIndex"`
	SpecialHandling      string               `json:"specialHandling" bson:"specialHandling"`
	OrderDeliveryDetails OrderDeliveryDetails `json:"orderDeliveryDetails" bson:"orderDeliveryDetails"`
}

type OrderWeightDetails struct {
	Weight         float64 `json:"weight" bson:"weight"`
	VolWeight      float64 `json:"volWeight" bson:"volWeight"`
	BillableWeight float64 `json:"billableWeight" bson:"billableWeight"`
}

type OrderDimensions struct {
	Length float64 `json:"length" bson:"length"`
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
	Unit   string  `json:"unit" bson:"unit"`
}

type NewOrderDeliveryDetails struct {
	OrderDate       string      `json:"orderDate" bson:"orderDate"`
	PickupDate      string      `json:"pickupDate" bson:"pickupDate"`
	OrderDistance   float64     `json:"orderDistance" bson:"orderDistance"`
	DeliveryFee     float64     `json:"deliveryFee" bson:"deliveryFee"`
	SpecialHandling string      `json:"specialHandling" bson:"specialHandling"`
	Linehaul        ApiLinehaul `json:"linehaul" bson:"linehaul"`
	Sla             Sla         `json:"sla" bson:"sla"`
	Bsht            BshtTag     `json:"bsht" bson:"bsht"`
}

type OrderDeliveryDetails struct {
	Linehaul      bool    `json:"linehaul" bson:"linehaul"`
	OrderDate     string  `json:"orderDate" bson:"orderDate"`
	PickupDate    string  `json:"pickupDate" bson:"pickupDate"`
	OrderDistance float64 `json:"orderDistance" bson:"orderDistance"`
}

type BshtTag struct {
	Id    int64  `json:"id" bson:"id"`
	Label string `json:"label" bson:"label"`
}

type ApiLinehaul struct {
	IsEnabled        int          `json:"isEnabled" bson:"isEnabled"`
	Condition        ApiCondition `json:"condition" bson:"condition"`
	IsLinehaulRouted bool         `json:"isLinehaulRouted" bson:"isLinehaulRouted"`
}

type ApiCondition struct {
	Field string  `json:"field" bson:"field"`
	Value float64 `json:"value" bson:"value"`
}

type Sla struct {
	Pickup  int64 `json:"pickup" bson:"pickup"`
	Dropoff int64 `json:"dropoff" bson:"dropoff"`
}

type PickupDetails struct {
	PickupInchargeDetails PersonalDetails `json:"pickupIncharge" bson:"pickupIncharge"`
	LocationDetails       LocationDetails `json:"locationDetails" bson:"locationDetails"`
	ExpectedPuDateAndTime string          `json:"expectedPuDateTime" bson:"expectedPuDateTime"`
	Slot                  string          `json:"slot" bson:"slot"`
	Note                  string          `json:"note" bson:"note"`
}

type DropOffDetails struct {
	RecipientDetails PersonalDetails `json:"recipientDetails" bson:"recipientDetails"`
	LocationDetails  LocationDetails `json:"locationDetails" bson:"locationDetails"`
	ReqDlTime        string          `json:"reqDlTime" bson:"reqDlTime"`
	DlNote           string          `json:"dlNote" bson:"dlNote"`
	Otp              string          `json:"otp" bson:"otp"`
}

type PersonalDetails struct {
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	PhoneNo string `json:"phone" bson:"phone"`
}

type LocationDetails struct {
	LocationName string   `json:"locationName" bson:"locationName"`
	Address      string   `json:"address" bson:"address"`
	Address1     string   `json:"address1" bson:"address1"`
	Address2     string   `json:"address2" bson:"address2"`
	SubDistrict  string   `json:"subDistrict" bson:"subDistrict"`
	District     string   `json:"district" bson:"district"`
	City         string   `json:"city" bson:"city"`
	PostalCode   string   `json:"postalCode" bson:"postalCode"`
	GeoPoint     GeoPoint `json:"geoPoint" bson:"geoPoint"`
	Type         string   `json:"type" bson:"type"`
	Province     string   `json:"province" bson:"province"`
	Id           int64    `json:"id" bson:"id"`
}

type GeoPoint struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

// type PackageDetails struct {
// 	Size           string     `json:"packageSize" bson:"packageSize"`
// 	Description    string     `json:"packageDescription" bson:"packageDescription"`
// 	Value          string     `json:"packageValue" bson:"packageValue"`
// 	NoOfPieces     int        `json:"numberofPiece" bson:"numberofPiece"`
// 	Dimensions     Dimensions `json:"dimensions" bson:"dimensions"`
// 	VolWeight      float64    `json:"volWeight" bson:"volWeight"`
// 	BillableWeight float64    `json:"billableWeight" bson:"billableWeight"`
// 	WeightIndex    float64    `json:"weightIndex" bson:"weightIndex"`
// }

type PaymentDetails struct {
	Method string  `json:"paymentMethod" bson:"paymentMethod"`
	Price  float64 `json:"price" bson:"price"`
}

type Hub struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Address    string             `json:"address" bson:"address"`
	PostalCode string             `json:"postalCode" bson:"postalCode"`
	Kelurahan  string             `json:"kelurahan" bson:"kelurahan"`
	Kecamatan  string             `json:"kecamatan" bson:"kecamatan"`
	Status     string             `json:"status" bson:"status"`
	City       City               `json:"city" bson:"city"`
	Pic        Pic                `json:"pic" bson:"pic"`
	Geo        GeoDetails         `json:"geo" bson:"geo"`
}

type City struct {
	Id   int32  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Slug string `json:"slug" bson:"slug"`
}

type Pic struct {
	Name          string `json:"name" bson:"name"`
	ContactNumber string `json:"contactNumber" bson:"contactNumber"`
}

type GeoDetails struct {
	LatLng LatLngData `json:"latlng" bson:"latlng"`
	H3Geo  H3GeoData  `json:"h3Geo" bson:"h3Geo"`
}

type LatLngData struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

type H3GeoData struct {
	Res6 H3Res `json:"res6" bson:"res6"`
	Res5 H3Res `json:"res5" bson:"res5"`
	Res4 H3Res `json:"res4" bson:"res4"`
}

type H3Res struct {
	Resolution int32  `json:"resolution" bson:"resolution"`
	HashValue  string `json:"hashValue" bson:"hashValue"`
}

type ZoneDetails struct {
	Pickup  Zone `json:"pickup" bson:"pickup"`
	Dropoff Zone `json:"dropoff" bson:"dropoff"`
}

type Zone struct {
	Id                  int64  `json:"id" bson:"id"`
	Type                string `json:"type" bson:"type"`
	Name                string `json:"name" bson:"name"`
	Status              string `json:"status" bson:"status"`
	Kecamatan           string `json:"kecamatan" bson:"kecamatan"`
	SearchableKecamatan string `json:"searchableKecamatan" bson:"searchableKecamatan"`
}

type BatchJourneyData struct {
	BatchHistory              []string `json:"batchHistory" bson:"batchHistory"`
	CreationTime              string   `json:"creationTime" bson:"creationTime"`
	Id                        string   `json:"id" bson:"id"`
	AssigningTime             string   `json:"assigningTime" bson:"assigningTime"`
	AcceptedTime              string   `json:"acceptedTime" bson:"acceptedTime"`
	Driver                    Driver   `json:"driver" bson:"driver"`
	Cancel                    Cancel   `json:"cancel" bson:"cancel"`
	FacilityArrivalTime       string   `json:"facilityArrivalTime" bson:"facilityArrivalTime"`
	OnHoldTime                string   `json:"onHoldTime" bson:"onHoldTime"`
	DepartedFromFacilityTime  string   `json:"departedFromFacilityTime" bson:"departedFromFacilityTime"`
	OrderReturnedTime         string   `json:"orderReturnedTime" bson:"orderReturnedTime"`
	BatchProcessingIdentifier string   `json:"batchProcessingIdentifier" bson:"batchProcessingIdentifier"`
	CreationTimeUnix          int64    `json:"creationTimeUnix" bson:"creationTimeUnix"`
	PickupDateTime            string   `json:"pickupDateTime" bson:"pickupDateTime"`
	PickupDateTimeUnix        int64    `json:"pickupDateTimeUnix" bson:"pickupDateTimeUnix"`
}
type BatchDetails struct {
	Current BatchJourneyData   `json:"current" bson:"current"`
	Journey []BatchJourneyData `json:"journey" bson:"journey"`
}

type Driver struct {
	Id           string `json:"id" bson:"id"`
	Name         string `json:"name" bson:"name"`
	MobileNumber string `json:"mobileNumber" bson:"mobileNumber"`
}

type Cancel struct {
	CancelledTime string `json:"cancelledTime" bson:"cancelledTime"`
	Reason        string `json:"reason" bson:"reason"`
}

type Notification struct {
	Message string `json:"message" bson:"message"`
	Status  string `json:"status" bson:"status"`
}

type Webhook struct {
	Purpose       []cmsdto.WebhookStatusMap `json:"purpose" bson:"purpose"`
	RequestMethod string                    `json:"requestMethod" bson:"requestMethod"`
	Url           string                    `json:"url" bson:"url"`
	Headers       []Headers                 `json:"headers" bson:"headers"`
	Payload       string                    `json:"payload" bson:"payload"`
}

type Headers struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

type StatusDetails struct {
	Current OrderStatus   `json:"current" bson:"current"`
	Journey []OrderStatus `json:"journey" bson:"journey"`
	Next    []OrderStatus `json:"next" bson:"next"`
}

type OrderStatus struct {
	Code                  string               `json:"code" bson:"code"`
	Status                string               `json:"status" bson:"status"`
	Languages             Languages            `json:"externalStatus" bson:"externalStatus"`
	CreatedAt             int64                `json:"createdAt" bson:"createdAt"`
	Attempt               int                  `json:"attempt" bson:"attempt"`
	DriverLatLng          GeoPoint             `json:"driver_latlng" bson:"driver_latlng"`
	Images                []string             `json:"images" bson:"images"`
	FailedTime            string               `json:"failedTime" bson:"failedTime"`
	FailedReason          string               `json:"failedReason" bson:"failedReason"`
	Comment               string               `json:"comment" bson:"comment"`
	RecipientName         string               `json:"recipientName" bson:"recipientName"`
	RecipientRelationship string               `json:"recipientRelationship" bson:"recipientRelationship"`
	WebhookResponse       WebhookResponse      `json:"webhookResponse,omitempty" bson:"webhookResponse"`
	NotificationResponse  NotificationResponse `json:"notificationResponse,omitempty" bson:"notificationResponse"`
}

type WebhookResponse struct {
	WebhookSent     string `json:"webhookSent" bson:"webhookSent"`
	CreatedAt       int64  `json:"createdAt" bson:"createdAt"`
	WebhookRequest  string `json:"webhookRequest" bson:"webhookRequest"`
	WebhookResponse string `json:"webhookResponse" bson:"webhookResponse"`
}

type NotificationResponse struct {
	NotificationSent     string `json:"notificationSent" bson:"notificationSent"`
	CreatedAt            int64  `json:"createdAt" bson:"createdAt"`
	Template             string `json:"template" bson:"template"`
	NotificationResponse string `json:"notificationResponse" bson:"notificationResponse"`
}

type Languages struct {
	Id string `json:"id" bson:"id"`
	En string `json:"en" bson:"en"`
}

type OrderDistanceMatrix struct {
	Distance struct {
		ValueInMeter int     `json:"valueInMeter"`
		ValueInKm    float64 `json:"valueInKm"`
	} `json:"distance"`
	Displacement struct {
		ValueInMeter int     `json:"valueInMeter"`
		ValueInKm    float64 `json:"valueInKm"`
	} `json:"displacement"`
	Duration struct {
		ValueInSeconds int     `json:"valueInSeconds"`
		ValueInMinutes float64 `json:"valueInMinutes"`
	} `json:"duration"`
}

type WebhookProducer struct {
	TemplateType string         `json:"templateType" bson:"templateType"`
	W            Webhook        `json:"w" bson:"w"`
	WebhookOrder NewOrderObject `json:"webhookOrder" bson:"webhookOrder"`
	Status       string         `json:"status" bson:"status"`
	Message      string         `json:"message" bson:"message"`
	Purpose      string         `json:"purpose" bson:"purpose"`
	Description  string         `json:"description" bson:"description"`
	RequestId    string         `json:"requestId" bson:"requestId"`
}
