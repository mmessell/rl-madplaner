package models

import (
	"encoding/json"

	"../io"
	. "../types"
	"../validator"
)

type Recipe struct {
	Path        string
	Name        string
	Calories    Double
	Image       string
	Portions    Double
	Ingredients []Ingredient
	Procedure   []string
}

func (rp Recipe) GetCaloriesPerPortions() Double {
	return rp.Calories / rp.Portions
}

func recipeMapValuesToSlice(recipeMap map[string]Recipe) []Recipe {
	mapValues := []Recipe{}
	for _, v := range recipeMap {
		mapValues = append(mapValues, v)
	}
	return mapValues
}

func getRecipe(dir string, filename string) Recipe {
	b := io.GetFileByteArray(dir, filename)
	recipe, err := bytesToRecipe(b)

	validator.ValidateError(err, "Couldn't Unmarshal bytes to recipe: "+dir+"/"+filename)
	recipe.Path = dir + "/" + filename
	return recipe
}

func bytesToRecipe(b []byte) (Recipe, error) {
	var recipe = Recipe{}
	err := json.Unmarshal(b, &recipe)
	return recipe, err
}
