package cmsdto

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/RaRa-Delivery/rara-ms-models/src/utility"
)

type CmsObject struct {
	BusinessDetails         BusinessDetails        `json:"businessDetails"`
	LinehaulRequired        bool                   `json:"linehaulRequired"`
	SpecialHandlingRequired bool                   `json:"specialHandlingRequired"`
	OperationRegions        CmsRegion              `json:"operationRegions"`
	Bsht                    BshtDetails            `json:"bshtDetails"`
	CancellationConditions  CancellationConditions `json:"cancellationConditions"`
	Webhooks                []WebhookData          `json:"webhooks"`
	DeliveryPackaging       []PackageData          `json:"deliveryPackaging"`
	OrderJourneyDetails     []DeliveryStatusDto    `json:"orderJourneyDetails"`
	CmsChatbotService       CmsChatbotService      `json:"chatbot"`
	//DeliveryFee             []DeliveryFeeDto       `json:"deliveryFee"`
	SLAservices []SlaServiceNew `json:"slaService"`
}

type SlaServiceNew struct {
	ID          string           `json:"id" bson:"id"`
	Pickup      SLAservice       `json:"pickup" bson:"pickup"`
	Dropoff     SLAservice       `json:"dropoff" bson:"dropoff"`
	DeliveryFee []DeliveryFeeDto `json:"deliveryFee"`
}

type BusinessDetails struct {
	BusinessId     int64  `json:"businessId"`
	AccountId      int64  `json:"accountId"`
	AccountName    string `json:"accountName"`
	PicName        string `json:"picName"`
	PicPhoneNumber string `json:"picPhoneNumber"`
}

type CmsRegion struct {
	Id         int64  `json:"id"`
	RegionName string `json:"regionName"`
}

type BshtDetails struct {
	IsEnabled bool  `json:"isEnabled"`
	Tags      []Tag `json:"tags"`
}

type Tag struct {
	Id    int64  `json:"id"`
	Label string `json:"label"`
}

type CancellationConditions struct {
	StatusIndex int      `json:"statusIndex"`
	IsEligible  bool     `json:"isEligible"`
	Type        []string `json:"type"`
	By          []string `json:"by"`
	Condition   string   `json:"condition"`
}

type CmsChatbotService struct {
	Language      string                   `json:"language" bson:"language"`
	Notifications []CmsChatbotNotification `json:"notifications" bson:"notifications"`
	PostBackUrl   string                   `json:"postBackUrl" bson:"postBackUrl"`
}

type CmsChatbotNotification struct {
	NotificationType string             `json:"notificationType" bson:"notificationType"`
	TemplateId       string             `json:"templateId" bson:"templateId"`
	Params           []CmsChatbotParams `json:"params" bson:"params"`
}

type CmsChatbotParams struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
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
