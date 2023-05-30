package search

type IndexObject struct {
	Service string `json:"service"`

	Type string `json:"type"`

	Provider string `json:"provider,omitempty"`

	Picture string `json:"picture,omitempty"`

	Slug        string `json:"slug,omitempty"`
	Name        string `json:"name,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	Genres    []string `json:"genres,omitempty"`
	Persons   []string `json:"persons,omitempty"`
	Countries []string `json:"countries,omitempty"`

	HasGenres    bool `json:"hasGenres,omitempty"`
	HasPersons   bool `json:"hasPersons,omitempty"`
	HasCountries bool `json:"hasCountries,omitempty"`

	Year      string `json:"year,omitempty"`
	YearEnd   string `json:"yearEnd,omitempty"`
	YearStart string `json:"yearStart,omitempty"`

	AgeRestriction string `json:"ageRestriction,omitempty"`

	IsActive bool `json:"isActive"`
}
