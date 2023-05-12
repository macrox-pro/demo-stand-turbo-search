package cmd

import (
	"log"

	"github.com/blevesearch/bleve/v2"

	"github.com/spf13/viper"

	_ "github.com/joho/godotenv/autoload"

	"github.com/spf13/cobra"
)

var (
	indexPath string
)

var rootCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := bleve.Open(indexPath)
		if err != nil {
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
		&indexPath, "index:path", "i", "./index.bleve",
		"index file (default is ./index.bleve)",
	)
}
