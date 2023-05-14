package search

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/blevesearch/bleve/v2/search/query"

	"github.com/legion-zver/premier-one-bleve-search/internal/grpc/nlp"

	"github.com/blevesearch/bleve/v2"
)

type Engine interface {
	Search(ctx context.Context, q string, isActive *bool) (*bleve.SearchResult, error)
}

type Options struct {
	Index bleve.Index
	NLP   nlp.NLPClient
}

type engine struct {
	Options
}

func (e *engine) detectScenario(ctx context.Context, q string) *Scenario {
	if e.NLP == nil {
		return nil
	}
	parsed, err := e.NLP.Parse(ctx, &nlp.Doc{Text: q})
	if err != nil {
		log.Println(err)
		return nil
	}
	if len(parsed.Tokens) < 2 {
		return nil
	}
	scenario := new(Scenario)
	for _, token := range parsed.Tokens {
		switch token.Lemma {
		case "сериал", "серил", "серия":
			scenario.AddType("сериал")
		case "фильм", "филь", "флм":
			scenario.AddType("фильм")
		default:
			if token.Pos == "PROPN" { // PERSONS
				scenario.AddPersons(token.Lemma)
			} else if token.Pos == "ADJ" { // DATES
				scenario.Year = token.Text
			} else if token.Pos != "ADP" && token.Rel != "fixed" {
				scenario.AddPhraseWord(token.Text)
			}
		}
	}
	return scenario
}

func (e *engine) newSearchQuery(ctx context.Context, q string, isActive *bool) (sq query.Query) {
	def := bleve.NewDisjunctionQuery(
		newFieldMatchQuery("name", q),
		newFieldMatchPhraseQuery("title", q),
		newFieldMatchPhraseQuery("description", q),
	)
	defer func() {
		if sq == nil {
			// use default
			sq = def
		}
		if isActive != nil {
			sq = bleve.NewConjunctionQuery(
				newBoolFieldQuery("isActive", *isActive),
				sq,
			)
		}
	}()
	scenario := e.detectScenario(ctx, q)
	if scenario != nil {
		conj := bleve.NewConjunctionQuery()
		if phrase := strings.Join(scenario.PhraseWords, " "); len(phrase) > 0 {
			conj.AddQuery(
				bleve.NewDisjunctionQuery(
					newFieldMatchQuery("name", phrase),
					newFieldMatchPhraseQuery("title", phrase),
					newFieldMatchPhraseQuery("description", phrase),
				))
		}
		if len(scenario.Types) > 0 {
			typesDis := bleve.NewDisjunctionQuery()
			if len(scenario.Types) > 1 {
				for _, t := range scenario.Types {
					typesDis.AddQuery(newFieldTermQuery("type", t))
				}
				conj.AddQuery(typesDis)
			} else {
				conj.AddQuery(newFieldTermQuery("type", scenario.Types[0]))
			}
		}
		if len(scenario.Year) > 0 {
			conj.AddQuery(bleve.NewDisjunctionQuery(
				newFieldTermQuery("year", scenario.Year),
				newFieldTermQuery("yearEnd", scenario.Year),
				newFieldTermQuery("yearStart", scenario.Year),
			))
		}
		for _, p := range scenario.Persons {
			conj.AddQuery(newFieldMatchQuery("keywords", p))
		}
		sq = conj
	}
	return
}

func (e *engine) Search(ctx context.Context, q string, isActive *bool) (*bleve.SearchResult, error) {
	req := bleve.NewSearchRequestOptions(e.newSearchQuery(ctx, q, isActive), 21, 0, false)
	req.Fields = []string{
		"type",
		"slug",
		"name",
		"year",
		"title",
		"yearEnd",
		"picture",
		"keywords",
		"provider",
		"isActive",
		"yearStart",
		"description",
		"ageRestriction",
	}
	req.SortBy([]string{
		"-_score",
		"-yearEnd",
		"-year",
		"-id",
	})
	return e.Index.Search(req)
}

func New(opts Options) (Engine, error) {
	if opts.Index == nil {
		return nil, fmt.Errorf("index is nil")
	}
	return &engine{Options: opts}, nil
}
