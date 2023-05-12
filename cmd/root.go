package cmd

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/legion-zver/premier-one-bleve-search/internal/graphql"

	"github.com/blevesearch/bleve/v2"

	"github.com/spf13/viper"

	_ "github.com/joho/godotenv/autoload"

	"github.com/spf13/cobra"
)

var (
	indexPath string
	addr      string
)

var rootCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		index, err := bleve.Open(indexPath)
		if err != nil {
			log.Fatalln(err)
		}
		server := graphql.NewServer(index)
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", server)
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
}
