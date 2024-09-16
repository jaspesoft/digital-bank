package criteria

type Paginate struct {
	Results  interface{} `json:"results"`
	NextPage int         `json:"nextPage"`
	Count    int64       `json:"count"`
}
