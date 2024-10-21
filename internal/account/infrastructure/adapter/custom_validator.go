package accountadapter

import (
	accountdomain "digital-bank/internal/account/domain"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var ValidateCompanyType validator.Func = func(fl validator.FieldLevel) bool {
	str := fl.Field().String()

	return str == string(accountdomain.PRIVATE) || str == string(accountdomain.PUBLIC) ||
		str == string(accountdomain.HNWI) || str == string(accountdomain.LLC) || str == string(accountdomain.LLP) ||
		str == string(accountdomain.LP) || str == string(accountdomain.S_CORP) || str == string(accountdomain.SOLE_PROP) ||
		str == string(accountdomain.TRUST) || str == string(accountdomain.NON_PROFIT)

}

var ValidateResidencyStatus validator.Func = func(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return str == string(accountdomain.US_CITIZEN) || str == string(accountdomain.RESIDENT_ALIEN) ||
		str == string(accountdomain.NON_RESIDENT_ALIEN)
}

var ValidateAccountType validator.Func = func(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return str == string(accountdomain.INDIVIDUAL_CLIENT) || str == string(accountdomain.COMPANY_CLIENT)
}

func AccountCustomValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		_ = v.RegisterValidation("companyTypeValidate", ValidateCompanyType)

		_ = v.RegisterValidation("residencyStatusValidate", ValidateResidencyStatus)

		_ = v.RegisterValidation("accountTypeValidate", ValidateAccountType)

	}
}
