package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/legion-zver/premier-one-bleve-search/internal/models/premier"

	"github.com/legion-zver/premier-one-bleve-search/internal/search"

	"github.com/blevesearch/bleve/v2"
	"github.com/corpix/uarand"
	"github.com/go-resty/resty/v2"
	"github.com/legion-zver/premier-one-bleve-search/internal/utils"
	"github.com/spf13/cobra"
)

var (
	cachePath string
)

func init() {
	indexSyncCmd.PersistentFlags().StringVarP(
		&cachePath, "cache:path", "c", "./cache.resty",
		"cache path (default is ./cache.resty)",
	)

	rootCmd.AddCommand(indexInitCmd)
	rootCmd.AddCommand(indexSyncCmd)
}

var indexInitCmd = &cobra.Command{
	Use:   "index:init",
	Short: "create search index",
	Run: func(cmd *cobra.Command, args []string) {
		index, err := bleve.New(indexPath, search.NewIndexMapping())
		if err != nil {
			log.Fatalln(err)
		}
		defer func(index bleve.Index) {
			_ = index.Close()
		}(index)
		log.Println("Ready!")
	},
}

var indexSyncCmd = &cobra.Command{
	Use:   "index:sync",
	Short: "sync search index",
	Run: func(cmd *cobra.Command, args []string) {
		index, err := bleve.Open(indexPath)
		if err != nil {
			log.Fatalln(err)
		}
		defer func(index bleve.Index) {
			_ = index.Close()
		}(index)

		// Create cache dir
		_ = os.MkdirAll(cachePath, os.ModePerm)

		page, limit, retryCount := 1, 100, 0
		client := resty.New()
		for {
			var (
				url      = fmt.Sprintf("https://premier.one/api/metainfo/tv/?page=%d&limit=%d&picture_type=card_group&style=portrait&device=web", page, limit)
				hashPath = path.Join(cachePath, fmt.Sprintf("%s.json", utils.SHA1(url)))
				body     []byte
				isCache  bool
			)
			if _, err := os.Stat(hashPath); err == nil || os.IsExist(err) {
				body, err = os.ReadFile(hashPath)
				if err == nil {
					isCache = true
				}
			}
			if !isCache {
				resp, err := client.R().
					SetDoNotParseResponse(false).
					SetHeader("Accept", "application/json").
					SetHeader("User-Agent", uarand.GetRandom()).
					Get(url)
				if err != nil {
					retryCount++
					if retryCount > 3 {
						log.Fatalln(err)
					}
					time.Sleep(1000)
				}
				if retryCount > 0 {
					retryCount = 0
				}
				body = resp.Body()
				if err = os.WriteFile(hashPath, body, os.ModePerm); err != nil {
					fmt.Println(err)
				}
			}
			var payload premier.Response[premier.TV]
			if err = json.Unmarshal(body, &payload); err != nil {
				log.Println(err)
			}
			if isCache {
				fmt.Println("Cached iteration", page, "with", len(payload.Results), "items...")
			} else {
				fmt.Println("Iteration", page, "with", len(payload.Results), "items...")
			}
			batch := index.NewBatch()
			for _, item := range payload.Results {
				if len(item.Externals) < 1 {
					continue
				}
				object := search.IndexObject{
					Slug:           strings.TrimSpace(item.Slug),
					Name:           strings.TrimSpace(item.Name),
					Year:           strings.TrimSpace(item.Year),
					Title:          strings.TrimSpace(item.OriginalTitle),
					YearEnd:        strings.TrimSpace(item.YearEnd),
					Picture:        strings.TrimSpace(item.Picture),
					Keywords:       strings.Split(strings.TrimSpace(item.Keywords), " "),
					IsActive:       item.IsActive,
					YearStart:      strings.TrimSpace(item.YearStart),
					Description:    strings.TrimSpace(item.Description),
					AgeRestriction: strings.TrimSpace(item.AgeRestriction),
				}
				if item.Type != nil {
					object.Type = item.Type.Title
				}
				if item.Provider != nil {
					object.Provider = item.Provider.Name
				}
				if err := batch.Index(fmt.Sprint(item.ID), object); err != nil {
					log.Println(err)
				}
			}
			if err := index.Batch(batch); err != nil {
				log.Println(err)
			}
			if !payload.HasNext {
				break
			}
			page++
		}
		log.Println("Ready!")
	},
}
