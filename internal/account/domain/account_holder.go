package accountdomain

import (
	systemdomain "digital-bank/internal/system/domain"
)

const (
	COMPANY_CLIENT    AccountType = "COMPANY"
	INDIVIDUAL_CLIENT AccountType = "INDIVIDUAL"

	US_CITIZEN         ResidencyStatus = "US_CITIZEN"
	RESIDENT_ALIEN     ResidencyStatus = "RESIDENT_ALIEN"
	NON_RESIDENT_ALIEN ResidencyStatus = "NON_RESIDENT_ALIEN"

	EMPLOYEE      EmploymentStatus = "EMPLOYEE"
	SELF_EMPLOYED EmploymentStatus = "SELF_EMPLOYED"
	RETIRED       EmploymentStatus = "RETIRED"
	UNEMPLOYED    EmploymentStatus = "UNEMPLOYED"
	OTHER         EmploymentStatus = "OTHER"
)

type (
	AccountType      string
	ResidencyStatus  string
	EmploymentStatus string

	AccountHolder interface {
		GetType() AccountType
		GetName() string
		GetIDNumber() string
		GetAddress() Address
		GetPhoneNumber() string
		GetInvestmentProfile() *InvestmentProfile
		SetAccountHolder(holder interface{}) *systemdomain.Error
		ToMap() map[string]interface{}
		SetKYC(kyc KYC)
		SetDocument(document Document, dni string) *systemdomain.Error
		GetDocuments() []Document
	}

	InvestmentProfile struct {
		PrimarySourceOfFunds              string `bson:"primarySourceOfFunds" json:"primarySourceOfFunds"`
		UsdValueOfFiat                    string `bson:"usdValueOfFiat" json:"usdValueOfFiat"`
		UsdValueOfCrypto                  string `bson:"usdValueOfCrypto" json:"usdValueOfCrypto"`
		MonthlyDeposits                   string `bson:"monthlyDeposits" json:"monthlyDeposits"`
		MonthlyCryptoDeposits             string `bson:"monthlyCryptoDeposits" json:"monthlyCryptoDeposits"`
		MonthlyInvestmentDeposit          string `bson:"monthlyInvestmentDeposit" json:"monthlyInvestmentDeposit"`
		MonthlyCryptoInvestmentDeposit    string `bson:"monthlyCryptoInvestmentDeposit" json:"monthlyCryptoInvestmentDeposit"`
		MonthlyWithdrawals                string `bson:"monthlyWithdrawals" json:"monthlyWithdrawals"`
		MonthlyCryptoWithdrawals          string `bson:"monthlyCryptoWithdrawals" json:"monthlyCryptoWithdrawals"`
		MonthlyInvestmentWithdrawal       string `bson:"monthlyInvestmentWithdrawal" json:"monthlyInvestmentWithdrawal"`
		MonthlyCryptoInvestmentWithdrawal string `bson:"monthlyCryptoInvestmentWithdrawal"`
	}

	Address struct {
		StreetOne       string `bson:"streetOne" json:"streetOne"`
		StreetTwo       string `bson:"streetTwo" json:"streetTwo,omitempty"`
		PostalCode      string `bson:"postalCode" json:"postalCode"`
		City            string `bson:"city" json:"city"`
		Region          string `bson:"region" json:"region"`
		Country         string `bson:"country" json:"country"`
		Number          string `bson:"number" json:"number,omitempty"`
		ApartmentNumber string `bson:"apartmentNumber" json:"apartmentNumber,omitempty"`
		IsShipping      bool   `bson:"isShipping" json:"isShipping,omitempty"`
	}

	KYC struct {
		CIPChecks          string            `bson:"CIPChecks" json:"cipChecks"`
		KYCRequiredActions map[string]string `bson:"KYCRequiredActions" json:"kycRequiredActions"`
	}

	KYCProfilePersonal struct {
		FundsSendReceiveJurisdictions string            `bson:"fundsSendReceiveJurisdictions" json:"fundsSendReceiveJurisdictions"`
		EngageInActivities            map[string]string `bson:"engageInActivities" json:"engageInActivities"`
	}

	Individual struct {
		FirstName            string              `bson:"firstName" json:"firstName"`
		DNI                  string              `bson:"dni" json:"dni"`
		MiddleName           *string             `bson:"middleName" json:"middleName"`
		LastName             string              `bson:"lastName" json:"lastName"`
		PhoneNumber          string              `bson:"phoneNumber" json:"phoneNumber"`
		TaxID                string              `bson:"taxID" json:"taxId,omitempty"`
		Email                string              `bson:"email" json:"email"`
		Passport             *string             `bson:"passport" json:"passport,omitempty"`
		DateBirth            *CustomTime         `bson:"dateBirth" json:"dateBirth,omitempty" time_utc:"1"`
		KYC                  *KYC                `bson:"kyc" json:"kyc,omitempty"`
		KYCProfile           *KYCProfilePersonal `bson:"kycProfile" json:"kycProfile"`
		Occupation           string              `bson:"occupation" json:"occupation"`
		EmploymentStatus     EmploymentStatus    `bson:"employmentStatus" json:"employmentStatus"`
		ResidencyStatus      ResidencyStatus     `bson:"residencyStatus" json:"residencyStatus"`
		Documents            []Document          `bson:"documents" json:"documents"`
		Address              *Address            `bson:"address" json:"address,omitempty"`
		InvestmentProfile    *InvestmentProfile  `bson:"investmentProfile" json:"investmentProfile,omitempty"`
		PartnerApplicationId string              `bson:"partnerApplicationId" json:"partnerApplicationId"`
	}
)

// Implementing AccountHolder methods for Individual

func (i *Individual) GetType() AccountType {
	return INDIVIDUAL_CLIENT
}

func (i *Individual) GetName() string {
	if i.MiddleName != nil {
		return i.FirstName + " " + *i.MiddleName + " " + i.LastName
	}
	return i.FirstName + " " + i.LastName
}

func (i *Individual) GetIDNumber() string {
	return i.DNI
}

func (i *Individual) GetAddress() Address {
	return *i.Address
}

func (i *Individual) GetPhoneNumber() string {
	return i.PhoneNumber
}

func (i *Individual) GetInvestmentProfile() *InvestmentProfile {
	return i.InvestmentProfile
}

func (i *Individual) SetAccountHolder(holder interface{}) *systemdomain.Error {
	if individual, ok := holder.(*Individual); ok {
		i.DNI = individual.DNI
		i.FirstName = individual.FirstName
		i.MiddleName = individual.MiddleName
		i.LastName = individual.LastName
		i.TaxID = individual.TaxID
		i.Passport = individual.Passport
		i.DateBirth = individual.DateBirth
		i.KYC = individual.KYC
		i.ResidencyStatus = individual.ResidencyStatus
		i.Documents = individual.Documents
		i.Address = individual.Address
		i.PhoneNumber = individual.PhoneNumber
		i.InvestmentProfile = individual.InvestmentProfile
		i.KYCProfile = individual.KYCProfile
		i.Email = individual.Email
		i.Occupation = individual.Occupation
		i.EmploymentStatus = individual.EmploymentStatus

		return nil
	}
	return systemdomain.NewError(400, "The type of Account Holder is not Individual")
}

func (i *Individual) SetKYC(kyc KYC) {
	i.KYC = &kyc
}

func (i *Individual) SetDocument(document Document, dni string) *systemdomain.Error {
	documentExist := false
	for idx, doc := range i.Documents {
		if doc.GetDocumentType() == document.GetDocumentType() {
			documentExist = true
			i.Documents[idx].UpdateDocument(document)
			break
		}
	}
	if !documentExist {
		i.Documents = append(i.Documents, document)
	}

	return nil
}

func (i *Individual) GetDocuments() []Document {
	return i.Documents
}

func (i *Individual) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"firstName":       i.FirstName,
		"dni":             i.DNI,
		"middleName":      i.MiddleName,
		"lastName":        i.LastName,
		"taxId":           i.TaxID,
		"passport":        i.Passport,
		"dateBirth":       i.DateBirth,
		"kyc":             i.KYC,
		"residencyStatus": i.ResidencyStatus,
		"documents":       i.Documents,
		"address":         i.Address,
	}
}
