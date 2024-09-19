package criteria

type Criteria struct {
	filters     []Filter
	order       *Order
	limit       *int
	currentPage int
	offset      int
}

func NewCriteria(order *Order, limit *int, currentPage int) *Criteria {
	var orderValue Order
	if order != nil {
		orderValue = *order
	}

	var limitValue int = 1
	if limit != nil {
		limitValue = *limit
	}

	var o int

	if currentPage > 0 {
		o = (currentPage - 1) * limitValue
	} else {
		o = 0
	}

	return &Criteria{
		order:       &orderValue,
		limit:       &limitValue,
		currentPage: currentPage,
		offset:      o,
	}
}

func (c *Criteria) SetFilter(filter Filter) {
	c.filters = append(c.filters, filter)
}

func (c *Criteria) HasFilters() bool {
	return c.filters != nil && len(c.filters) > 0
}

func (c *Criteria) GetFilters() *[]Filter {
	return &c.filters
}

func (c *Criteria) GetLimit() int {
	return *c.limit
}

func (c *Criteria) GetOrdering() Order {
	return *c.order
}

func (c *Criteria) GetOffset() int {
	return c.offset
}

func (c *Criteria) GetNextPage(docs int64) *int {
	var nextPage int

	if int64(c.currentPage*c.GetLimit()) > docs {
		nextPage = 0
		return &nextPage
	}

	nextPage = c.currentPage + 1
	return &nextPage

}

func (c *Criteria) GetPrevPage(nextPage *int) *int {

	if nextPage == nil {
		return nil
	}

	prevPage := *(nextPage) - 1
	return &prevPage

}
