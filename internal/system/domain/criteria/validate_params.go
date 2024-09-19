package criteria

import (
	"digital-bank/internal/system/domain/criteria/errors"
)

func CheckParamsIsValid(va string, list []interface{}) {
	for _, v := range list {
		if v == va {
			return
		}
	}

	panic(errors.NewInvalidArgumentError("Params " + va + " is valid"))
}
