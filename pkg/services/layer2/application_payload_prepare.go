package layer2

import (
	accountdomain "digital-bank/internal/account/domain"
	"os"
)

func ApplicationPayloadPrepare(a *accountdomain.Account) map[string]interface{} {
	productId := "DEPOSIT_BASIC"
	assetTypeId := "FIAT_MAINNET_USD"

	if os.Getenv("GO_ENV") != "prod" {
		productId = "DEPOSIT_FORT_FIAT"
		assetTypeId = "FIAT_TESTNET_USD"
	}

	if a.GetType() == accountdomain.COMPANY_CLIENT {
		holder, _ := a.GetAccountHolder().(accountdomain.CompanyAccountHolder)

		return map[string]interface{}{
			"application_type": "CORPORATION",
			"account_to_open": map[string]interface{}{
				"account_id": a.GetAccountID(),
				"asset_type": assetTypeId,
				"product_id": productId,
			},
			"terms_and_conditions_accepted": true,
			"customer_id":                   a.GetAccountID(),
			"customer_details": map[string]interface{}{
				"registered_name":   a.GetName(),
				"registered_number": holder.GetIDNumber(),
				"registered_address": map[string]interface{}{
					"address_line1": holder.GetAddress().StreetOne,
					"address_line2": holder.GetAddress().StreetTwo,
					"address_line3": "",
					"postal_code":   holder.GetAddress().PostalCode,
					"city":          holder.GetAddress().City,
					"state":         holder.GetAddress().Region,
					"country_code":  holder.GetAddress().Country,
				},
				"physical_address": map[string]interface{}{
					"address_line1": holder.GetRegisteredAddress().StreetOne,
					"address_line2": holder.GetRegisteredAddress().StreetTwo,
					"address_line3": "",
					"postal_code":   holder.GetRegisteredAddress().PostalCode,
					"city":          holder.GetRegisteredAddress().City,
					"state":         holder.GetRegisteredAddress().Region,
					"country_code":  holder.GetRegisteredAddress().Country,
				},
			},
			"telephone_number":             holder.GetPhoneNumber(),
			"website_address":              holder.GetCompanyData().WebSite,
			"state_of_incorporation":       holder.GetRegisteredAddress().Region,
			"country_of_incorporation":     holder.GetRegisteredAddress().Country,
			"corporate_entity_type":        holder.GetCompanyData().CompanyType,
			"corporate_entity_description": holder.GetCompanyData().CompanyType,
			"email_address":                holder.GetCompanyData().Email,
			"established_on":               holder.GetCompanyData().EstablishedDate,
			"naics":                        holder.GetCompanyData().NAICS,
			"naics_description":            holder.GetCompanyData().NAICSDescription,
			"investment_profile": map[string]interface{}{
				"primary_source_of_funds":              holder.GetInvestmentProfile().PrimarySourceOfFunds,
				"primary_source_of_funds_description":  "",
				"usd_value_of_fiat":                    holder.GetInvestmentProfile().UsdValueOfFiat,
				"monthly_deposits":                     holder.GetInvestmentProfile().MonthlyDeposits,
				"monthly_withdrawals":                  holder.GetInvestmentProfile().MonthlyWithdrawals,
				"monthly_investment_deposit":           holder.GetInvestmentProfile().MonthlyInvestmentDeposit,
				"monthly_investment_withdrawal":        holder.GetInvestmentProfile().MonthlyInvestmentWithdrawal,
				"usd_value_of_crypto":                  holder.GetInvestmentProfile().UsdValueOfCrypto,
				"monthly_crypto_deposits":              holder.GetInvestmentProfile().MonthlyCryptoDeposits,
				"monthly_crypto_withdrawals":           holder.GetInvestmentProfile().MonthlyCryptoWithdrawals,
				"monthly_crypto_investment_deposit":    holder.GetInvestmentProfile().MonthlyCryptoInvestmentDeposit,
				"monthly_crypto_investment_withdrawal": holder.GetInvestmentProfile().MonthlyCryptoInvestmentWithdrawal,
			},
			"kyc_profile": map[string]interface{}{
				"description_of_business_nature":   holder.GetKYCProfile().DescriptionBusinessNature,
				"business_jurisdictions":           holder.GetKYCProfile().BusinessJurisdictions,
				"funds_send_receive_jurisdictions": holder.GetKYCProfile().FundsSendReceiveJurisdictions,
				"engage_in_activities":             holder.GetKYCProfile().EngageInActivities,
				"regulated_status":                 holder.GetKYCProfile().RegulatedStatus,
				"primary_business":                 holder.GetKYCProfile().PrimaryBusiness,
			},
		}
	}

	holder := a.GetAccountHolder().(*accountdomain.Individual)

	resident := "NON_RESIDENT_ALIEN"
	if holder.GetAddress().Country == "" {
		resident = "US_CITIZEN"
	}

	employmentStatus := holder.EmploymentStatus
	if employmentStatus == accountdomain.OTHER {
		employmentStatus = accountdomain.EMPLOYEE
	}

	return map[string]interface{}{
		"application_type": "INDIVIDUAL",
		"account_to_open": map[string]interface{}{
			"account_id": a.GetAccountID(),
			"asset_type": assetTypeId,
			"product_id": productId,
		},
		"customer_id": a.GetAccountID(),
		"customer_details": map[string]interface{}{
			"date_of_birth":     holder.DateBirth,
			"email_address":     holder.Email,
			"first_name":        holder.FirstName,
			"last_name":         holder.LastName,
			"employment_status": employmentStatus,
			"occupation":        holder.Occupation,
			"mailing_address": map[string]interface{}{
				"address_line1": holder.GetAddress().StreetOne,
				"address_line2": holder.GetAddress().StreetTwo,
				"address_line3": "",
				"postal_code":   holder.GetAddress().PostalCode,
				"city":          holder.GetAddress().City,
				"state":         holder.GetAddress().Region,
				"country_code":  holder.GetAddress().Country,
			},
			"investment_profile": map[string]interface{}{
				"primary_source_of_funds":              holder.GetInvestmentProfile().PrimarySourceOfFunds,
				"usd_value_of_fiat":                    holder.GetInvestmentProfile().UsdValueOfFiat,
				"monthly_deposits":                     holder.GetInvestmentProfile().MonthlyDeposits,
				"monthly_withdrawals":                  holder.GetInvestmentProfile().MonthlyWithdrawals,
				"monthly_investment_deposit":           holder.GetInvestmentProfile().MonthlyInvestmentDeposit,
				"monthly_investment_withdrawal":        holder.GetInvestmentProfile().MonthlyInvestmentWithdrawal,
				"usd_value_of_crypto":                  holder.GetInvestmentProfile().UsdValueOfCrypto,
				"monthly_crypto_deposits":              holder.GetInvestmentProfile().MonthlyCryptoDeposits,
				"monthly_crypto_withdrawals":           holder.GetInvestmentProfile().MonthlyCryptoWithdrawals,
				"monthly_crypto_investment_deposit":    holder.GetInvestmentProfile().MonthlyCryptoInvestmentDeposit,
				"monthly_crypto_investment_withdrawal": holder.GetInvestmentProfile().MonthlyCryptoInvestmentWithdrawal,
			},
			"kyc_profile": map[string]interface{}{
				"funds_send_receive_jurisdictions": holder.KYCProfile.FundsSendReceiveJurisdictions,
				"engage_in_activities":             holder.KYCProfile.EngageInActivities,
			},
			"middle_name":          holder.MiddleName,
			"nationality":          holder.GetAddress().Country,
			"passport_number":      holder.Passport,
			"tax_reference_number": holder.TaxID,
			"telephone_number":     holder.PhoneNumber,
			"us_residency_status":  resident,
		},
		"terms_and_conditions_accepted": true,
	}
}

func PartnerPayloadPrepare(partner accountdomain.Individual) map[string]interface{} {
	return map[string]interface{}{
		"individual_type": []string{"BENEFICIAL_OWNER"},
		"first_name":      partner.FirstName,
		"middle_name":     partner.MiddleName,
		"last_name":       partner.LastName,
		"email_address":   partner.Email,
		"mailing_address": map[string]interface{}{
			"address_line1": partner.Address.StreetOne,
			"address_line2": partner.Address.StreetTwo,
			"city":          partner.Address.City,
			"state":         partner.Address.Region,
			"postal_code":   partner.Address.PostalCode,
			"country_code":  partner.Address.Country,
		},
		"telephone_number":     partner.PhoneNumber,
		"tax_reference_number": partner.TaxID,
		"passport_number":      partner.Passport,
		"date_of_birth":        partner.DateBirth,
	}
}
