package models

type BankDetails struct {
	BeneficiaryName          string `json:"beneficiaryName" bson:"beneficiaryName"`
	BeneficiaryEmail         string `json:"beneficiaryEmail" bson:"beneficiaryEmail"`
	AccountNumber            string `json:"accountNumber" bson:"accountNumber"`
	BankName                 string `json:"bankName" bson:"bankName"`
	BeneficiaryBankCode      string `json:"beneficiaryBankCode" bson:"beneficiaryBankCode"`
	TransferType             string `json:"transferType" bson:"transferType"`
	BeneficiaryCustResidence string `json:"beneficiaryCustResidence" bson:"beneficiaryCustResidence"`
	BeneficiaryCustType      string `json:"beneficiaryCustType" bson:"beneficiaryCustType"`
	BankBookPhoto            string `json:"bankBookPhoto" bson:"bankBookPhoto"`
	RabbitMobile             string `json:"rabbitMobile" bson:"rabbitMobile"`
	CreatedAt                int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt                int64  `json:"updatedAt" bson:"updatedAt"`
	Status                   string `json:"status" bson:"status"`
}
