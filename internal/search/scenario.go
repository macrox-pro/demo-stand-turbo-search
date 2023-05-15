package search

type Person struct {
	Text  string
	Lemma string
}

type Scenario struct {
	Types       []string
	Year        string
	Persons     []Person
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

func (s *Scenario) AddPersons(text, lemma string) *Scenario {
	if s.Persons == nil {
		s.Persons = make([]Person, 0)
	}
	s.Persons = append(s.Persons, Person{Text: text, Lemma: lemma})
	return s
}
