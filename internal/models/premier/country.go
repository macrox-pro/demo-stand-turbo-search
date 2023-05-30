package premier

type Country struct {
	TwoLetter string `json:"two_letter,omitempty"`
	Name      string `json:"name"`
}
