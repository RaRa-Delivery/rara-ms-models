package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RaRa-Delivery/rara-ms-models/src/models/cmsdto"
	"github.com/RaRa-Delivery/rara-ms-models/src/utility/lg"
)

func StoreNewCmsContract(secretKey string, accountId int64, token string) (cmsdto.CmsObject, error) {
	//log.Println("secret key: ", secretKey)
	cmsData := cmsdto.CmsObject{}
	cmsObj := cmsdto.CmsBasic{}
	cms := cmsdto.CMS{Token: token}

	res, resE := cms.GetBusinessDetails(secretKey)
	json.Unmarshal([]byte(res), &cmsObj)

	if resE != nil {
		log.Println(lg.Error(resE))
		return cmsData, resE
	}

	bbb, _ := json.Marshal(&cmsObj)

	log.Println(lg.Green("resresres: ", string(bbb)))
	/**Business details**/
	cmsData.BusinessDetails = cmsdto.BusinessDetails{BusinessId: int64(cmsObj.Data.BusinessID), AccountId: int64(cmsObj.Data.ID), AccountName: cmsObj.Data.AccountName, PicName: cmsObj.Data.PicName, PicPhoneNumber: cmsObj.Data.PicPhoneNumber}

	/**Linehaul and special handling boolean check**/
	cmsData.LinehaulRequired = cmsObj.Data.LinehaulRequired
	cmsData.SpecialHandlingRequired = cmsObj.Data.SpecialHandlingRequired

	/**Operation city**/
	//cmsData.OperationCities = cmsObj.Data.OperationCityID
	cityObj := cmsdto.CmsCityObj{}
	cityId := cmsObj.Data.OperationRegionId
	cityRes, cityError := cms.GetOperationalRegionDetails(int64(cityId))

	if cityError != nil {
		log.Println(lg.Error(cityError))
		return cmsData, cityError
	}
	json.Unmarshal([]byte(cityRes), &cityObj)
	cmsData.OperationRegions = cmsdto.CmsRegion{Id: int64(cityObj.Data.ID), RegionName: cityObj.Data.RegionName}
	log.Println(lg.Yellow(cityRes))

	/**Cancellation policy**/
	cmsData.CancellationConditions = cmsdto.CancellationConditions{StatusIndex: 20, IsEligible: true, Type: []string{"ORDER"}, By: []string{"RARA", "CUSTOMER", "END_CUSTOMER"}, Condition: "LESS_THAN"}

	/**Webhook**/
	res, webhookEr := cms.GetWebhookDetails(int64(cmsObj.Data.BusinessID))
	if webhookEr != nil {
		log.Println(lg.Error(webhookEr))
		return cmsData, webhookEr
	}
	webhookObj := cmsdto.WebhookBuilder{}
	json.Unmarshal([]byte(res), &webhookObj)

	webhookActualDataByte, webhookActualDataError := json.Marshal(&webhookObj.Data)
	if webhookActualDataError != nil {
		log.Println(lg.Error(webhookActualDataError))
		return cmsData, webhookActualDataError
	}
	json.Unmarshal(webhookActualDataByte, &cmsData.Webhooks)
	if len(cmsData.Webhooks) > 0 {
		purpose := []string{"OR",
			"OP",
			"BA",
			"PS",
			"PA",
			"PP",
			"PF",
			"SD",
			"AD",
			"NS",
			"RTH",
			"DF",
			"DL",
			"RTW",
			"PWH",
			"DTH",
			"RS",
			"SA",
			"RTS",
			"PWH",
			"DTH",
		}
		cmsData.Webhooks[0].Purpose = append(cmsData.Webhooks[0].Purpose, purpose...)
	}

	/**Chatbot**/

	chatbot := []byte(`{
        "language": "id",
        "notifications": [
            {
                "notificationType": "chatStarter",
                "templateId": "chat_starter_new"
            },
            {
                "notificationType": "csatDelivered",
                "templateId": "csat_delivery_new"
            },
            {
                "notificationType": "orderCreated",
                "templateId": "order_tracking_bahasa_2",
                "params": [
                    {
                        "key": "1",
                        "value": "RECIPIENT_NAME"
                    },
                    {
                        "key": "2",
                        "value": "TRACKING_ID"
                    },
                    {
                        "key": "3",
                        "value": "BUSINESS_NAME"
                    },
                    {
                        "key": "4",
                        "value": "TRACKING_LINK"
                    }
                ]
            },
            {
                "notificationType": "startDelivery",
                "templateId": "start_delivery_2_new",
                "params": [
                    {
                        "key": "1",
                        "value": "RECIPIENT_NAME"
                    },
                    {
                        "key": "2",
                        "value": "TRACKING_ID"
                    },
                    {
                        "key": "3",
                        "value": "BUSINESS_NAME"
                    },
                    {
                        "key": "4",
                        "value": "ETA"
                    },
                    {
                        "key": "5",
                        "value": "OTP"
                    }
                ]
            },
            {
                "notificationType": "failed",
                "templateId": "failed_1_new",
                "params": [
                    {
                        "key": "1",
                        "value": "RECIPIENT_NAME"
                    },
                    {
                        "key": "2",
                        "value": "TRACKING_ID"
                    },
                    {
                        "key": "3",
                        "value": "BUSINESS_NAME"
                    },
                    {
                        "key": "4",
                        "value": "BUSINESS_NAME"
                    }
                ]
            }
        ],
        "postBackUrl": "` + os.Getenv("WHATSAPP_CALLBACK_URL") + `"
    }`)

	cmsChat := cmsdto.CmsChatbotService{}

	log.Println(lg.Green("chatbot: ", string(chatbot)))

	waError := json.Unmarshal(chatbot, &cmsChat)
	log.Println("waError: ", waError)
	if waError != nil {
		log.Println(lg.Error("waError: ", waError))
	}

	cmsData.CmsChatbotService = cmsChat

	/**Order journey**/
	cmsData.OrderJourneyDetails, _ = cmsdto.GetDeliveryStatusList()

	/**service type**/

	if len(cmsObj.Data.ServiceAndPricingConfig) > 0 {

		cmsData.SLAservices = []cmsdto.SlaServiceNew{}

		for _, deliveryService := range cmsObj.Data.ServiceAndPricingConfig {
			slaService := cmsdto.SlaServiceNew{}
			serviceTypeId := deliveryService.DeliveryService.ID
			serviceTypeRes, serviceTypeErr := cms.GetServiceTypeDetails(int64(serviceTypeId))
			if serviceTypeErr != nil {
				log.Println(lg.Error(serviceTypeErr))
				return cmsData, serviceTypeErr
			}
			serviceTypeObj := cmsdto.ServiceTypeObject{}
			json.Unmarshal([]byte(serviceTypeRes), &serviceTypeObj)

			////////////////////////////////////////////////////////////////
			log.Println(lg.Yellow("serviceTypeRes:: ", serviceTypeRes))
			slaService.ID = fmt.Sprint(serviceTypeId)
			//pickup sla
			pickupSla := cmsdto.SLAservice{}
			pickupSla.ID = fmt.Sprint(serviceTypeObj.Data.ID)
			pickupSla.Name = serviceTypeObj.Data.Name
			pickupSla.SLAValue = int32(serviceTypeObj.Data.PickupSLA)
			pickupSla.SLAUnit = "minutes"

			if serviceTypeObj.Data.OperationStartTime > 0 && serviceTypeObj.Data.OperationEndTime > 0 {
				startDateTime := serviceTypeObj.Data.OperationStartTime / 1000
				endDateTime := serviceTypeObj.Data.OperationEndTime / 1000

				startDateTimeUnix := time.Unix(startDateTime, 0)
				st := startDateTimeUnix.UTC().Format("15:04:05")

				endDateTimeUnix := time.Unix(endDateTime, 0)
				ed := endDateTimeUnix.UTC().Format("15:04:05")

				pickupSla.StartTime = st
				pickupSla.EndTime = ed
			}

			pickupSla.CutOffTime = "22:00:00"
			pickupSla.Type = "pickup"
			slaService.Pickup = pickupSla
			//cmsData.SLAservices = append(cmsData.SLAservices, pickupSla)

			//drop sla
			dropSla := cmsdto.SLAservice{}
			dropSla.ID = fmt.Sprint(serviceTypeObj.Data.ID)
			dropSla.Name = serviceTypeObj.Data.Name
			dropSla.SLAValue = int32(serviceTypeObj.Data.SLAType.Duration)
			dropSla.SLAUnit = serviceTypeObj.Data.SLAType.DurationUnitLabel.Label

			if serviceTypeObj.Data.OperationStartTime > 0 && serviceTypeObj.Data.OperationEndTime > 0 {
				startDateTime := serviceTypeObj.Data.OperationStartTime / 1000
				endDateTime := serviceTypeObj.Data.OperationEndTime / 1000

				startDateTimeUnix := time.Unix(startDateTime, 0)
				st := startDateTimeUnix.UTC().Format("15:04:05")

				endDateTimeUnix := time.Unix(endDateTime, 0)
				ed := endDateTimeUnix.UTC().Format("15:04:05")

				dropSla.StartTime = st
				dropSla.EndTime = ed
			}

			dropSla.CutOffTime = "22:00:00"
			dropSla.Type = "dropoff"
			slaService.Dropoff = dropSla
			//cmsData.SLAservices = append(cmsData.SLAservices, dropSla)

			////////////////////////////////////////////////////////////////

			deliveryFeeArray := []cmsdto.DeliveryFeeDto{}

			for _, l := range deliveryService.DeliveryFeeSchemes {
				log.Println("l.Value:", l.Value)
				delVal := strings.ReplaceAll(l.Value, "\u003c", "<")
				delVal = strings.ReplaceAll(delVal, "\u003e", ">")
				delFee := cmsdto.DeliveryFeeDto{}
				delFeeError := json.Unmarshal([]byte(delVal), &delFee)

				if delFeeError != nil {
					return cmsData, delFeeError
				}

				deliveryFeeArray = append(deliveryFeeArray, delFee)

			}

			slaService.DeliveryFee = deliveryFeeArray

			cmsData.SLAservices = append(cmsData.SLAservices, slaService)

		}

	}

	/**Special tags**/
	if cmsData.SpecialHandlingRequired {
		bsht := cmsObj.Data.BusinessAccountProperties.Bsht
		cmsData.Bsht.IsEnabled = bsht.IsEnabled
		//cmsData.Bsht.Tags
		for _, tag := range bsht.Tags {
			//tagIds = append(tagIds, int64(tag.ID))
			//cmsData.Bsht.Tags = append(cmsData.Bsht.Tags, cmsdto.Tag{Id: int64(tag.ID)})

			tagRes, tagErr := cms.GetSpecialHandlingDetails(int64(tag.ID))
			if tagErr != nil {
				log.Println(lg.Error(tagErr))
				return cmsData, tagErr
			}

			bshtdto := cmsdto.CmsBsht{}

			beee := json.Unmarshal([]byte(tagRes), &bshtdto)
			if beee != nil {
				log.Println(beee)
				log.Println(lg.Error(beee))
				return cmsData, beee
			}

			tt := bshtdto.Data
			tagId := tt.ID
			tagLabel := tt.Label
			cmsData.Bsht.Tags = append(cmsData.Bsht.Tags, cmsdto.Tag{Id: int64(tagId), Label: tagLabel})

		}

	}
	/**Pacakge type**/
	packageDataArray := []cmsdto.PackageData{}
	for _, deliveryPackage := range cmsObj.Data.DeliveryPackaging {
		cms := cmsdto.CMS{Token: token}
		packageData := cmsdto.PackageData{}
		wightIndexStr := deliveryPackage.WeightIndex
		weightIndex, weightError := strconv.Atoi(wightIndexStr)
		if weightError != nil {
			log.Println(lg.Error(weightError))
			return cmsData, weightError
		}
		//log.Println("deliveryPackage.PackageTypeScheme.Value:", deliveryPackage.PackageTypeScheme.Value)
		packageData.WeightIndex = float64(weightIndex)

		packageNameObj := cmsdto.BshtDto{}
		packageRes, packageError := cms.GetCommonLabel(int64(deliveryPackage.PackageName.ID))

		if packageError != nil {
			log.Println(lg.Error(packageError))
			return cmsData, packageError
		}

		//	log.Println("packageRes: ", packageRes)

		packageUnmarshalError := json.Unmarshal([]byte(packageRes), &packageNameObj)
		if packageUnmarshalError != nil {
			log.Println(lg.Error(packageUnmarshalError))
			return cmsData, packageUnmarshalError
		}

		if len(packageNameObj.Data) == 0 {
			log.Println(lg.Error("No package name found"))
			return cmsData, errors.New("No package name found")
		}

		packageName := packageNameObj.Data[0].Label
		packageData.Name = packageName
		packageData.Id = int64(deliveryPackage.PackageName.ID)

		str, _ := cms.GetPackageSchemaDetails(int64(deliveryPackage.PackageTypeScheme.ID))
		log.Println("str: ", str)
		log.Println(lg.Yellow(str))
		packageDto := cmsdto.PackDto{}
		pe := json.Unmarshal([]byte(str), &packageDto)
		if pe != nil {
			log.Println(lg.Error(pe))
			return cmsData, pe
		}

		//log.Println("deliveryPackage.PackageTypeScheme.Value: ", deliveryPackage.PackageTypeScheme.Value)
		packageData.Type = packageDto.Data.Name
		packageDefDto := cmsdto.PackageDefinitionDto{}
		packageDefDtoError := json.Unmarshal([]byte(deliveryPackage.PackageTypeScheme.Value), &packageDefDto)
		if packageDefDtoError != nil {
			log.Println(lg.Error(packageDefDtoError))
			return cmsData, packageDefDtoError
		}

		formula := packageDefDto.Formula
		formula = strings.ReplaceAll(formula, "\u003c", "<")
		formula = strings.ReplaceAll(formula, "\u003e", ">")
		packageDefDto.Formula = formula

		packageData.PackageDef = packageDefDto

		log.Println(lg.Yellow("packageData.PackageDef: ", packageData.PackageDef))

		packageDataArray = append(packageDataArray, packageData)

	}

	cmsData.DeliveryPackaging = packageDataArray

	resByte, resError := json.Marshal(&cmsData)
	if resError != nil {
		log.Println(lg.Error(resError))
		return cmsData, resError
	}

	log.Println("NewCmsContract:"+secretKey, " -- ", string(resByte))
	log.Println(lg.Green("NewCmsContract:"+secretKey, " -- ", string(resByte)))

	return cmsData, nil

}
