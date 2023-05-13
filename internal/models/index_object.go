package models

type IndexObject struct {
	Type string `json:"type"`

	Provider string `json:"provider,omitempty"`

	Slug        string `json:"slug,omitempty"`
	Name        string `json:"name,omitempty"`
	Title       string `json:"title,omitempty"`
	Picture     string `json:"picture,omitempty"`
	Description string `json:"description,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	Year      string `json:"year,omitempty"`
	YearEnd   string `json:"yearEnd,omitempty"`
	YearStart string `json:"yearStart,omitempty"`

	AgeRestriction string `json:"ageRestriction,omitempty"`

	IsActive bool `json:"isActive"`
}
