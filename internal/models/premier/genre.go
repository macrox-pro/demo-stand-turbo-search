package premier

type Genre struct {
	ID int `json:"id"`

	Name string `json:"name"`

	Main bool   `json:"main,omitempty"`
	Url  string `json:"url,omitempty"`
}
