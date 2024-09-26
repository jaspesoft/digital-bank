package systemdomain

type (
	SystemParameters struct {
		commissions *TransactionFee
	}

	SystemParametersRepository interface {
		GetSystemParameters() (*SystemParameters, error)
	}
)

func NewSystemParameters(commissions *TransactionFee) *SystemParameters {
	return &SystemParameters{
		commissions: commissions,
	}
}

func (s *SystemParameters) GetCommissions() *TransactionFee {
	return s.commissions
}

func (s *SystemParameters) ToMpa() map[string]interface{} {
	return map[string]interface{}{
		"commissions": s.commissions.ToMap(),
	}

}

func SystemParametersFromPrimitive(s map[string]interface{}) *SystemParameters {

	return &SystemParameters{
		commissions: TransactionFeeFromPrimitive(s["commissions"].(map[string]interface{})),
	}
}
