package search

import "github.com/blevesearch/bleve/v2/search/query"

func newFieldMatchQuery(field, q string, b float64) query.Query {
	m := query.NewMatchQuery(q)
	m.SetField(field)
	if b != 0 {
		m.SetBoost(b)
	}
	return m
}
