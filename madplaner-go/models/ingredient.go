package models

import (
	"errors"
	"sort"
	"strings"

	. "../types"
)

type Ingredient struct {
	Amount       Double
	Unit         string
	Name         string
	ShoppingList bool
}

func (ing1 Ingredient) validateUnitIsSame(ing2 Ingredient) error {

	if strings.ToLower(ing1.Unit) != strings.ToLower(ing2.Unit) {
		return errors.New("Differences in units (" + ing1.Unit + "|" + ing2.Unit + ") for " + ing1.Name)
	}

	return nil
}

func sortIngredients(ingMap map[string]Ingredient) []Ingredient {

	keys := []string{}
	for k := range ingMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	mapValues := []Ingredient{}
	for _, k := range keys {
		mapValues = append(mapValues, ingMap[k])
	}

	return mapValues
}
