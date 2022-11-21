package models

type BankDetails struct {
	BeneficiaryName     string `json:"beneficiaryName" bson:"beneficiaryName"`
	BeneficiaryEmail    string `json:"beneficiaryEmail" bson:"beneficiaryEmail"`
	AccountNumber       string `json:"accountNumber" bson:"accountNumber"`
	BankName            string `json:"bankName" bson:"bankName"`
	BeneficiaryBankCode string `json:"beneficiaryBankCode" bson:"beneficiaryBankCode"`
	TransferType        string `json:"transferType" bson:"transferType"`
	BankBookPhoto       string `json:"bankBookPhoto" bson:"bankBookPhoto"`
}
