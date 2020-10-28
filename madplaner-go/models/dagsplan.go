package models

import (
	"../log"
)

type Dagsplan struct {
	Day       string
	Breakfast Recipe
	Lunch     Recipe
	Dinner    Recipe
}

func getDagsplaner(recipeDir string, dagsplanFiles []DagsplanFile) []Dagsplan {
	log.Trace("--- dagsplan.go start getDagsplaner ---")

	dagsplaner := []Dagsplan{}
	for _, dpFile := range dagsplanFiles {
		dagsplaner = append(dagsplaner, dpFile.toDagsplan(recipeDir))
	}

	log.Trace("--- dagsplan.go end getDagsplaner ---")
	return dagsplaner
}

func (dpFile DagsplanFile) toDagsplan(dir string) Dagsplan {
	return Dagsplan{
		Day:       dpFile.Day,
		Breakfast: getRecipe(dir, dpFile.Breakfast),
		Lunch:     getRecipe(dir, dpFile.Lunch),
		Dinner:    getRecipe(dir, dpFile.Dinner),
	}
}
