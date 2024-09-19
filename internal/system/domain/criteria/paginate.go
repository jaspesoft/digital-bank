package criteria

type Paginate struct {
	Results  interface{} `json:"results"`
	NextPage *int        `json:"nextPage"`
	PrevPage *int        `json:"prevPage"`
	Count    int64       `json:"count"`
}
