package accountdomain

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

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	str = str[1 : len(str)-1] // Remove quotes
	ct.Time, err = time.Parse(ctLayout, str)
	return
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Format(ctLayout) + `"`), nil
}

type (
	CompanyType string

	CompanyAccountHolder interface {
		AccountHolder
		AddPartner(partner Individual)
		EditPartner(dni string, updatedPartner Individual)
		setDocument(document Document, dni string) *systemdomain.Error
		UpdatePartnerKYC(dni string, kyc *KYC)
		GetAddress() Address
		GetRegisteredAddress() Address
		GetCompanyData() Company
		GetPartners() []Individual
		GetKYCProfile() *KYCProfile
	}

	KYCProfile struct {
		DescriptionBusinessNature     string   `json:"descriptionBusinessNature" bson:"descriptionBusinessNature"`
		BusinessJurisdictions         []string `json:"businessJurisdictions" bson:"businessJurisdictions"`
		FundsSendReceiveJurisdictions []string `json:"fundsSendReceiveJurisdictions" bson:"fundsSendReceiveJurisdictions"`
		EngageInActivities            []string `json:"engageInActivities" bson:"engageInActivities"`
		RegulatedStatus               string   `json:"regulatedStatus" bson:"regulatedStatus"`
		PrimaryBusiness               string   `json:"primaryBusiness" bson:"primaryBusiness"`
	}

	Company struct {
		Name              string             `bson:"name" json:"name" binding:"required"`
		RegisterNumber    string             `bson:"registerNumber" json:"registerNumber"`
		NAICS             string             `bson:"naics" json:"naics" binding:"required"`
		NAICSDescription  string             `bson:"naicsDscription" json:"naicsDescription"`
		CompanyType       CompanyType        `bson:"companyType" json:"companyType"`
		EstablishedDate   CustomTime         `bson:"establishedDate" json:"establishedDate"`
		WebSite           string             `bson:"webSite" json:"webSite" binding:"required" `
		RegisteredAddress Address            `bson:"registeredAddress "json:"registeredAddress" `
		PhysicalAddress   Address            `bson:"physicalAddress" json:"physicalAddress,omitempty" `
		PhoneCountry      string             `bson:"phoneCountry" json:"phoneCountry"`
		PhoneNumber       string             `bson:"phoneNumber" json:"phoneNumber"`
		Email             string             `bson:"email" json:"email"`
		Documents         []Document         `bson:"documents" json:"documents"`
		KYC               *KYC               `bson:"kyc" json:"kyc,omitempty"`
		InvestmentProfile *InvestmentProfile `bson:"investmentProfile" json:"investmentProfile"`
		KYCProfile        *KYCProfile        `bson:"kycProfile" json:"kycProfile"`
		Partners          *[]Individual      `bson:"partners" json:"partners"`
	}
)

func (c *Company) GetAddress() Address {
	panic("not implement for company")
}

func (c *Company) GetType() AccountType {
	return COMPANY_CLIENT
}

func (c *Company) GetName() string {
	return c.Name
}

func (c *Company) SetAccountHolder(holder interface{}) *systemdomain.Error {
	if company, ok := holder.(*Company); ok {
		c.Name = company.Name
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
		c.InvestmentProfile = company.InvestmentProfile
		c.Partners = company.Partners
		c.Email = company.Email

		return nil
	}

	return systemdomain.NewError(400, "The type of Account Holder is not Company")
}

func (c *Company) SetKYC(kyc KYC) {
	c.KYC = &kyc
}

func (c *Company) AddPartner(partner Individual) {
	*c.Partners = append(*c.Partners, partner)

}

func (c *Company) EditPartner(dni string, updatedPartner Individual) {
	for index, partner := range *c.Partners {
		if partner.DNI == dni {
			(*c.Partners)[index].FirstName = updatedPartner.FirstName
			(*c.Partners)[index].MiddleName = updatedPartner.MiddleName
			(*c.Partners)[index].LastName = updatedPartner.LastName
			(*c.Partners)[index].TaxID = updatedPartner.TaxID
			(*c.Partners)[index].Passport = updatedPartner.Passport
			(*c.Partners)[index].DateBirth = updatedPartner.DateBirth
			(*c.Partners)[index].ResidencyStatus = updatedPartner.ResidencyStatus
			(*c.Partners)[index].Address = updatedPartner.Address

			break
		}
	}
}

func (c *Company) UpdatePartnerKYC(dni string, kyc *KYC) {
	for index, partner := range *c.Partners {
		if partner.DNI == dni {
			(*c.Partners)[index].KYC = kyc
			break
		}
	}

}

func (c *Company) GetIDNumber() string {
	return c.RegisterNumber
}

func (c *Company) SetDocument(document Document, dni string) *systemdomain.Error {
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

	if len(*c.Partners) == 0 {
		fmt.Println("Company partners not found")
		return systemdomain.NewError(404, "Company partners not found")

	}

	// assign or updated document to partner
	for _, partner := range *c.Partners {
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

func (c *Company) GetPhoneNumber() string {
	return c.PhoneNumber
}

func (c *Company) GetCompanyData() Company {
	return *c
}

func (c *Company) GetInvestmentProfile() *InvestmentProfile {
	return c.InvestmentProfile
}

func (c *Company) GetKYCProfile() *KYCProfile {
	return c.KYCProfile
}

func (c *Company) GetPartners() []Individual {
	return *c.Partners
}

func (c *Company) GetDocuments() []Document {
	return c.Documents
}

func (c *Company) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":              c.Name,
		"registerNumber":    c.RegisterNumber,
		"naics":             c.NAICS,
		"naicsDescription":  c.NAICSDescription,
		"companyType":       c.CompanyType,
		"establishedDate":   c.EstablishedDate,
		"webSite":           c.WebSite,
		"registeredAddress": c.RegisteredAddress,
		"physicalAddress":   c.PhysicalAddress,
		"phoneCountry":      c.PhoneCountry,
		"phoneNumber":       c.PhoneNumber,
		"documents":         c.Documents,
		"kyc":               c.KYC,
		"investmentProfile": c.InvestmentProfile,
		"partners":          c.Partners,
	}
}
