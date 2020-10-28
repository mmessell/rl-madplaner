package models

import (
	"fmt"
	"os"

	"../log"
	. "../types"
)

type Madplan struct {
	Filename   string
	Name       string
	Persons    Double
	Weeks      []Double
	Dagsplaner []Dagsplan
}

func GetMadplan(fn string, mpDir string, recipeDir string) Madplan {
	log.Start("madplan", "GetMadplan")

	mpFile := getMadplanFile(mpDir, fn)
	dagsplaner := getDagsplaner(recipeDir, mpFile.Dagsplaner)

	log.End("madplan", "GetMadplan")

	mp := Madplan{
		Filename:   fn,
		Name:       mpFile.Name,
		Persons:    mpFile.Persons,
		Weeks:      mpFile.Weeks,
		Dagsplaner: dagsplaner,
	}

	mp.validateMadplan()

	return mp
}

func (mp Madplan) GetUniqueBreakfastRecipes() []Recipe {
	log.Start("madplan", "GetUniqueBreakfastRecipes")

	var recipes []Recipe
	for _, dp := range mp.Dagsplaner {
		recipes = append(recipes, dp.Breakfast)
	}

	urs := getUniqueRecipes(recipes)

	log.End("madplan", "GetUniqueBreakfastRecipes")
	return urs
}

func (mp Madplan) GetUniqueLunchRecipes() []Recipe {
	log.Start("madplan", "GetUniqueLunchRecipes")

	var recipes []Recipe
	for _, dp := range mp.Dagsplaner {
		recipes = append(recipes, dp.Lunch)
	}

	urs := getUniqueRecipes(recipes)

	log.End("madplan", "GetUniqueLunchRecipes")
	return urs
}

func (mp Madplan) GetUniqueDinnerRecipes() []Recipe {
	log.Start("madplan", "GetUniqueDinnerRecipes")

	var recipes []Recipe
	for _, dp := range mp.Dagsplaner {
		recipes = append(recipes, dp.Dinner)
	}

	urs := getUniqueRecipes(recipes)

	log.End("madplan", "GetUniqueDinnerRecipes")
	return urs
}

func getUniqueRecipes(recipes []Recipe) []Recipe {
	recipeMap := make(map[string]Recipe)
	for _, recipe := range recipes {
		recipeMap[recipe.Path] = recipe
	}

	return recipeMapValuesToSlice(recipeMap)
}

func (mp Madplan) GetShoppingList() []Ingredient {
	log.Start("madplan", "GetShoppingList")

	ingMap := make(map[string]Ingredient)

	rcs := mp.GetRecipeCounts()
	for _, rc := range rcs {

		for _, ing := range rc.Recipe.Ingredients {
			if v, ok := ingMap[ing.Name]; ok {
				err := ing.validateUnitIsSame(v)

				if err != nil {
					log.Error(err, "GetShoppingList")
					printAllRecipesWithIngredient(rcs, ing)
					os.Exit(1)
				}
				v.Amount = v.Amount + ing.Amount*Double(rc.Count)
			} else {
				ingMap[ing.Name] = Ingredient{
					Amount:       ing.Amount * Double(rc.Count),
					Unit:         ing.Unit,
					Name:         ing.Name,
					ShoppingList: ing.ShoppingList,
				}
			}
		}
	}

	ingredients := sortIngredients(ingMap)

	log.End("madplan", "GetShoppingList")
	return ingredients
}

func printAllRecipesWithIngredient(rcs map[string]RecipeCount, ing Ingredient) {
	for _, rc := range rcs {
		for _, i := range rc.Recipe.Ingredients {
			if i.Name == ing.Name {
				fmt.Println(rc.Recipe.Path, i.Unit)
			}
		}
	}
}

func (mp Madplan) validateMadplan() {
	recipeCounts := mp.getRecipeCountMap()

	for _, rc := range recipeCounts {
		if int(rc.Count)*int(mp.Persons)%int(rc.Recipe.Portions) != 0 {
			fmt.Println("Der er mismatch mellem antal portioner og antal gange opskriften for", rc.Path, "forekommer")
			os.Exit(1)
		}
	}
}

func (mp Madplan) GetRecipeCounts() map[string]RecipeCount {
	log.Start("madplan", "GetRecipeCounts")

	recipeCounts := mp.getRecipeCountMap()

	for path, rc := range recipeCounts {
		rc.Count = rc.Count * mp.Persons / rc.Portions
		recipeCounts[path] = rc
	}

	log.End("madplan", "GetRecipeCounts")
	return recipeCounts
}

func (mp Madplan) getRecipeCountMap() map[string]RecipeCount {
	recipeCounts := make(map[string]RecipeCount)

	for _, dagsplan := range mp.Dagsplaner {
		recipeCounts = addRecipe(recipeCounts, dagsplan.Breakfast)
		recipeCounts = addRecipe(recipeCounts, dagsplan.Lunch)
		recipeCounts = addRecipe(recipeCounts, dagsplan.Dinner)
	}
	return recipeCounts
}

func addRecipe(recipeCount map[string]RecipeCount, recipe Recipe) map[string]RecipeCount {
	if v, ok := recipeCount[recipe.Path]; ok {
		recipeCount[recipe.Path] = RecipeCount{
			Recipe: recipe,
			Count:  v.Count + 1,
		}
	} else {
		recipeCount[recipe.Path] = RecipeCount{
			Recipe: recipe,
			Count:  1,
		}
	}

	return recipeCount
}
