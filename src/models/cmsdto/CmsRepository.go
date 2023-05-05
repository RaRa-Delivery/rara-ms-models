package cmsdto

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/RaRa-Delivery/rara-ms-models/src/utility"
	"github.com/RaRa-Delivery/rara-ms-models/src/utility/lg"
)

type CMS struct {
	Token string `json:"token"`
}

type IamToken struct {
	Token    string `json:"token,omitempty"`
	ExpireAt int64  `json:"expireAt,omitempty"`
}
type WebhookTypeRequest struct {
	Filters []WebhookTypeFilter `json:"filters"`
	Size    int32               `json:"size"`
}

type WebhookTypeFilter struct {
	FilterKey      string  `json:"key"`
	FilterOperator string  `json:"operator"`
	FilterType     string  `json:"field_type"`
	FilterValue    int64   `json:"value"`
	FilterValues   []int64 `json:"values"`
}

// func SetIamToken() (string, error) {
// 	url := os.Getenv("IAM_UI_URL") + "/api/login"

// 	auth := make(map[string]string)
// 	auth["username"] = "prajwal@rara.delivery"
// 	auth["password"] = "qwerty123"

// 	payloadByte, _ := json.Marshal(&auth)

// 	headers := make(map[string]string)
// 	headers["Content-Type"] = "application/json "
// 	response, err := utility.ApiResponse("POST", url, headers, string(payloadByte))
// 	if err != nil {
// 		log.Println(lg.Error(err))
// 		return "", err
// 	}

// 	tokenPayload := IamToken{}
// 	er := json.Unmarshal([]byte(response), &tokenPayload)
// 	if er != nil {
// 		return "", er
// 	}

// 	SaveByKey("serviceToken", tokenPayload.Token)

// 	return tokenPayload.Token, nil

// }

// func GetSecretKeyByAccountId(id int64) (string, error) {

// 	secretKey, secretSt := FindByKey("secret:" + fmt.Sprint(id))
// 	if !secretSt {

// 		url := os.Getenv("IAM_URL") + "/business/account/secret?businessAccountId=" + fmt.Sprint(id)
// 		token, _ := FindByKey("serviceToken")
// 		headers := make(map[string]string)
// 		headers["Authorization"] = "Bearer " + token
// 		response, err := utility.GetApiResponse(url, headers)
// 		if err != nil {

// 			log.Println(lg.Error(err))
// 			return "", err
// 		}

// 		SaveByKey("secret:"+fmt.Sprint(id), response)

// 		return response, nil

// 	} else {
// 		return secretKey, nil
// 	}

// }

func (c *CMS) GetBusinessDetailsBusinessId(businessId int64) (BusinessData, error) {
	url := os.Getenv("CMS_URL") + "/business/get/" + fmt.Sprint(businessId)
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return BusinessData{}, err
	}

	obj := CmsBusinessDto{}

	eee := json.Unmarshal([]byte(response), &obj)
	if eee != nil {
		log.Println(lg.Error(eee))
		return BusinessData{}, eee
	}
	//log.Println("GetBusinessDetails: ", response)
	return obj.Data, nil

}

func (c *CMS) GetBusinessDetailsByAccountId(accountId int64) (string, error) {
	url := os.Getenv("CMS_URL") + "/business-account/get/" + fmt.Sprint(accountId)
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return "", err
	}
	//log.Println("GetBusinessDetails: ", response)
	return response, nil

}

func (c *CMS) GetBusinessDetailsByOperationRegion(secretKey string, operationRegion int64) (string, error) {
	url := os.Getenv("IAM_URL") + "/business/secret/validate?operationRegionId=" + fmt.Sprint(operationRegion)
	token := c.Token
	headers := make(map[string]string)
	headers["secret"] = secretKey
	headers["Authorization"] = "Bearer " + token
	response, err := utility.ApiResponse("POST", url, headers, "")
	if err != nil {
		log.Println(lg.Error(err))
		return "", err
	}
	//log.Println("GetBusinessDetails: ", response)
	return response, nil

}

func (c *CMS) GetWebhookDetails(businessId int64) (string, error) {
	url := os.Getenv("CMS_URL") + "/business-webhook/search/non-paginated"
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token

	webhookFilter := WebhookTypeFilter{FilterKey: "businessId", FilterOperator: "EQUAL", FilterType: "INTEGER", FilterValue: businessId}
	webhookFilters := []WebhookTypeFilter{}
	webhookFilters = append(webhookFilters, webhookFilter)
	req := WebhookTypeRequest{Filters: webhookFilters, Size: 1}
	reqStr, reqError := json.Marshal(&req)

	if reqError != nil {
		log.Println(lg.Error(reqError))
		return "", reqError
	}

	response, err := utility.ApiResponse("POST", url, headers, string(reqStr))
	if err != nil {
		log.Println(lg.Error(err))
		return "", err
	}

	return response, nil

}

func (c *CMS) GetSpecialHandlingDetails(id int64) (string, error) {
	//defer utility.HandlePanic()
	url := os.Getenv("MDS_URL") + "/bsht-tag/get/" + fmt.Sprint(id)
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token

	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return "", err
	}

	//log.Println("GetSpecialHandlingDetails: ", response)

	return response, nil

}

func (c *CMS) GetOperationalRegionDetails(id int64) (string, error) {

	url := os.Getenv("MDS_URL") + "/operation-region/get/" + fmt.Sprint(id)
	log.Println("operational url: ", url)
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token

	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return "", err
	}
	//log.Println("GetOperationalRegionDetails: ", response)

	return response, nil

}

func (c *CMS) GetServiceTypeDetails(id int64) (string, error) {

	url := os.Getenv("MDS_URL") + "/service/master/get/" + fmt.Sprint(id)
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token

	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return "", err
	}

	return response, nil

}

func (c *CMS) GetPackageSchemaDetails(id int64) (string, error) {

	url := os.Getenv("MDS_URL") + "/package-type-scheme/get/" + fmt.Sprint(id)
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token

	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return "", err
	}

	return response, nil

}

func (c *CMS) GetCommonLabel(id int64) (string, error) {
	url := os.Getenv("MDS_URL") + "/common-labels/search/non-paginated"
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token

	webhookFilter := WebhookTypeFilter{FilterKey: "id", FilterOperator: "EQUAL", FilterType: "INTEGER", FilterValue: id}
	webhookFilters := []WebhookTypeFilter{}
	webhookFilters = append(webhookFilters, webhookFilter)
	req := WebhookTypeRequest{Filters: webhookFilters, Size: 1}
	reqStr, reqError := json.Marshal(&req)

	if reqError != nil {
		log.Println(lg.Error(reqError))
		return "", reqError
	}
	//log.Println("url: ", url)
	//log.Println("Headers: ", headers)
	//log.Println("reqStr: ", string(reqStr))
	response, err := utility.ApiResponse("POST", url, headers, string(reqStr))
	if err != nil {
		log.Println(lg.Error(err))
		return "", err
	}

	return response, nil

}

func (c *CMS) GetBshtPickup(accountId int64, regionId int64) (string, error) {
	url := os.Getenv("MDS_URL") + "/business-pickup-zone-map/search/filterV1"
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token

	req := BshtPickupRequest{}
	req.BusinessAccountIds = []int64{accountId}
	req.IsActive = true
	req.OperationRegionIds = []int64{}
	if regionId > 0 {
		req.OperationRegionIds = append(req.OperationRegionIds, regionId)
	}
	req.BshtTagsIds = []int64{}
	req.PickupLocationName = ""
	req.PageNo = 0
	req.PageSize = 1
	req.SortBy = "pickupLocationName"
	req.SortDir = "asc"
	req.PickupLocationType = "REGISTERED"

	reqStr, reqError := json.Marshal(&req)

	if reqError != nil {
		return "", reqError
	}

	response, err := utility.ApiResponse("POST", url, headers, string(reqStr))
	if err != nil {
		log.Println(err)
		return "", err
	}

	//log.Println(response)

	return response, nil

}

func (c *CMS) GetPickupLocation(locationId string) (CmsPickupDto, error) {
	url := os.Getenv("MDS_URL") + "/business-pickup-zone-map/details/v2?pickupLocationId=" + locationId
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	pickupDto := CmsPickupDto{}
	response, err := utility.GetApiResponse(url, headers)
	//log.Println(response)
	if err != nil {
		log.Println(lg.Error(err))
		return pickupDto, err
	}

	error := json.Unmarshal([]byte(response), &pickupDto)
	if error != nil {
		log.Println(lg.Error(error))
		return pickupDto, error
	}

	return pickupDto, nil
}

func (c *CMS) GetPickupDetailsByDropoffKecamatan(accountId int64, kecamatan string) (PickupByDropoffDto, error) {

	kecamatan = strings.ReplaceAll(kecamatan, " ", "%20")

	url := os.Getenv("MDS_URL") + "/business-pickup-zone-map/reverse-mapping/dropoff-province?dropoffProvince=" + kecamatan + "&accountId=" + fmt.Sprint(accountId)
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	pickupDto := PickupByDropoffDto{}
	response, err := utility.GetApiResponse(url, headers)
	//log.Println(response)
	if err != nil {
		log.Println(lg.Error(err))
		return pickupDto, err
	}

	error := json.Unmarshal([]byte(response), &pickupDto)
	if error != nil {
		log.Println(lg.Error(error))
		return pickupDto, error
	}

	return pickupDto, nil
}

func (c *CMS) GetDropoffZoneByDropoffKecamatan(bshtTag int64, kecamatan string) (DropoffZoneDto, error) {
	kecamatan = strings.ReplaceAll(kecamatan, " ", "%20")
	url := os.Getenv("MDS_URL") + "global-zone/search/district-name?districtName=" + kecamatan + "&zoneType=DROPOFF"
	if bshtTag > 0 {
		url = url + "&bshtTagId=" + fmt.Sprint(bshtTag)
	}
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	pickupDto := DropoffZoneDto{}
	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return pickupDto, err
	}
	//log.Println(response)
	error := json.Unmarshal([]byte(response), &pickupDto)
	if error != nil {
		log.Println(lg.Error(error))
		return pickupDto, error
	}

	return pickupDto, nil
}

func (c *CMS) GetPickupZoneByPickupKecamatan(bshtTag int64, kecamatan string) (DropoffZoneDto, error) {
	kecamatan = strings.ReplaceAll(kecamatan, " ", "%20")
	url := os.Getenv("MDS_URL") + "global-zone/search/district-name?districtName=" + kecamatan + "&zoneType=PICKUP"
	if bshtTag > 0 {
		url = url + "&bshtTagId=" + fmt.Sprint(bshtTag)
	}
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	pickupDto := DropoffZoneDto{}
	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return pickupDto, err
	}
	//(response)
	error := json.Unmarshal([]byte(response), &pickupDto)
	if error != nil {
		log.Println(lg.Error(error))
		return pickupDto, error
	}

	return pickupDto, nil
}

func (c *CMS) GetZoneByKecamatan(bshtTag int64, kecamatan string, zoneType string) (DropoffZoneDto, error) {
	kecamatan = strings.ReplaceAll(kecamatan, " ", "%20")
	url := os.Getenv("MDS_URL") + "/global-zone/search/district-name?districtName=" + kecamatan + "&zoneType=" + zoneType
	if bshtTag > 0 {
		url = url + "&bshtTagId=" + fmt.Sprint(bshtTag)
	}
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	pickupDto := DropoffZoneDto{}
	response, err := utility.GetApiResponse(url, headers)
	//log.Println(response)
	if err != nil {
		log.Println(lg.Error(err))
		return pickupDto, err
	}

	error := json.Unmarshal([]byte(response), &pickupDto)
	if error != nil {
		log.Println(lg.Error(error))
		return pickupDto, error
	}

	return pickupDto, nil
}

func (c *CMS) GetOperationRegionByPostalCode(postalCode string) (OperationRegionByPostalCode, error) {
	opr := OperationRegionByPostalCode{}
	url := os.Getenv("MDS_URL") + "/operation-region/postal-code/" + postalCode
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return opr, err
	}

	oprError := json.Unmarshal([]byte(response), &opr)
	if oprError != nil {
		log.Println(lg.Error(oprError))
		return opr, oprError
	}
	//log.Println("GetOperationRegionByPostalCode: ", response)
	return opr, nil

}

func (c *CMS) GetPackageNameDetails(id int64) (PackageNameData, error) {
	pnDto := PackageNameDto{}
	pn := PackageNameData{}
	url := os.Getenv("MDS_URL") + "/package-type/get/" + fmt.Sprint(id)
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token

	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return pn, err
	}

	oprError := json.Unmarshal([]byte(response), &pnDto)
	if oprError != nil {
		log.Println(lg.Error(oprError))
		return pn, oprError
	}
	//("GetPackageNameDetails: ", response)
	return pnDto.Data, nil

}

func (c *CMS) GetWebhookConfigDetails(accountId int64) (WebhookConfigDto, error) {
	webhookDto := WebhookConfigDto{}
	url := os.Getenv("CMS_URL") + "/notifications/config/web-hook/order-status?accountId=" + fmt.Sprint(accountId)
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token

	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return webhookDto, err
	}

	//log.Println(lg.Mg(response))

	eeee := json.Unmarshal([]byte(response), &webhookDto)
	if eeee != nil {
		log.Println(eeee)
		return webhookDto, eeee
	}

	return webhookDto, nil

}

func (c *CMS) GetNotificationConfigDetails(accountId int64, notificationType string) (NotificationConfigDto, error) {
	webhookDto := NotificationConfigDto{}
	url := os.Getenv("CMS_URL") + "/notifications/config/chat-bot/order-status?accountId=" + fmt.Sprint(accountId) + "&notificationTargetType=" + notificationType
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	//log.Println("chatbot url: ", url)
	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return webhookDto, err
	}

	//log.Println(lg.Mg("chatbot: ", response))

	eeee := json.Unmarshal([]byte(response), &webhookDto)
	if eeee != nil {
		return webhookDto, eeee
	}

	return webhookDto, nil

}

func (c *CMS) GetOperationRegionByKecamatan(kec string) (OperationRegionData, error) {

	kec = strings.ReplaceAll(kec, " ", "%20")

	url := os.Getenv("MDS_URL") + "/operation-region/by-kecametan/" + kec
	token := c.Token
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token

	response, err := utility.GetApiResponse(url, headers)
	if err != nil {
		log.Println(lg.Error(err))
		return OperationRegionData{}, err
	}

	operationRegionDto := OperationRegionDto{}
	eee := json.Unmarshal([]byte(response), &operationRegionDto)
	if err != nil {
		log.Println(lg.Error(err))
		return OperationRegionData{}, eee
	}

	return operationRegionDto.Data, nil

}

// curl -X 'GET' \
//   'https://rara-saas-ms-mds.dev.rara.delivery/business-pickup-zone-map/details/v2?pickupLocationId=130' \
//   -H 'accept: */*' \
//   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwcmFqd2FsQHJhcmEuZGVsaXZlcnkiLCJyb2xlcyI6WyJBRE1JTiJdLCJpYXQiOjE2NzU3OTU0MjcsImV4cCI6MTY3NTg1MzAyN30.Zt8SU4dan00z1_w7ew9T37A3MfSwd-lw6DrK9SljiIY'

// curl -X 'GET' \
//   'https://rara-saas-ms-mds.dev.rara.delivery/sla-type/get/5' \
//   -H 'accept: */*' \
//   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwcmFqd2FsQHJhcmEuZGVsaXZlcnkiLCJyb2xlcyI6WyJBRE1JTiJdLCJpYXQiOjE2NzM2MjkzNjYsImV4cCI6MTY3MzY4Njk2Nn0.gwrAktCAGlktOI1EVcPPSTH1hAJ20PW1HU80Zzm1Zw4'

// BSHT
//   curl -X 'POST' \
//   'https://rara-saas-ms-mds.dev.rara.delivery/common-labels/search/non-paginated' \
//   -H 'accept: */*' \
//   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwcmFqd2FsQHJhcmEuZGVsaXZlcnkiLCJyb2xlcyI6WyJBRE1JTiJdLCJpYXQiOjE2NzM2MjkzNjYsImV4cCI6MTY3MzY4Njk2Nn0.gwrAktCAGlktOI1EVcPPSTH1hAJ20PW1HU80Zzm1Zw4' \
//   -H 'Content-Type: application/json' \
//   -d '{
//   "filters": [
//     {
//       "key": "id",
//       "operator": "IN",
//       "field_type": "INTEGER",

//       "values": [
//         1,2
//       ]
//     }
//   ]
// }'
