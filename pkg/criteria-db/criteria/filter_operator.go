package criteria

const (
	EQUAL       = "=="
	NOT_EQUAL   = "!="
	GT          = ">"
	GTE         = ">="
	LT          = "<"
	LTE         = "<="
	CONTAINS    = "array-contains"
	OR_CONTAINS = "array-contains-any"
	IN          = "in"
	NOT_IN      = "not-in"
)

const (
	Equal       Operator = "EQUAL"
	NotEqual    Operator = "NOT_EQUAL"
	GreaterThan Operator = "GT"
	LessThan    Operator = "LT"
	Contains    Operator = "CONTAINS"
	NotContains Operator = "NOT_CONTAINS"
)

type Operator string

type FilterOperator struct {
	Operator string
}

var validOperator = []interface{}{
	EQUAL,
	NOT_EQUAL,
	GT, GTE, LTE,
	LT,
	CONTAINS,
	OR_CONTAINS,
	IN,
	NOT_IN,
}

func NewFilterOperator(op string) *FilterOperator {
	CheckParamsIsValid(op, validOperator)
	return &FilterOperator{
		Operator: op,
	}
}

func (fo *FilterOperator) GetOP() string {
	return fo.Operator
}
