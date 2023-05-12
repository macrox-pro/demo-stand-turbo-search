package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/blevesearch/bleve/v2"
	"github.com/legion-zver/premier-one-bleve-search/internal/graphql/graph"
)

func NewServer(index bleve.Index) *handler.Server {
	return handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					Index: index,
				},
			},
		),
	)
}
