package accountadapter

import (
	accountdomain "digital-bank/internal/account/domain"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var ValidateCompanyType validator.Func = func(fl validator.FieldLevel) bool {
	speed := fl.Field().String()
	return speed == string(accountdomain.PRIVATE) || speed == string(accountdomain.PUBLIC) ||
		speed == string(accountdomain.HNWI) || speed == string(accountdomain.LLC) || speed == string(accountdomain.LLP) ||
		speed == string(accountdomain.LP) || speed == string(accountdomain.S_CORP) || speed == string(accountdomain.SOLE_PROP) ||
		speed == string(accountdomain.TRUST) || speed == string(accountdomain.NON_PROFIT)

}

var ValidateResidencyStatus validator.Func = func(fl validator.FieldLevel) bool {
	speed := fl.Field().String()
	return speed == string(accountdomain.US_CITIZEN) || speed == string(accountdomain.RESIDENT_ALIEN) ||
		speed == string(accountdomain.NON_RESIDENT_ALIEN)
}

func AccountCustomValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		_ = v.RegisterValidation("companyTypeValidate", ValidateCompanyType)

		_ = v.RegisterValidation("residencyStatusValidate", ValidateResidencyStatus)

	}
}
