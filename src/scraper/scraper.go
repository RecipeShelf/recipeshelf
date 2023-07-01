package main

import (
	"time"

	recipePkg "github.com/kkyr/go-recipe/pkg/recipe"
)

func main() {
	InfoLog.Println("Import Queue:", importQueue, "Export Queue:", exportQueue, "Dead Letter Queue:", deadLetterQueue)

	for {
		result, err := rdb.BLPop(ctx, 0*time.Second, importQueue).Result()
		if err != nil {
			ErrorLog.Fatal(err)
		}

		url := result[1]
		InfoLog.Println("Scraping", url)

		recipe, err := recipePkg.ScrapeURL(url)
		if err != nil {
			HandleScrapeError(url, err)
			continue
		}

		recipeJson, err := Marshal(url, recipe)
		if err != nil {
			HandleScrapeError(url, err)
			continue
		}
		rdb.RPush(ctx, exportQueue, recipeJson)
		InfoLog.Println("Scraped", url)
	}
}

func HandleScrapeError(url string, err error) {
	rdb.RPush(ctx, deadLetterQueue, url)
	WarningLog.Println(err)
}
