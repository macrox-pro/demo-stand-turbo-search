package search

import (
	"context"
	"fmt"
	"log"

	"github.com/blevesearch/bleve/v2/search/query"

	"github.com/legion-zver/premier-one-bleve-search/internal/grpc/nlp"

	"github.com/blevesearch/bleve/v2"
)

type Engine interface {
	Search(ctx context.Context, q string) (*bleve.SearchResult, error)
}

type Options struct {
	Index bleve.Index
	NLP   nlp.NLPClient
}

type engine struct {
	Options
}

func (e *engine) newSearchQuery(ctx context.Context, q string) query.Query {
	if e.NLP != nil {
		result, err := e.NLP.Parse(ctx, &nlp.Doc{Text: q})
		if err != nil {
			log.Println(err)
		} else if len(result.Tokens) > 1 {
			fmt.Println(result)
		}
	}
	return bleve.NewDisjunctionQuery(
		newFieldMatchQuery("title", q, 5),
		newFieldMatchQuery("description", q, 0),
	)
}

func (e *engine) Search(ctx context.Context, q string) (*bleve.SearchResult, error) {
	req := bleve.NewSearchRequestOptions(e.newSearchQuery(ctx, q), 21, 0, false)
	req.Fields = []string{
		"type",
		"slug",
		"name",
		"year",
		"title",
		"yearEnd",
		"picture",
		"provider",
		"isActive",
		"yearStart",
		"description",
		"ageRestriction",
	}
	return e.Index.Search(req)
}

func New(opts Options) (Engine, error) {
	if opts.Index == nil {
		return nil, fmt.Errorf("index is nil")
	}
	return &engine{Options: opts}, nil
}
