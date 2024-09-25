package reqaccount

import (
	accountdomain "digital-bank/internal/account/domain"
	systemdomain "digital-bank/internal/system/domain"
)

type (
	ApplicationAccountIndividualRequest struct {
		accountdomain.Individual
		ResidencyStatus string `json:"residencyStatus" binding:"required,residencyStatusValidate"`
	}

	ApplicationAccountCompanyRequest struct {
		accountdomain.Company
		CompanyType string                                `json:"companyType" binding:"required,companyTypeValidate"`
		Partners    []ApplicationAccountIndividualRequest `json:"partners"`
		OwnerRecord systemdomain.AppClientIdentifier      `json:"ownerRecord"`
	}
)