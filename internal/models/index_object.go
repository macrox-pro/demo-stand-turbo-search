package models

type IndexObject struct {
	Type string `json:"type"`

	Provider string `json:"provider,omitempty"`

	Name  string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`

	Description string `json:"description,omitempty"`
}
