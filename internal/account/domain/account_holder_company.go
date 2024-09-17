package clientdomain

import (
	"digital-bank/internal/system/domain"
	"fmt"
	"time"
)

const (
	PRIVATE    CompanyType = "C_CORP_PRIVATE"
	PUBLIC     CompanyType = "C_CORP_PUBLIC"
	HNWI       CompanyType = "HNWI"
	LLC        CompanyType = "LLC"
	LLP        CompanyType = "LLP"
	LP         CompanyType = "LP"
	S_CORP     CompanyType = "S_CORP"
	SOLE_PROP  CompanyType = "SOLE_PROP"
	TRUST      CompanyType = "TRUST"
	NON_PROFIT CompanyType = "NON_PROFIT"
)

type (
	CompanyType string

	CompanyAccountHolder interface {
		AccountHolder
		AddPartner(partner Individual)
		EditPartner(dni string, updatedPartner Individual)
		setDocument(document Document, dni string)
		UpdatePartnerKYC(dni string, kyc *KYC)
	}

	CompanyQuestionnaire struct {
		PurposeAccount                         string `json:"purposeAccount"`
		SourceAssetsAndIncome                  string `json:"sourceAssetsAndIncome"`
		IntendedUseAccount                     string `json:"intendedUseAccount"`
		AnticipatedTypesAssets                 string `json:"anticipatedTypesAssets"`
		AnticipatedMonthlyCashVolume           string `json:"anticipatedMonthlyCashVolume"`
		AnticipatedTradingPatterns             string `json:"anticipatedTradingPatterns"`
		AnticipatedMonthlyTransactionsIncoming string `json:"anticipatedMonthlyTransactionsIncoming"`
		AnticipatedMonthlyTransactionsOutgoing string `json:"anticipatedMonthlyTransactionsOutgoing"`
		NatureBusinessCompany                  string `json:"natureBusinessCompany"`
	}

	Company struct {
		Name                      string                `json:"name"`
		PrimaryBusiness           string                `json:"primaryBusiness"`
		DescriptionBusinessNature string                `json:"descriptionBusinessNature"`
		RegisterNumber            string                `json:"registerNumber"`
		NAICS                     string                `json:"naics"`
		NAICSDescription          string                `json:"naicsDescription"`
		CompanyType               CompanyType           `json:"companyType"`
		EstablishedDate           time.Time             `json:"establishedDate"`
		WebSite                   string                `json:"webSite"`
		RegisteredAddress         Address               `json:"registeredAddress"`
		PhysicalAddress           Address               `json:"physicalAddress,omitempty"`
		PhoneCountry              string                `json:"phoneCountry"`
		PhoneNumber               string                `json:"phoneNumber"`
		Documents                 []Document            `json:"documents"`
		KYC                       *KYC                  `json:"kyc,omitempty"`
		Questionnaire             *CompanyQuestionnaire `json:"questionnaire"`
		Partners                  []Individual          `json:"partners"`
	}
)

func (c *Company) GetType() AccountType {
	return COMPANY_CLIENT
}

func (c *Company) GetName() string {
	return c.Name
}

func (c *Company) SetAccountHolder(holder interface{}) *domain.Result[string] {
	if company, ok := holder.(*Company); ok {
		c.Name = company.Name
		c.PrimaryBusiness = company.PrimaryBusiness
		c.DescriptionBusinessNature = company.DescriptionBusinessNature
		c.RegisterNumber = company.RegisterNumber
		c.NAICS = company.NAICS
		c.NAICSDescription = company.NAICSDescription
		c.CompanyType = company.CompanyType
		c.EstablishedDate = company.EstablishedDate
		c.WebSite = company.WebSite
		c.RegisteredAddress = company.RegisteredAddress
		c.PhysicalAddress = company.PhysicalAddress
		c.PhoneCountry = company.PhoneCountry
		c.PhoneNumber = company.PhoneNumber
		c.Documents = company.Documents
		c.KYC = company.KYC
		c.Questionnaire = company.Questionnaire
		c.Partners = company.Partners

		return nil
	}

	return domain.NewResult("", &domain.ErrorMessage{
		HttpCode: 400,
		Message:  "The type of Account Holder is not Company",
	})
}

func (c *Company) SetKYC(kyc KYC) {
	c.KYC = &kyc
}

func (c *Company) AddPartner(partner Individual) {
	c.Partners = append(c.Partners, partner)
}

func (c *Company) EditPartner(dni string, updatedPartner Individual) {
	for index, partner := range c.Partners {
		if partner.DNI == dni {
			c.Partners[index].FirstName = updatedPartner.FirstName
			c.Partners[index].MiddleName = updatedPartner.MiddleName
			c.Partners[index].LastName = updatedPartner.LastName
			c.Partners[index].TaxID = updatedPartner.TaxID
			c.Partners[index].Passport = updatedPartner.Passport
			c.Partners[index].DateBirth = updatedPartner.DateBirth
			c.Partners[index].ResidencyStatus = updatedPartner.ResidencyStatus
			c.Partners[index].Address = updatedPartner.Address

			break
		}
	}

}

func (c *Company) UpdatePartnerKYC(dni string, kyc *KYC) {
	for index, partner := range c.Partners {
		if partner.DNI == dni {
			c.Partners[index].KYC = kyc
			break
		}
	}

}

func (c *Company) GetIDNumber() string {
	return c.RegisterNumber
}

func (c *Company) setDocument(document Document, dni string) *domain.Result[string] {
	// assign or updated document to company
	if c.GetIDNumber() == dni {
		documentExist := false
		for _, doc := range c.Documents {
			if doc.GetDocumentType() == document.GetDocumentType() {
				documentExist = true
				doc.UpdateDocument(document)
				break
			}
		}

		if !documentExist {
			c.Documents = append(c.Documents, document)
		}
		return nil
	}

	if len(c.Partners) == 0 {
		fmt.Println("Company partners not found")
		return domain.NewResult("", &domain.ErrorMessage{
			HttpCode: 404,
			Message:  "Company partners not found",
		})
	}

	// assign or updated document to partner
	for _, partner := range c.Partners {
		if partner.DNI != dni {
			continue
		}

		documentExist := false
		for _, doc := range partner.Documents {
			if doc.GetDocumentType() == document.GetDocumentType() {
				documentExist = true
				doc.UpdateDocument(document)
				break
			}
		}

		if !documentExist {
			partner.Documents = append(partner.Documents, document)
		}
	}

	return nil

}

func (c *Company) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":                      c.Name,
		"primaryBusiness":           c.PrimaryBusiness,
		"descriptionBusinessNature": c.DescriptionBusinessNature,
		"registerNumber":            c.RegisterNumber,
		"naics":                     c.NAICS,
		"naicsDescription":          c.NAICSDescription,
		"companyType":               c.CompanyType,
		"establishedDate":           c.EstablishedDate,
		"webSite":                   c.WebSite,
		"registeredAddress":         c.RegisteredAddress,
		"physicalAddress":           c.PhysicalAddress,
		"phoneCountry":              c.PhoneCountry,
		"phoneNumber":               c.PhoneNumber,
		"documents":                 c.Documents,
		"kyc":                       c.KYC,
		"questionnaire":             c.Questionnaire,
		"partners":                  c.Partners,
	}
}
