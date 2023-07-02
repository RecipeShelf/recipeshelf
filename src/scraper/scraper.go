package main

import (
	"time"

	recipePkg "github.com/kkyr/go-recipe/pkg/recipe"
	"github.com/redis/go-redis/v9"
)

func main() {
	importTimeout, err := time.ParseDuration(importQueueTimeout)
	if err != nil {
		ErrorLog.Fatal(err)
	}

	for {
		result, err := rdb.BLPop(ctx, importTimeout, importQueue).Result()
		switch {
		case err == redis.Nil:
			InfoLog.Println("Nothing to scrape")
			return
		case err != nil:
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
	WarningLog.Println("Skipped", url)
	json, _ := MarshalErr(url, err)
	rdb.RPush(ctx, deadLetterQueue, json)
}
