package search

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/search/query"

	"github.com/legion-zver/vss-brain-search/internal/grpc/nlp"
)

var nonNumericRegex = regexp.MustCompile(`\D`)

type Engine interface {
	Search(ctx context.Context, q string, useNLP, isActive *bool) (*bleve.SearchResult, error)
}

type Options struct {
	Index bleve.Index
	NLP   nlp.NLPClient
}

type engine struct {
	Options
}

func (e *engine) newSearchQuery(ctx context.Context, q string, useNLP, isActive *bool) (sq query.Query) {
	defer func() {
		if sq == nil {
			// use default
			sq = bleve.NewDisjunctionQuery(
				newFieldMatchQuery("name", q),
				newFieldMatchPhraseQuery("title", q),
				newFieldMatchPhraseQuery("description", q),
			)
		}
		if isActive != nil {
			sq = bleve.NewConjunctionQuery(
				newBoolFieldQuery("isActive", *isActive),
				sq,
			)
		}
	}()
	if useNLP != nil && *useNLP && e.NLP != nil {
		result, err := e.NLP.Parse(ctx, &nlp.Doc{Text: q})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(q, result)
		if result.Intent == nil {
			return
		}
		if result.Intent.Confidence < 0.2 {
			// TODO: Add to log for next train
			return
		}
		if result.Intent.Name == "out_of_score" {
			return
		}
		conj := bleve.NewConjunctionQuery()
		switch result.Intent.Name {
		case "films_by_person":
			conj.AddQuery(newFieldMatchQuery("type", "фильм"))
		case "serials_by_person":
			conj.AddQuery(newFieldMatchQuery("type", "сериал"))
		case "shows_by_person":
			conj.AddQuery(newFieldMatchQuery("type", "шоу"))
		default:
			log.Println("intent", result.Intent.Name, "not supported! skip brain search query")
			return
		}
		for _, entity := range result.Entities {
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
			sq = conj
		}
	}
	return
}

func (e *engine) Search(ctx context.Context, q string, useNLP, isActive *bool) (*bleve.SearchResult, error) {
	req := bleve.NewSearchRequestOptions(e.newSearchQuery(ctx, q, useNLP, isActive), 21, 0, false)
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
	return e.Index.Search(req)
}

func New(opts Options) (Engine, error) {
	if opts.Index == nil {
		return nil, fmt.Errorf("index is nil")
	}
	return &engine{Options: opts}, nil
}
