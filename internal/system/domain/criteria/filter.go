package criteria

type (
	Filter struct {
		field string
		value interface{}
		OP    *FilterOperator
	}
)

func NewFilter(Field string, Value interface{}, OP string) *Filter {

	return &Filter{
		field: Field,
		value: Value,
		OP:    NewFilterOperator(OP),
	}
}

func (f *Filter) GetValue() interface{} {
	return f.value
}

func (f *Filter) GetField() string {
	return f.field
}

func (f *Filter) GetOP() string {
	return f.OP.GetOP()
}
