package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type Bank struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	BankName string             `json:"bank_name" bson:"bank_name"`
	Status   string             `json:"status" bson:"status"`
	BankCode string             `json:"bank_code" bson:"bank_code"`
	Version  int                `json:"version" bson:"version"`
}

type BankValidator struct {
	BankAccountNumber string `json:"bank_account_number" bson:"bank_account_number"`
	BankCode          string `json:"bank_code" bson:"bank_code"`
	GivenName         string `json:"given_name" bson:"given_name"`
	SurName           string `json:"surname" bson:"surname"`
	Mobile            string `json:"mobile" bson:"mobile"`
	ReferenceId       string `json:"referenceId" bson:"referenceId"`
}

type BankValidatorResponse struct {
	Status            string    `json:"status"`
	BankAccountNumber string    `json:"bank_account_number"`
	BankCode          string    `json:"bank_code"`
	Created           time.Time `json:"created"`
	Updated           time.Time `json:"updated"`
	ID                string    `json:"id"`
	Result            struct {
		IsFound            bool   `json:"is_found"`
		IsVirtualAccount   bool   `json:"is_virtual_account"`
		NeedReview         bool   `json:"need_review"`
		NameMatchingResult string `json:"name_matching_result"`
	} `json:"result"`
}
