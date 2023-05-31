package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/legion-zver/vss-brain-search/internal/graphql/graph"
	"github.com/legion-zver/vss-brain-search/internal/search"
)

func NewServer(engine search.Engine) *handler.Server {
	return handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					SearchEngine: engine,
				},
			},
		),
	)
}
