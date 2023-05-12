package models

type TV struct {
	ID   uint64 `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name,omitempty"`

	Type     *Type     `json:"type,omitempty"`
	Provider *Provider `json:"provider,omitempty"`

	OriginalTitle string `json:"original_title,omitempty"`

	Keywords    string `json:"keywords,omitempty"`
	Description string `json:"description,omitempty"`

	IsActive bool `json:"is_active,omitempty"`
}
