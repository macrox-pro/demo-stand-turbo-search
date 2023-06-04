package search

import (
	"strconv"

	"github.com/blevesearch/bleve/v2/search"
)

var yearFields = []string{"year", "yearStart", "yearEnd"}

type HitsWithSortByYears search.DocumentMatchCollection

func (h HitsWithSortByYears) Year(i int) int {
	for _, fieldName := range yearFields {
		if v, ok := h[i].Fields[fieldName]; ok && v != nil {
			if year, err := strconv.Atoi(v.(string)); err == nil {
				return year
			}
		}
	}
	return -1
}

func (h HitsWithSortByYears) Len() int { return len(h) }

func (h HitsWithSortByYears) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h HitsWithSortByYears) Less(i, j int) bool {
	if yi := h.Year(i); yi > 0 {
		if yj := h.Year(j); yj > 0 {
			return yi > yj
		}
	}
	return h[i].Score > h[j].Score
}
