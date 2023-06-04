package cmd

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/legion-zver/vss-brain-search/internal/grpc/nlp"

	"google.golang.org/grpc"

	"github.com/legion-zver/vss-brain-search/internal/search"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/legion-zver/vss-brain-search/internal/graphql"

	"github.com/blevesearch/bleve/v2"

	"github.com/spf13/viper"

	_ "github.com/joho/godotenv/autoload"

	"github.com/spf13/cobra"
)

var (
	nlpGrpcAddr string
	indexPath   string
	addr        string
)

var rootCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		index, err := bleve.Open(indexPath)
		if err != nil {
			log.Fatalln(err)
		}
		options := search.Options{
			Index: index,
		}
		if len(nlpGrpcAddr) > 1 {
			cc, err := grpc.Dial(
				nlpGrpcAddr,
				grpc.WithTransportCredentials(
					insecure.NewCredentials(),
				),
			)
			if err != nil {
				log.Println(err)
			} else {
				defer func(cc *grpc.ClientConn) {
					_ = cc.Close()
				}(cc)
				options.NLP = nlp.NewNLPClient(cc)
			}
		}
		engine, err := search.New(options)
		if err != nil {
			log.Fatalln(err)
		}
		server := graphql.NewServer(engine)
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", cors.Default().Handler(server))
		log.Println("Listen GraphQL on", addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalln(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
	_ = viper.ReadInConfig()
	rootCmd.PersistentFlags().StringVarP(
		&addr, "addr", "a", ":9000",
		"listen address for graphql",
	)
	rootCmd.PersistentFlags().StringVarP(
		&indexPath, "index:path", "i", "./index.bleve",
		"index file (default is ./index.bleve)",
	)
	rootCmd.PersistentFlags().StringVar(
		&nlpGrpcAddr, "nlp:grpc:addr", "127.0.0.1:50051",
		"nlp grpc address",
	)
}
