package accountdomain

import (
	systemdomain "digital-bank/domain/system/domain"
	"time"
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
		StreetOne       string `json:"streetOne"`
		StreetTwo       string `json:"streetTwo,omitempty"`
		PostalCode      string `json:"postalCode"`
		City            string `json:"city"`
		Region          string `json:"region"`
		Country         string `json:"country"`
		Number          string `json:"number,omitempty"`
		ApartmentNumber string `json:"apartmentNumber,omitempty"`
		IsShipping      bool   `json:"isShipping,omitempty"`
	}

	KYC struct {
		CIPChecks          string            `json:"cipChecks"`
		KYCRequiredActions map[string]string `json:"kycRequiredActions"`
	}

	Individual struct {
		FirstName       string          `json:"firstName"`
		DNI             string          `json:"dni"`
		MiddleName      *string         `json:"middleName"`
		LastName        string          `json:"lastName"`
		TaxID           *string         `json:"taxId,omitempty"`
		Passport        *string         `json:"passport,omitempty"`
		DateBirth       *time.Time      `json:"dateBirth,omitempty"`
		KYC             *KYC            `json:"kyc,omitempty"`
		ResidencyStatus ResidencyStatus `json:"residencyStatus"`
		Documents       []Document      `json:"documents"`
		Address         *Address        `json:"address,omitempty"`
	}
)

func (i *Individual) GetType() AccountType {
	return INDIVIDUAL_CLIENT
}

func (i *Individual) GetName() string {
	return i.FirstName + " " + i.LastName
}

func (i *Individual) GetIDNumber() string {
	return i.DNI
}

func (i *Individual) SetAccountHolder(holder interface{}) *systemdomain.Result[string] {
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

	return systemdomain.NewResult("", &systemdomain.ErrorMessage{
		HttpCode: 400,
		Message:  "The type of Account Holder is not Individual",
	})

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
