package kit

// KV key value type
type KV map[string]interface{}

const (
	SortRequestMissingFirst = "first"
	SortRequestMissingLast  = "last"
)

// SortRequest request to sort data
type SortRequest struct {
	Field   string `json:"field"`
	Asc     bool   `json:"asc"`
	Missing string `json:"missing"`
}

// PagingRequest paging request
type PagingRequest struct {
	Size   int            `json:"size"`
	Index  int            `json:"index"`
	SortBy []*SortRequest `json:"sortBy"`
}

// PagingResponse paging response
type PagingResponse struct {
	Total int `json:"total"`
	Index int `json:"index"`
}
