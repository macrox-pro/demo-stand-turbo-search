package search

import (
	"context"
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/search/query"

	"github.com/legion-zver/vss-brain-search/internal/grpc/nlp"
)

var nonNumericRegex = regexp.MustCompile(`\D`)

type Where struct {
	Service *string
	Active  *bool
}

type Engine interface {
	Search(ctx context.Context, text string, w *Where, useNLP *bool) (*bleve.SearchResult, *nlp.Result, error)
}

type Options struct {
	Index bleve.Index
	NLP   nlp.NLPClient
}

type engine struct {
	Options
}

func (e *engine) newSearchQuery(ctx context.Context, text string, w *Where, useNLP *bool) (q query.Query, nlpResult *nlp.Result) {
	defer func() {
		if q == nil {
			// use default
			q = bleve.NewDisjunctionQuery(
				newFieldMatchQuery("name", text),
				newFieldMatchPhraseQuery("title", text),
				newFieldMatchPhraseQuery("description", text),
			)
		}
		if w != nil {
			if w.Active != nil {
				q = bleve.NewConjunctionQuery(
					newBoolFieldQuery("isActive", *w.Active),
					q,
				)
			}
			if w.Service != nil {
				q = bleve.NewConjunctionQuery(
					newFieldTermQuery("service", *w.Service),
					q,
				)
			}
		}
	}()
	if !(useNLP != nil && *useNLP && e.NLP != nil) {
		return
	}
	var err error
	nlpResult, err = e.NLP.Parse(ctx, &nlp.Doc{Text: text})
	if err != nil {
		log.Println(err)
		return
	}
	if nlpResult.Intent == nil ||
		nlpResult.Intent.Confidence < 0.2 ||
		nlpResult.Intent.Name == "out_of_score" {
		nlpResult = nil
		return
	}
	conj := bleve.NewConjunctionQuery()
	switch nlpResult.Intent.Name {
	case "films_by_person":
		conj.AddQuery(newFieldMatchQuery("type", "фильм"))
	case "serials_by_person":
		conj.AddQuery(newFieldMatchQuery("type", "сериал"))
	case "shows_by_person":
		conj.AddQuery(newFieldMatchQuery("type", "шоу"))
	default:
		log.Println("intent", nlpResult.Intent.Name, "not supported! skip brain search query")
		return
	}
	for _, entity := range nlpResult.Entities {
		switch entity.Type {
		case "person":
			conj.AddQuery(
				newBoolFieldQuery("hasPersons", true),
			)
			conj.AddQuery(
				bleve.NewDisjunctionQuery(
					newFieldMatchPhraseQuery("persons", entity.Value),
					newFieldMatchPhraseQuery("persons", entity.NormalValue),
				),
			)
		case "title":
			conj.AddQuery(
				newFieldMatchPhraseQuery("title", entity.Value),
			)
		case "details":
			conj.AddQuery(
				newFieldMatchPhraseQuery("description", entity.Value),
			)
		case "genre":
			conj.AddQuery(
				newBoolFieldQuery("hasGenres", true),
			)
			conj.AddQuery(
				newFieldMatchQuery("genres", entity.NormalValue),
			)
		case "country_production":
			conj.AddQuery(
				newBoolFieldQuery("hasCountries", true),
			)
			conj.AddQuery(
				newFieldMatchQuery("countries", entity.NormalValue),
			)
		case "year_production":
			year := strings.TrimSpace(nonNumericRegex.ReplaceAllString(entity.Value, ""))
			if len(year) > 0 {
				conj.AddQuery(
					bleve.NewDisjunctionQuery(
						newFieldTermQuery("year", year),
						newFieldTermQuery("yearEnd", year),
						newFieldTermQuery("yearStart", year),
					),
				)
			} else {
				log.Println("fail extract number year from string -", entity.Value)
			}
		default:
			log.Println("entity", entity.Type, "not supported! ignore it for search query")
		}
	}
	if len(conj.Conjuncts) > 0 {
		q = conj
	}
	return
}

func (e *engine) Search(ctx context.Context, text string, where *Where, useNLP *bool) (*bleve.SearchResult, *nlp.Result, error) {
	q, nlpResult := e.newSearchQuery(ctx, text, where, useNLP)
	req := bleve.NewSearchRequestOptions(q, 21, 0, false)
	req.Fields = []string{
		"type",
		"slug",
		"name",
		"year",
		"title",
		"genres",
		"service",
		"yearEnd",
		"picture",
		"persons",
		"provider",
		"countries",
		"isActive",
		"yearStart",
		"description",
		"ageRestriction",
	}
	req.SortBy([]string{
		"-_score",
		"-year",
		"-yearEnd",
		"-id",
	})
	res, err := e.Index.Search(req)
	if err != nil {
		return nil, nil, err
	}
	for _, hit := range res.Hits {
		hit.Score = math.Floor(hit.Score*100) / 100
	}
	sort.Sort(HitsWithSortByYears(res.Hits))
	return res, nlpResult, nil
}

func New(opts Options) (Engine, error) {
	if opts.Index == nil {
		return nil, fmt.Errorf("index is nil")
	}
	return &engine{Options: opts}, nil
}
