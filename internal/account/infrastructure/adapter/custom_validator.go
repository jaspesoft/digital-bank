package adapter

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func CreateCustomValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		_ = v.RegisterValidation("companyType", requests.ValidateSpeedTX)

	}
}
