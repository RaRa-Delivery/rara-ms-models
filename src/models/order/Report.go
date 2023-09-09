package order

import "go.mongodb.org/mongo-driver/bson/primitive"

type Report struct {
	Id                  primitive.ObjectID `json:"_id" bson:"_id"`
	TrackingID          string             `json:"trackingId" bson:"trackingId,omitempty"`
	OrderDate           string             `json:"orderDate" bson:"orderDate,omitempty"`
	OrderDateUnix       int64              `json:"orderDateUnix" bson:"orderDateUnix,omitempty"`
	OperationRegion     string             `json:"operationRegion" bson:"operationRegion,omitempty"`
	OperationRegionID   int                `json:"operationRegionId" bson:"operationRegionId,omitempty"`
	BusinessID          int                `json:"buisnessId" bson:"buisnessId,omitempty"`
	AccountNumber       int                `json:"accountNumber" bson:"accountNumber,omitempty"`
	BusinessName        string             `json:"buisnessName" bson:"buisnessName,omitempty"`
	AccountName         string             `json:"accountName" bson:"accountName,omitempty"`
	OrderType           string             `json:"orderType" bson:"orderType,omitempty"`
	BSHT                string             `json:"bsht" bson:"bsht,omitempty"`
	BSHTID              int                `json:"bshtId" bson:"bshtId,omitempty"`
	ServiceTypeId       int64              `json:"serviceTypeId" bson:"serviceTypeId,omitempty"`
	ServiceType         string             `json:"serviceType" bson:"serviceType,omitempty"`
	ServiceID           int                `json:"serviceId" bson:"serviceId,omitempty"`
	OrderDistance       float64            `json:"orderDistance" bson:"orderDistance,omitempty"`
	Status              string             `json:"status" bson:"status,omitempty"`
	StatusCode          string             `json:"statusCode" bson:"statusCode,omitempty"`
	Weight              float64            `json:"weight" bson:"weight,omitempty"`
	ParcelSize          string             `json:"parcelSize" bson:"parcelSize,omitempty"`
	PackageID           int                `json:"packageId" bson:"packageId,omitempty"`
	CODAmount           float64            `json:"codAmount" bson:"codAmount,omitempty"`
	CODDepositStatus    string             `json:"codDepositStatus" bson:"codDepositStatus,omitempty"`
	BatchID             string             `json:"batchId" bson:"batchId,omitempty"`
	DriverID            string             `json:"driverId" bson:"driverId,omitempty"`
	DriverMobile        string             `json:"driverMobile" bson:"driverMobile,omitempty"`
	DriverName          string             `json:"driverName" bson:"driverName,omitempty"`
	SenderName          string             `json:"senderName" bson:"senderName,omitempty"`
	SenderPhone         string             `json:"senderPhone" bson:"senderPhone,omitempty"`
	PickupLocationName  string             `json:"pickupLocationName" bson:"pickupLocationName,omitempty"`
	PickupAddressLine1  string             `json:"pickupAddressLine1" bson:"pickupAddressLine1,omitempty"`
	PickupAddressLine2  string             `json:"pickupAddressLine2" bson:"pickupAddressLine2,omitempty"`
	PickupKecamatan     string             `json:"pickupKecamatan" bson:"pickupKecamatan,omitempty"`
	PickupCity          string             `json:"pickupCity" bson:"pickupCity,omitempty"`
	PickupZipCode       string             `json:"pickupZipCode" bson:"pickupZipCode,omitempty"`
	PickupZoneType      string             `json:"pickupZoneType" bson:"pickupZoneType,omitempty"`
	PickupZoneID        int                `json:"pickupZoneId" bson:"pickupZoneId,omitempty"`
	PickupZoneName      string             `json:"pickupZoneName" bson:"pickupZoneName,omitempty"`
	PickupLatitude      float64            `json:"pickupLatitude" bson:"pickupLatitude,omitempty"`
	PickupLongitude     float64            `json:"pickupLongitude" bson:"pickupLongitude,omitempty"`
	RecipientName       string             `json:"recipientName" bson:"recipientName,omitempty"`
	RecipientPhone      string             `json:"recipientPhone" bson:"recipientPhone,omitempty"`
	DropoffAddressLine1 string             `json:"dropoffAddressLine1" bson:"dropoffAddressLine1,omitempty"`
	DropoffAddressLine2 string             `json:"dropoffAddressLine2" bson:"dropoffAddressLine2,omitempty"`
	DropoffKecamatan    string             `json:"dropoffKecamatan,omitempty" bson:"dropoffKecamatan,omitempty"`
	DropoffCity         string             `json:"dropoffCity,omitempty" bson:"dropoffCity,omitempty"`
	DropoffProvince     string             `json:"dropoffProvince,omitempty" bson:"dropoffProvince,omitempty"`
	DropoffZipCode      string             `json:"dropoffZipCode,omitempty" bson:"dropoffZipCode,omitempty"`
	DropoffZoneId       int                `json:"dropoffZoneId,omitempty" bson:"dropoffZoneId,omitempty"`
	DropoffZoneName     string             `json:"dropoffZoneName,omitempty" bson:"dropoffZoneName,omitempty"`
	DropoffLatitude     float64            `json:"dropoffLatitude,omitempty" bson:"dropoffLatitude,omitempty"`
	DropoffLongitude    float64            `json:"dropoffLongitude,omitempty" bson:"dropoffLongitude,omitempty"`
	BatchingTime1       string             `json:"batchingTime1,omitempty" bson:"batchingTime1,omitempty"`
	BatchingTime2       string             `json:"batchingTime2,omitempty" bson:"batchingTime2,omitempty"`
	Pwh                 string             `json:"pwh,omitempty" bson:"pwh,omitempty"`
	Dth                 string             `json:"dth,omitempty" bson:"dth,omitempty"`
	DriverAcceptance    string             `json:"driverAcceptance,omitempty" bson:"driverAcceptance,omitempty"`
	StartPickup         string             `json:"startPickup,omitempty" bson:"startPickup,omitempty"`
	ArrivedAtPickup     string             `json:"arrivedAtPickup,omitempty" bson:"arrivedAtPickup,omitempty"`
	PickupCompleted     string             `json:"pickupCompleted,omitempty" bson:"pickupCompleted,omitempty"`
	PickupFailed        string             `json:"pickupFailed,omitempty" bson:"pickupFailed,omitempty"`
	StartDelivery       string             `json:"startDelivery,omitempty" bson:"startDelivery,omitempty"`
	ArrivedAtDelivery   string             `json:"arrivedAtDelivery,omitempty" bson:"arrivedAtDelivery,omitempty"`
	Delivered           string             `json:"delivered,omitempty" bson:"delivered,omitempty"`
	DeliveryFailed      string             `json:"deliveryFailed,omitempty" bson:"deliveryFailed,omitempty"`
	StartReturn         string             `json:"startReturn,omitempty" bson:"startReturn,omitempty"`
	ArrivedAtSender     string             `json:"arrivedAtSender,omitempty" bson:"arrivedAtSender,omitempty"`
	ReturnedToSender    string             `json:"returnedToSender,omitempty" bson:"returnedToSender,omitempty"`
	PfReason            string             `json:"pfReason,omitempty" bson:"pfReason,omitempty"`
	DfReason            string             `json:"dfReason,omitempty" bson:"dfReason,omitempty"`
	OrderCancel         string             `json:"orderCancel,omitempty" bson:"orderCancel,omitempty"`
	OrderCancelReason   string             `json:"orderCancelReason,omitempty" bson:"orderCancelReason,omitempty"`
	PopImages           string             `json:"popImages,omitempty" bson:"popImages,omitempty"`
	PopfImages          string             `json:"popfImages,omitempty" bson:"popfImages,omitempty"`
	PodImages           string             `json:"podImages,omitempty" bson:"podImages,omitempty"`
	PodfImages          string             `json:"podfImages,omitempty" bson:"podfImages,omitempty"`
}

func (f *Report) GetStructFieldBsonFieldNameMap() map[string]string {
	m := make(map[string]string)
	m["trackingId"] = "TrackingID"
	m["orderDate"] = "OrderDate"
	m["orderDateUnix"] = "OrderDateUnix"
	m["operationRegion"] = "OperationRegion"
	m["operationRegionId"] = "OperationRegionID"
	m["buisnessId"] = "BusinessID"
	m["accountNumber"] = "AccountNumber"
	m["buisnessName"] = "BusinessName"
	m["accountName"] = "AccountName"
	m["orderType"] = "OrderType"
	m["bsht"] = "BSHT"
	m["bshtId"] = "BSHTID"
	m["serviceTypeId"] = "ServiceTypeId"
	m["serviceType"] = "ServiceType"
	m["serviceId"] = "ServiceID"
	m["orderDistance"] = "OrderDistance"
	m["status"] = "Status"
	m["statusCode"] = "StatusCode"
	m["weight"] = "Weight"
	m["parcelSize"] = "ParcelSize"
	m["packageId"] = "PackageID"
	m["codAmount"] = "CODAmount"
	m["codDepositStatus"] = "CODDepositStatus"
	m["batchId"] = "BatchID"
	m["driverId"] = "DriverID"
	m["driverMobile"] = "DriverMobile"
	m["driverName"] = "DriverName"
	m["senderName"] = "SenderName"
	m["senderPhone"] = "SenderPhone"
	m["pickupLocationName"] = "PickupLocationName"
	m["pickupAddressLine1"] = "PickupAddressLine1"
	m["pickupAddressLine2"] = "PickupAddressLine2"
	m["pickupKecamatan"] = "PickupKecamatan"
	m["pickupCity"] = "PickupCity"
	m["pickupZipCode"] = "PickupZipCode"
	m["pickupZoneType"] = "PickupZoneType"
	m["pickupZoneId"] = "PickupZoneID"
	m["pickupZoneName"] = "PickupZoneName"
	m["pickupLatitude"] = "PickupLatitude"
	m["pickupLongitude"] = "PickupLongitude"
	m["recipientName"] = "RecipientName"
	m["recipientPhone"] = "RecipientPhone"
	m["dropoffAddressLine1"] = "DropoffAddressLine1"
	m["dropoffAddressLine2"] = "DropoffAddressLine2"
	m["dropoffKecamatan"] = "DropoffKecamatan"
	m["dropoffCity"] = "DropoffCity"
	m["dropoffProvince"] = "DropoffProvince"
	m["dropoffZipCode"] = "DropoffZipCode"
	m["dropoffZoneId"] = "DropoffZoneId"
	m["dropoffZoneName"] = "DropoffZoneName"
	m["dropoffLatitude"] = "DropoffLatitude"
	m["dropoffLongitude"] = "DropoffLongitude"
	m["batchingTime1"] = "BatchingTime1"
	m["batchingTime2"] = "BatchingTime2"
	m["pwh"] = "Pwh"
	m["dth"] = "Dth"
	m["driverAcceptance"] = "DriverAcceptance"
	m["startPickup"] = "StartPickup"
	m["arrivedAtPickup"] = "ArrivedAtPickup"
	m["pickupCompleted"] = "PickupCompleted"
	m["pickupFailed"] = "PickupFailed"
	m["startDelivery"] = "StartDelivery"
	m["arrivedAtDelivery"] = "ArrivedAtDelivery"
	m["delivered"] = "Delivered"
	m["deliveryFailed"] = "DeliveryFailed"
	m["startReturn"] = "StartReturn"
	m["arrivedAtSender"] = "ArrivedAtSender"
	m["returnedToSender"] = "ReturnedToSender"
	m["pfReason"] = "PfReason"
	m["dfReason"] = "DfReason"
	m["orderCancel"] = "OrderCancel"
	m["orderCancelReason"] = "OrderCancelReason"
	m["popImages"] = "PopImages"
	m["popfImages"] = "PopfImages"
	m["podImages"] = "PodImages"
	m["podfImages"] = "PodfImages"

	return m
}
