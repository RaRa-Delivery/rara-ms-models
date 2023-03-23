package cmsdto

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/RaRa-Delivery/rara-ms-models/src/utility"
)

type CmsObject struct {
	BusinessDetails         BusinessDetails        `json:"businessDetails,omitempty"`
	LinehaulRequired        bool                   `json:"linehaulRequired,omitempty"`
	SpecialHandlingRequired bool                   `json:"specialHandlingRequired,omitempty"`
	OperationRegions        CmsRegion              `json:"operationRegions,omitempty"`
	Bsht                    BshtDetails            `json:"bshtDetails,omitempty"`
	CancellationConditions  CancellationConditions `json:"cancellationConditions,omitempty"`
	Webhooks                []WebhookData          `json:"webhooks,omitempty"`
	DeliveryPackaging       []PackageData          `json:"deliveryPackaging,omitempty"`
	OrderJourneyDetails     []DeliveryStatusDto    `json:"orderJourneyDetails,omitempty"`
	CmsChatbotService       CmsChatbotService      `json:"chatbot,omitempty"`
	//DeliveryFee             []DeliveryFeeDto       `json:"deliveryFee"`
	SLAservices      []SlaServiceNew `json:"slaService,omitempty"`
	OffhoursHandling string          `json:"offhoursHandling,omitempty"`

	IsOtpRequiredDuringDropoff          bool   `json:"isOtpRequiredDuringDropoff,omitempty"`
	GeoLocationCheck                    string `json:"geoLocationCheck,omitempty"`
	GeoLocationDropoffDistanceThreshold int    `json:"geoLocationDropoffDistanceThreshold,omitempty"`
}

type SlaServiceNew struct {
	ID          string           `json:"id,omitempty" bson:"id"`
	Pickup      SLAservice       `json:"pickup,omitempty" bson:"pickup"`
	Dropoff     SLAservice       `json:"dropoff,omitempty" bson:"dropoff"`
	DeliveryFee []DeliveryFeeDto `json:"deliveryFee,omitempty"`
}

type BusinessDetails struct {
	BusinessId     int64  `json:"businessId,omitempty"`
	AccountId      int64  `json:"accountId,omitempty"`
	AccountName    string `json:"accountName,omitempty"`
	PicName        string `json:"picName,omitempty"`
	PicPhoneNumber string `json:"picPhoneNumber,omitempty"`
}

type CmsRegion struct {
	Id         int64  `json:"id,omitempty"`
	RegionName string `json:"regionName,omitempty"`
}

type BshtDetails struct {
	IsEnabled bool  `json:"isEnabled,omitempty"`
	Tags      []Tag `json:"tags,omitempty"`
}

type Tag struct {
	Id    int64  `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
}

type CancellationConditions struct {
	StatusIndex float32  `json:"statusIndex,omitempty"`
	IsEligible  bool     `json:"isEligible,omitempty"`
	Type        []string `json:"type,omitempty"`
	By          []string `json:"by,omitempty"`
	Condition   string   `json:"condition,omitempty"`
}

type CmsChatbotService struct {
	Language      string                   `json:"language,omitempty" bson:"language"`
	Notifications []CmsChatbotNotification `json:"notifications,omitempty" bson:"notifications"`
	PostBackUrl   string                   `json:"postBackUrl,omitempty" bson:"postBackUrl"`
}

type CmsChatbotNotification struct {
	NotificationType string             `json:"notificationType,omitempty" bson:"notificationType"`
	TemplateId       string             `json:"templateId,omitempty" bson:"templateId"`
	Params           []CmsChatbotParams `json:"params,omitempty" bson:"params"`
	ReceiverType     string             `json:"receiverType,omitempty" bson:"receiverType"`
}

type CmsChatbotParams struct {
	Key   string `json:"key,omitempty" bson:"key"`
	Value string `json:"value,omitempty" bson:"value"`
}

func (cmsResponse CmsObject) EvalueateExpressionByDeliveryFee(input float64, serviceFormulaObj DeliveryFeeDto) (bool, float64) {
	defer utility.HandlePanic()

	log.Println("serviceFormulaObj.Formula: ", serviceFormulaObj.Formula)

	serviceFormulaObj.Formula = `if (^weight < #a) then {#b} else {#b + ((^weight - #a) * #c))}`

	formulaExpression := strings.ReplaceAll(serviceFormulaObj.Formula, "^", "")
	formulaExpression = strings.ReplaceAll(formulaExpression, "#", "")
	metricObj := serviceFormulaObj.Metric[0]
	inputValue := strings.ReplaceAll(metricObj.Ref, "^", "")
	log.Println("input", input)
	if input > metricObj.MaxValue || input < metricObj.MinValue {
		return false, 0.0
	}

	log.Println("Delivery fee formula: ", formulaExpression)

	expression, error := govaluate.NewEvaluableExpression(formulaExpression)
	if error != nil {
		log.Println(error)
		return false, 0.0
	}
	parameters := make(map[string]interface{}, 8)
	parameters[inputValue] = input

	for i := 0; i < len(serviceFormulaObj.Variables); i++ {
		varObj := serviceFormulaObj.Variables[i]
		variableData := strings.ReplaceAll(varObj.Ref, "#", "")
		parameters[variableData] = varObj.Value

	}

	eb, _ := json.Marshal(&parameters)
	log.Println("eb: ", eb)

	result, err := expression.Evaluate(parameters)
	if err != nil {
		log.Println(err)
		return false, 0.0
	}

	i, ok := result.(float64)
	log.Println("Checking delivery fee")
	log.Println(i)
	if !ok {
		return false, 0.0
	}

	return true, i
}

func (cmsResponse *CmsObject) EvalueateExpressionByDeliveryFeeByType(input float64, shippingType string, unit string) (bool, float64, string, string) {
	//input is the distance
	serviceFormulaObj := cmsResponse.SLAservices[0].DeliveryFee[0]

	log.Println(serviceFormulaObj)
	formulaExpression := strings.ReplaceAll(serviceFormulaObj.Formula, "^", "")
	formulaExpression = strings.ReplaceAll(formulaExpression, "#", "")
	metricObj := serviceFormulaObj.Metric[0]
	inputValue := strings.ReplaceAll(metricObj.Ref, "^", "")
	log.Println("input", input)
	if input > metricObj.MaxValue || input < metricObj.MinValue {
		errorMessage := "<SHIPPING_TYPE> should be in the range of <L>-<M> in <UNIT>"
		errorMessage = strings.Replace(errorMessage, "<SHIPPING_TYPE>", shippingType, 1)
		errorMessage = strings.Replace(errorMessage, "<L>", fmt.Sprint(metricObj.MinValue), 1)
		errorMessage = strings.Replace(errorMessage, "<M>", fmt.Sprint(metricObj.MaxValue), 1)
		errorMessage = strings.Replace(errorMessage, "<UNIT>", unit, 1)

		return false, 0.0, errorMessage, "SH40001"
	}
	log.Println("formulaExpression: ", formulaExpression)
	expression, error := govaluate.NewEvaluableExpression(formulaExpression)
	if error != nil {
		log.Println(error)
		return false, 0.0, "Unknown error in shipping fee calculation", "SH40002"
	}
	parameters := make(map[string]interface{}, 8)
	parameters[inputValue] = input

	for i := 0; i < len(serviceFormulaObj.Variables); i++ {
		varObj := serviceFormulaObj.Variables[i]
		variableData := strings.ReplaceAll(varObj.Ref, "#", "")
		parameters[variableData] = varObj.Value

	}

	log.Println("parameters: ", parameters)

	result, err := expression.Evaluate(parameters)
	if err != nil {
		log.Println(err)
		return false, 0.0, err.Error(), "SH40003"
	}

	i, ok := result.(float64)
	log.Println("Checking delivery fee")
	log.Println(i)
	if !ok {
		return false, 0.0, "", "SH40004"
	}

	return true, i, "", ""
}

func (cmsObj *CmsObject) SetDeliveryFee(val float64, shippingType string, unit string) (float64, string, bool, string) {
	//orderWeight := d.PackageDetails.BillableWeight

	//temp fix to be reviewd ??
	//orderWeight := 0.0
	// for _, ele := range ItemWeights{
	// 	orderWeight += ele
	// }

	st, data, message, code := cmsObj.EvalueateExpressionByDeliveryFeeByType(val, shippingType, unit)
	// if data > 0 {
	// 	d.OrderDetails.DeliveryFee = data
	// }
	return data, message, st, code
}
