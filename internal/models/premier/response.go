package premier

type Response[T any] struct {
	HasNext     bool    `json:"has_next,omitempty"`
	Next        *string `json:"next,omitempty"`
	HasPrevious bool    `json:"has_previous,omitempty"`
	Previous    *string `json:"previous,omitempty"`

	Page uint64 `json:"page,omitempty"`

	PerPage uint64 `json:"per_page,omitempty"`

	Results []T `json:"results"`
}
