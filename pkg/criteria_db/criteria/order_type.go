package criteria

const (
	ASC  = "asc"
	DESC = "desc"
)

var validTypes = []interface{}{
	ASC,
	DESC,
}

type OrderType struct {
	orderType string
}

func NewOrderType(ot string) *OrderType {
	CheckParamsIsValid(ot, validTypes)
	return &OrderType{
		orderType: ot,
	}
}

func (o *OrderType) GetOrderType() string {
	return o.orderType
}
