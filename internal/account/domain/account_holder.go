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
)

type (
	AccountType     string
	ResidencyStatus string

	AccountHolder interface {
		GetType() AccountType
		GetName() string
		GetIDNumber() string
		SetAccountHolder(holder interface{})
		ToMap() map[string]interface{}
		SetKYC(kyc KYC)
		setDocument(document Document, dni string)
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

	Individual struct {
		FirstName       string          `bson:"firstName" json:"firstName"`
		DNI             string          `bson:"dni" json:"dni"`
		MiddleName      *string         `bson:"middleName" json:"middleName"`
		LastName        string          `bson:"lastName" json:"lastName"`
		TaxID           *string         `bson:"taxID" json:"taxId,omitempty"`
		Passport        *string         `bson:"passport" json:"passport,omitempty"`
		DateBirth       *CustomTime     `bson:"dateBirth" json:"dateBirth,omitempty" time_utc:"1"`
		KYC             *KYC            `bson:"kyc" json:"kyc,omitempty"`
		ResidencyStatus ResidencyStatus `bson:"residencyStatus" json:"residencyStatus"`
		Documents       []Document      `bson:"documents" json:"documents"`
		Address         *Address        `bson:"address" json:"address,omitempty"`
	}
)

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

		return nil
	}

	return systemdomain.NewError(400, "The type of Account Holder is not Individual")

}

func (i *Individual) SetKYC(kyc KYC) {
	i.KYC = &kyc
}

func (i *Individual) setDocument(document Document, dni string) {
	documentExist := false
	for _, doc := range i.Documents {
		if doc.GetDocumentType() == document.GetDocumentType() {
			documentExist = true
			doc.UpdateDocument(document)
			break
		}
	}

	if !documentExist {
		i.Documents = append(i.Documents, document)
	}
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
