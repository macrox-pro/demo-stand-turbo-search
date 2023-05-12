package models

type Provider struct {
	ID   uint64 `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name,omitempty"`
}
