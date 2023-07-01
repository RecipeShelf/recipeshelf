package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/redis/go-redis/v9"
)

type ScrapedRecipe struct {
	Url           string
	Author        string
	Categories    []string
	CookTime      time.Duration
	Cuisine       []string
	Description   string
	ImageURL      string
	Ingredients   []string
	Instructions  []string
	Language      string
	Name          string
	Nutrition     recipe.Nutrition
	PrepTime      time.Duration
	SuitableDiets []recipe.Diet
	TotalTime     time.Duration
	Yields        string
}

var WarningLog = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
var InfoLog = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var ErrorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("redis-host"),
	Password: os.Getenv("redis-password"),
	DB:       0,
})

var importQueue = os.Getenv("import-queue")
var exportQueue = os.Getenv("export-queue")
var deadLetterQueue = os.Getenv("dead-letter-queue")
