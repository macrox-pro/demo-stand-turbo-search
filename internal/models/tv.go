package models

type TV struct {
	ID   uint64 `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name,omitempty"`

	Picture string `json:"picture,omitempty"`

	Type     *Type     `json:"type,omitempty"`
	Provider *Provider `json:"provider,omitempty"`

	OriginalTitle string `json:"original_title,omitempty"`

	Year      string `json:"year,omitempty"`
	YearEnd   string `json:"year_end,omitempty"`
	YearStart string `json:"year_start,omitempty"`

	Keywords    string `json:"keywords,omitempty"`
	Description string `json:"description,omitempty"`

	AgeRestriction string `json:"age_restriction,omitempty"`

	IsActive bool `json:"is_active,omitempty"`
}
