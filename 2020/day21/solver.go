package day21

import (
	"fmt"
	"mrose.de/aoc/utility"
	"sort"
	"strings"
)

/*
Rules:
	Input: Ingredients followed by some or all allergens the recipe contains.
	Ingredient: 1 or 0 allergens
	Allergen: Contained in exactly one recipe / not always marked
*/
type recipe struct {
	ingredients []string
	allergens   []string
}

func Solve() (result int) {
	foodList := foodListFromInput(prepareInput())
	//map with all allergens and possible containing ingredients
	possibleAllergenMap := foodAllergensPoss(foodList)
	allergenMap := findIngredientForAllergen(possibleAllergenMap, foodList)
	ingredientsWithoutAllergens := ingredientsWithoutAllergens(allergenMap, foodList)
	return appearanceOfIngredientsWithoutAllergens(foodList, ingredientsWithoutAllergens)
}
//return the number of appearances for ingredients without allergens
func appearanceOfIngredientsWithoutAllergens(list []recipe, withoutAllergens []string) (result int) {
	for _, v := range list {
		for _, ingredient := range v.ingredients {
			if utility.ContainsString(withoutAllergens, ingredient) {
				result++
			}
		}
	}
	return
}
//find all ingredients without allergens
func ingredientsWithoutAllergens(allergenMap map[string][]string, foodList []recipe) (ingredientsWithoutAllergen []string) {
	ingredientList := allIngredients(foodList)
	neverContained := true
	for _, ingredient := range ingredientList {
		neverContained = true
		for k := range allergenMap {
			if utility.ContainsString(allergenMap[k], ingredient) {
				neverContained = false
				continue
			}
		}
		if neverContained {
			ingredientsWithoutAllergen = append(ingredientsWithoutAllergen, ingredient)
		}
	}
	return
}
//remove ingredients from the allergenMap of ingredient x, if it is already uniquely identified for y
func makeUnique(allergenMap *map[string][]string) {
	run := true
	for run {
		for k := range *allergenMap {
			if len((*allergenMap)[k]) == 1 {
				for key := range *allergenMap {
					if key != k {
						for i, v := range (*allergenMap)[key] {
							if v == (*allergenMap)[k][0] {
								(*allergenMap)[key] = utility.RemoveIndexStr((*allergenMap)[key], i)
							}
						}
					}
				}
			}
		}
		allUnique := true
		for k := range *allergenMap {
			if len((*allergenMap)[k]) != 1 {
				allUnique = false
			}
		}
		if allUnique {
			run = false
		}
	}
}

//list all ingredients
func allIngredients(list []recipe) (ingredients []string) {
	for _, food := range list {
		for _, ingredient := range food.ingredients {
			if !utility.ContainsString(ingredients, ingredient) {
				ingredients = append(ingredients, ingredient)
			}
		}
	}
	return
}

//map with key=allergen and v=list of possible ingredients containing k
func foodAllergensPoss(list []recipe) (allergenMap map[string][]string) {
	allergenMap = allergens(list)
	for key := range allergenMap {
		for _, vFood := range list {
			if utility.ContainsString(vFood.allergens, key) {
				for _, ingredient := range vFood.ingredients {
					if !utility.ContainsString(allergenMap[key], ingredient) {
						allergenMap[key] = append(allergenMap[key], ingredient)
					}
				}
			}
		}
	}
	return
}

//any ingredient can only contain one allergen. Find the unique ingredient for each allergen
func findIngredientForAllergen(allergens map[string][]string, foodList []recipe) map[string][]string {
	for allergen := range allergens {
		for _, currentFood := range foodList {
			if utility.ContainsString(currentFood.allergens, allergen) {
				for index := len(allergens[allergen]) - 1; index >= 0; index-- {
					ingredient := allergens[allergen][index]
					if utility.ContainsString(currentFood.ingredients, ingredient) {
						continue
					}
					allergens[allergen] = utility.RemoveIndexStr(allergens[allergen], index)
				}
				if len(allergens[allergen]) == 1 {
					for key := range allergens {
						if key != allergen {
							for i, v := range allergens[key] {
								if v == allergens[allergen][0] {
									allergens[key] = utility.RemoveIndexStr(allergens[key], i)
								}
							}
						}
					}
				}
			}
		}
	}
	makeUnique(&allergens)
	orderAlphabetically(&allergens)
	return allergens
}

//Part 2. Sort allergens alphabetically. Print the ingredients sorted by allergen and separated by ','
func orderAlphabetically(m *map[string][]string) {
	allergens := make([]string, 0)
	for k := range *m {
		allergens = append(allergens, k)
	}
	sort.Strings(allergens)
	result := ""
	for i, v := range allergens {
		result += (*m)[v][0]
		if i < len(allergens)-1 {
			result+=","
		}
	}
	fmt.Print(result + "\n")
}

//Initialize map with all allergens as keys. Values are all empty []string
func allergens(list []recipe) (result map[string][]string) {
	allergens := make([]string, 0)
	result = make(map[string][]string)
	for _, v := range list {
		for _, allergen := range v.allergens {
			if !utility.ContainsString(allergens, allergen) {
				allergens = append(allergens, allergen)
			}
		}
	}
	for _, v := range allergens {
		result[v] = make([]string, 0)
	}
	return
}

func prepareInput() []string {
	input := utility.Input2020Day21()
	input = strings.Replace(input, ",", "", -1)
	input = strings.Replace(input, ")", "", -1)
	return utility.StrArr(input)
}

func foodListFromInput(arr []string) (foodList []recipe) {
	foodList = make([]recipe, 0)
	isAllergen := false
	for _, inputRecipe := range arr {
		ingredients := make([]string, 0)
		allergens := make([]string, 0)
		for _, part := range strings.Split(inputRecipe, " ") {
			if !isAllergen {
				if !strings.Contains(part, "(") {
					ingredients = append(ingredients, part)
				} else {
					isAllergen = true
					continue
				}
			} else {
				allergens = append(allergens, part)
			}
		}
		foodList = append(foodList,
			recipe{
				ingredients: ingredients,
				allergens:   allergens,
			},
		)
		isAllergen = false
	}
	return
}

/*
--- Day 21: Allergen Assessment ---
You reach the train's last stop and the closest you can get to your vacation island without getting wet. There aren't even any boats here, but nothing can stop you now: you build a raft. You just need a few days' worth of recipe for your journey.

You don't speak the local language, so you can't read any ingredients lists. However, sometimes, allergens are listed in a language you do understand. You should be able to use this information to determine which ingredient contains which allergen and work out which foods are safe to take with you on your trip.

You start by compiling a list of foods (your puzzle input), one recipe per line. Each line includes that recipe's ingredients list followed by some or all of the allergens the recipe contains.

Each allergen is found in exactly one ingredient. Each ingredient contains zero or one allergen. Allergens aren't always marked; when they're listed (as in (contains nuts, shellfish) after an ingredients list), the ingredient that contains each listed allergen will be somewhere in the corresponding ingredients list. However, even if an allergen isn't listed, the ingredient that contains that allergen could still be present: maybe they forgot to label it, or maybe it was labeled in a language you don't know.

For example, consider the following list of foods:

mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
The first recipe in the list has four ingredients (written in a language you don't understand): mxmxvkd, kfcds, sqjhc, and nhms. While the recipe might contain other allergens, a few allergens the recipe definitely contains are listed afterward: dairy and fish.

The first step is to determine which ingredients can't possibly contain any of the allergens in any recipe in your list. In the above example, none of the ingredients kfcds, nhms, sbzzf, or trh can contain an allergen. Counting the number of times any of these ingredients appear in any ingredients list produces 5: they all appear once each except sbzzf, which appears twice.

Determine which ingredients cannot possibly contain any of the allergens in your list. How many times do any of those ingredients appear?
--- Part Two ---
Now that you've isolated the inert ingredients, you should have enough information to figure out which ingredient contains which allergen.

In the above example:

mxmxvkd contains dairy.
sqjhc contains fish.
fvjkl contains soy.
Arrange the ingredients alphabetically by their allergen and separate them by commas to produce your canonical dangerous ingredient list. (There should not be any spaces in your canonical dangerous ingredient list.) In the above example, this would be mxmxvkd,sqjhc,fvjkl.

Time to stock your raft with supplies. What is your canonical dangerous ingredient list?
*/
