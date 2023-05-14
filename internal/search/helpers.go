package search

import "github.com/blevesearch/bleve/v2/search/query"

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
