package premier

type PersonData struct {
	ID          int    `json:"id"`
	Type        *Type  `json:"type,omitempty"`
	Name        string `json:"name"`
	Picture     string `json:"picture,omitempty"`
	Content     string `json:"content,omitempty"`
	Description string `json:"description,omitempty"`

	IsActive bool `json:"is_active,omitempty"`

	CanSubscribe bool `json:"can_subscribe,omitempty"`
}

type Person struct {
	ID int `json:"id"`

	PersonData PersonData `json:"person"`
	LinkType   *Type      `json:"link_type,omitempty"`

	Role string `json:"role,omitempty"`

	Weight int `json:"weight,omitempty"`
}
