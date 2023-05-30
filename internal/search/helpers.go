package search

import "github.com/blevesearch/bleve/v2/search/query"

func useMatchQueryOperatorAnd(q query.Query) query.Query {
	if m, ok := q.(*query.MatchQuery); ok {
		m.SetOperator(query.MatchQueryOperatorAnd)
	}
	return q
}

func useBoost(q query.Query, b float64) query.Query {
	switch v := q.(type) {
	case *query.MatchQuery:
		v.SetBoost(b)
	case *query.MatchPhraseQuery:
		v.SetBoost(b)
	case *query.MatchAllQuery:
		v.SetBoost(b)
	case *query.TermQuery:
		v.SetBoost(b)
	case *query.TermRangeQuery:
		v.SetBoost(b)
	case *query.BooleanQuery:
		v.SetBoost(b)
	}
	return q
}

func newFieldMatchPhraseQuery(field, q string) query.Query {
	m := query.NewMatchPhraseQuery(q)
	m.SetField(field)
	return m
}

func newFieldMatchQuery(field, q string) query.Query {
	m := query.NewMatchQuery(q)
	m.SetField(field)
	return m
}

func newFieldPhraseQuery(field string, terms []string) query.Query {
	m := query.NewPhraseQuery(terms, field)
	return m
}

func newFieldFuzzyQuery(field, term string) query.Query {
	m := query.NewFuzzyQuery(term)
	m.SetField(field)
	return m
}

func newFieldTermQuery(field, q string) query.Query {
	m := query.NewTermQuery(q)
	m.SetField(field)
	return m
}

func newBoolFieldQuery(field string, v bool) query.Query {
	m := query.NewBoolFieldQuery(v)
	m.SetField(field)
	return m
}
