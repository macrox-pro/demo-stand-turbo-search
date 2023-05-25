package search

import (
	"context"
	"fmt"
	"log"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/search/query"

	"github.com/legion-zver/premier-one-bleve-search/internal/grpc/nlp"
)

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
	if useNLP != nil && *useNLP && e.NLP != nil {
		result, err := e.NLP.Parse(ctx, &nlp.Doc{Text: q})
		if err != nil {
			log.Println(err)
			return
		}
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
			conj.AddQuery(newFieldTermQuery("type", "фильм"))
		case "serials_by_person":
			conj.AddQuery(newFieldTermQuery("type", "сериал"))
		case "shows_by_person":
			conj.AddQuery(newFieldTermQuery("type", "шоу"))
		default:
			log.Println("intent", result.Intent.Name, "not supported! skip brain search query")
			return
		}
		for _, entity := range result.Entities {
			switch entity.Type {
			case "person":
				conj.AddQuery(
					newFieldMatchQuery("keywords", entity.NormalValue),
				)
			}
		}
		if len(conj.Conjuncts) > 0 {
			sq = conj
		}
		//TODO: Deprecated code
		//if scenario != nil {
		//	log.Println(q, *scenario)
		//	conj := bleve.NewConjunctionQuery()
		//	if phrase := strings.Join(scenario.PhraseWords, " "); len(phrase) > 0 {
		//		conj.AddQuery(
		//			bleve.NewDisjunctionQuery(
		//				newFieldMatchQuery("name", phrase),
		//				newFieldMatchPhraseQuery("title", phrase),
		//				newFieldMatchPhraseQuery("description", phrase),
		//			))
		//	}
		//	if len(scenario.Types) > 0 {
		//		typesDis := bleve.NewDisjunctionQuery()
		//		if len(scenario.Types) > 1 {
		//			for _, t := range scenario.Types {
		//				typesDis.AddQuery(newFieldTermQuery("type", t))
		//			}
		//			conj.AddQuery(typesDis)
		//		} else {
		//			conj.AddQuery(newFieldTermQuery("type", scenario.Types[0]))
		//		}
		//	}
		//	if len(scenario.Year) > 0 {
		//		conj.AddQuery(
		//			bleve.NewDisjunctionQuery(
		//				newFieldTermQuery("year", scenario.Year),
		//				newFieldTermQuery("yearEnd", scenario.Year),
		//				newFieldTermQuery("yearStart", scenario.Year),
		//			),
		//		)
		//	}
		//	for _, person := range scenario.Persons {
		//		conj.AddQuery(
		//			bleve.NewDisjunctionQuery(
		//				newFieldMatchQuery("keywords", person.Text),
		//				newFieldMatchQuery("keywords", person.Lemma),
		//			),
		//		)
		//	}
		//	sq = conj
		//}
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
