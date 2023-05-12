package models

type Type struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`

	Title string `json:"title,omitempty"`

	SerialContent bool `json:"serial_content,omitempty"`
}
