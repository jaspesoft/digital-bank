package criteria

type Order struct {
	field     string
	orderType *OrderType
}

func NewOrder(f string, ot string) *Order {
	return &Order{
		field:     f,
		orderType: NewOrderType(ot),
	}
}

func (o *Order) GetField() string {
	return o.field
}

func (o *Order) IsAsc() bool {
	return o.GetDirection() == ASC
}

func (o *Order) IsDesc() bool {
	return o.GetDirection() == DESC
}

func (o *Order) GetDirection() string {
	return o.orderType.GetOrderType()
}
