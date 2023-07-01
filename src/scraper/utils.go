package main

import (
	"encoding/json"

	"github.com/kkyr/go-recipe"
)

func Marshal(url string, recipe recipe.Scraper) (string, error) {
	var scrapedRecipe ScrapedRecipe

	scrapedRecipe.Url = url
	author, ok := recipe.Author()
	if ok {
		scrapedRecipe.Author = author
	}
	categories, ok := recipe.Categories()
	if ok {
		scrapedRecipe.Categories = categories
	}
	cookTime, ok := recipe.CookTime()
	if ok {
		scrapedRecipe.CookTime = cookTime
	}
	cuisine, ok := recipe.Cuisine()
	if ok {
		scrapedRecipe.Cuisine = cuisine
	}
	description, ok := recipe.Description()
	if ok {
		scrapedRecipe.Description = description
	}
	imageURL, ok := recipe.ImageURL()
	if ok {
		scrapedRecipe.ImageURL = imageURL
	}
	ingredients, ok := recipe.Ingredients()
	if ok {
		scrapedRecipe.Ingredients = ingredients
	}
	instructions, ok := recipe.Instructions()
	if ok {
		scrapedRecipe.Instructions = instructions
	}
	language, ok := recipe.Language()
	if ok {
		scrapedRecipe.Language = language
	}
	name, ok := recipe.Name()
	if ok {
		scrapedRecipe.Name = name
	}
	nutrition, ok := recipe.Nutrition()
	if ok {
		scrapedRecipe.Nutrition = nutrition
	}
	prepTime, ok := recipe.PrepTime()
	if ok {
		scrapedRecipe.PrepTime = prepTime
	}
	suitableDiets, ok := recipe.SuitableDiets()
	if ok {
		scrapedRecipe.SuitableDiets = suitableDiets
	}
	totalTime, ok := recipe.TotalTime()
	if ok {
		scrapedRecipe.TotalTime = totalTime
	}
	yields, ok := recipe.Yields()
	if ok {
		scrapedRecipe.Yields = yields
	}

	recipeBytes, err := json.Marshal(scrapedRecipe)
	if err != nil {
		return "", err
	}

	return string(recipeBytes), nil
}
