package model

type ResponseSearchConfig struct {
	Source []string `json:"source" elastic:"_source"`
	Size   int      `json:"limit" elastic:"size"`
	From   int      `json:"offset" elastic:"from"`
}

type SearchResults struct {
	Pagination Paging   `json:"pagination"`
	Data       []string `json:"data,omitempty"`
}

type RangeQueryReq struct {
	From any `json:"from"`
	To   any `json:"to"`
}
