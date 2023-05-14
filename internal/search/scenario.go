package search

type Scenario struct {
	Types       []string
	Year        string
	Persons     []string
	PhraseWords []string
}

func (s *Scenario) AddType(v string) *Scenario {
	if s.Types == nil {
		s.Types = make([]string, 0)
	}
	s.Types = append(s.Types, v)
	return s
}

func (s *Scenario) AddPhraseWord(v string) *Scenario {
	if s.PhraseWords == nil {
		s.PhraseWords = make([]string, 0)
	}
	s.PhraseWords = append(s.PhraseWords, v)
	return s
}

func (s *Scenario) AddPersons(v string) *Scenario {
	if s.Persons == nil {
		s.Persons = make([]string, 0)
	}
	s.Persons = append(s.Persons, v)
	return s
}
