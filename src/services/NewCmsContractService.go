package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RaRa-Delivery/rara-ms-models/src/models/cmsdto"
	"github.com/RaRa-Delivery/rara-ms-models/src/utility/lg"
)

func GetAccountBySecretKey(secretKey string, token string, operationRegion int64) (int64, error) {

	cmsObj := cmsdto.CmsBasic{}
	cms := cmsdto.CMS{Token: token}

	res, resE := cms.GetBusinessDetailsByOperationRegion(secretKey, operationRegion)
	json.Unmarshal([]byte(res), &cmsObj)

	if resE != nil {
		log.Println(lg.Error(resE))
		return 0, resE
	}

	_, e := json.Marshal(&cmsObj)
	if e != nil {
		log.Println(lg.Error(e))
		return 0, e
	}

	return int64(cmsObj.Data.ID), nil

}

func GetAccountByAccountId(token string, accountId int64) (int64, error) {

	cmsObj := cmsdto.CmsBasic{}
	cms := cmsdto.CMS{Token: token}

	res, resE := cms.GetBusinessDetailsByAccountId(accountId)
	json.Unmarshal([]byte(res), &cmsObj)

	if resE != nil {
		log.Println(lg.Error(resE))
		return 0, resE
	}

	_, e := json.Marshal(&cmsObj)
	if e != nil {
		log.Println(lg.Error(e))
		return 0, e
	}

	return int64(cmsObj.Data.ID), nil

}

func StoreNewCmsContract(accountId int64, token string) (cmsdto.CmsObject, error) {
	//log.Println("secret key: ", secretKey)
	cmsData := cmsdto.CmsObject{}
	cmsObj := cmsdto.CmsBasic{}
	cms := cmsdto.CMS{Token: token}

	res, resE := cms.GetBusinessDetailsByAccountId(accountId)
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
	log.Println(lg.Yellow("Webhook"))
	res, webhookEr := cms.GetWebhookDetails(int64(cmsObj.Data.BusinessID))
	if webhookEr != nil {
		log.Println(lg.Error("webhookEr: ", webhookEr))
		return cmsData, webhookEr
	}

	log.Println(lg.Green("Webhook: ", res))
	webhookObj := cmsdto.WebhookBuilder{}
	json.Unmarshal([]byte(res), &webhookObj)

	webhookActualDataByte, webhookActualDataError := json.Marshal(&webhookObj.Data)
	if webhookActualDataError != nil {
		log.Println(lg.Error("webhookActualDataError: ", webhookActualDataError))
		return cmsData, webhookActualDataError
	}
	json.Unmarshal(webhookActualDataByte, &cmsData.Webhooks)
	if len(cmsData.Webhooks) > 0 {
		webhookConfigDto, ee := cms.GetWebhookConfigDetails(int64(cmsObj.Data.ID))
		if ee != nil {
			return cmsData, ee
		}
		purpose := []cmsdto.WebhookStatusMap{
			{Code: "OR", BusinessStatus: "Order rejected"},
			{Code: "OP", BusinessStatus: "Order placed"},
			{Code: "BA", BusinessStatus: "Accepted"},
			{Code: "PS", BusinessStatus: "Start pickup"},
			{Code: "PA", BusinessStatus: "Arrived at pickup"},
			{Code: "PP", BusinessStatus: "Parcel picked"},
			{Code: "PF", BusinessStatus: "Pickup failed"},
			{Code: "SD", BusinessStatus: "Start delivery"},
			{Code: "AD", BusinessStatus: "Arrived at dropoff"},
			{Code: "NS", BusinessStatus: "No show"},
			{Code: "RTH", BusinessStatus: "Returned to hub"},
			{Code: "DF", BusinessStatus: "Delivery failed"},
			{Code: "DL", BusinessStatus: "Delivered"},
			{Code: "RTW", BusinessStatus: "Returned to WH"},
			{Code: "PWH", BusinessStatus: "Picked up from WH"},
			{Code: "DTH", BusinessStatus: "Delivered to Hub"},
			{Code: "RS", BusinessStatus: "Start return"},
			{Code: "SA", BusinessStatus: "Arrived at sender"},
			{Code: "RTS", BusinessStatus: "Returned to sender"},
		}

		purpose1 := []cmsdto.WebhookStatusMap{}
		for _, webhookConfig := range webhookConfigDto.Data {

			code := webhookConfig.OrderStatus.Code
			if code != "" {

				purpose1 = append(purpose1, cmsdto.WebhookStatusMap{Code: code, BusinessStatus: webhookConfig.WebhookStatusName})
			}

		}
		if len(purpose1) > 0 {
			cmsData.Webhooks[0].Purpose = append(cmsData.Webhooks[0].Purpose, purpose1...)
		} else {
			cmsData.Webhooks[0].Purpose = append(cmsData.Webhooks[0].Purpose, purpose...)
		}

	}

	/**Chatbot**/

	notifArr := []cmsdto.CmsChatbotNotification{}

	notif := cmsdto.CmsChatbotNotification{ReceiverType: "RECIPIENT", NotificationType: "chatStarter", TemplateId: "chat_starter_new"}
	notifArr = append(notifArr, notif)

	notif = cmsdto.CmsChatbotNotification{ReceiverType: "RECIPIENT", NotificationType: "DL", TemplateId: "raranow_delivered"}
	notifArr = append(notifArr, notif)

	//RECIPIENT
	recipientNotifiyConfigDto, rnee := cms.GetNotificationConfigDetails(int64(cmsObj.Data.ID), "RECIPIENT")
	if rnee == nil {

		for _, recipientNotify := range recipientNotifiyConfigDto.Data {
			if recipientNotify.Status == "ACTIVE" && recipientNotify.NotificationTargetType == "RECIPIENT" {

				params := []cmsdto.CmsChatbotParams{}
				if recipientNotify.ChatbotNotificationTemplate.TemplateID != "" {
					if recipientNotify.ChatbotNotificationTemplate.Params.Num1 != "" {
						par := cmsdto.CmsChatbotParams{Key: "1", Value: recipientNotify.ChatbotNotificationTemplate.Params.Num1}
						params = append(params, par)
					}
					if recipientNotify.ChatbotNotificationTemplate.Params.Num2 != "" {
						par := cmsdto.CmsChatbotParams{Key: "2", Value: recipientNotify.ChatbotNotificationTemplate.Params.Num2}
						params = append(params, par)
					}
					if recipientNotify.ChatbotNotificationTemplate.Params.Num3 != "" {
						par := cmsdto.CmsChatbotParams{Key: "3", Value: recipientNotify.ChatbotNotificationTemplate.Params.Num3}
						params = append(params, par)
					}
					if recipientNotify.ChatbotNotificationTemplate.Params.Num4 != "" {
						par := cmsdto.CmsChatbotParams{Key: "4", Value: recipientNotify.ChatbotNotificationTemplate.Params.Num4}
						params = append(params, par)
					}
					if recipientNotify.ChatbotNotificationTemplate.Params.Num5 != "" {
						par := cmsdto.CmsChatbotParams{Key: "5", Value: recipientNotify.ChatbotNotificationTemplate.Params.Num5}
						params = append(params, par)
					}
					if recipientNotify.ChatbotNotificationTemplate.Params.Num6 != "" {
						par := cmsdto.CmsChatbotParams{Key: "6", Value: recipientNotify.ChatbotNotificationTemplate.Params.Num6}
						params = append(params, par)
					}
					notif = cmsdto.CmsChatbotNotification{ReceiverType: "RECIPIENT", NotificationType: recipientNotify.OrderStatus.Code, TemplateId: recipientNotify.ChatbotNotificationTemplate.TemplateID, Params: params}
					notifArr = append(notifArr, notif)
				}

			}
		}

	}

	cmsChat := cmsdto.CmsChatbotService{Notifications: notifArr, Language: "id", PostBackUrl: os.Getenv("WHATSAPP_CALLBACK_URL")}

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

			cmsData.OffhoursHandling = deliveryService.OffhoursHandling

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

	cmsData.IsOtpRequiredDuringDropoff = cmsObj.Data.BusinessAccountProperties.IsOtpRequiredDuringDropoff
	cmsData.GeoLocationCheck = cmsObj.Data.BusinessAccountProperties.GeoLocationCheck
	cmsData.GeoLocationDropoffDistanceThreshold = cmsObj.Data.BusinessAccountProperties.GeoLocationDropoffDistanceThreshold

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

		packageNameObj, packageError := cms.GetPackageNameDetails(int64(deliveryPackage.PackageName.ID))

		if packageError != nil {
			log.Println(lg.Error("Issue with package name: ", packageError))
			return cmsData, packageError
		}

		//	log.Println("packageRes: ", packageRes)

		packageName := packageNameObj.PackageName
		packageData.Name = packageName
		packageData.Id = int64(packageNameObj.ID)

		packageDef := packageNameObj.PackageDefinitions[0]

		// str, _ := cms.GetPackageSchemaDetails(int64(deliveryPackage.PackageTypeScheme.ID))
		// log.Println("str: ", str)
		// log.Println(lg.Yellow(str))
		// packageDto := cmsdto.PackDto{}
		// pe := json.Unmarshal([]byte(str), &packageDto)
		// if pe != nil {
		// 	log.Println(lg.Error(pe))
		// 	return cmsData, pe
		// }

		//log.Println("deliveryPackage.PackageTypeScheme.Value: ", deliveryPackage.PackageTypeScheme.Value)
		packageData.Type = packageDef.Metric
		metrics := []cmsdto.Matric{{Ref: packageDef.Reference}}
		vars := []cmsdto.Variable{}

		for _, variable := range packageDef.PackageVariables {
			v := cmsdto.Variable{Ref: variable.Reference, Name: variable.Name, Label: cmsdto.PackageLabel(variable.Label), Type: variable.Type, IsCurrency: variable.IsCurrency, Value: variable.Value}
			vars = append(vars, v)
		}

		packageDefDto := cmsdto.PackageDefinitionDto{Formula: packageDef.Formula, Metric: metrics, Variables: vars}

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

	log.Println("NewCmsContract:"+fmt.Sprint(accountId), " -- ", string(resByte))
	log.Println(lg.Green("NewCmsContract:"+fmt.Sprint(accountId), " -- ", string(resByte)))

	return cmsData, nil

}
