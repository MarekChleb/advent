package ex21

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const Path = "./solutions/ex21"

func A(lines []string) {
	count := 0

	possibleIngredients := map[string]map[string]struct{}{} // allergen -> ingredients

	for _, l := range lines {
		ingredients, allergens := parseLine(l)

		for _, a := range allergens {
			if _, ok := possibleIngredients[a]; !ok {
				possibleIngredients[a] = map[string]struct{}{}
				for _, i := range ingredients {
					possibleIngredients[a][i] = struct{}{}
				}
			} else {
				temp := map[string]struct{}{}
				for _, i := range ingredients {
					temp[i] = struct{}{}
				}

				toRemove := []string{}
				for i := range temp {
					if _, ok := possibleIngredients[a][i]; !ok {
						toRemove = append(toRemove, i)
					}
				}

				for _, s := range toRemove {
					delete(temp, s)
				}
				possibleIngredients[a] = temp
			}
		}
	}

	reducedIngredients := reduceIngredients(possibleIngredients)

	for _, l := range lines {
		ingredients, _ := parseLine(l)
		for _, i := range ingredients {
			if _, ok := reducedIngredients[i]; !ok {
				count++
			}
		}
	}

	fmt.Println(count)
}

func parseLine(l string) ([]string, []string) {
	re := regexp.MustCompile(`(.*) \(contains (.*)\)$`)
	info := re.FindStringSubmatch(l)
	ingredients := strings.Split(info[1], " ")
	allergens := strings.Split(info[2], ", ")

	return ingredients, allergens
}

func reduceIngredients(possibleIngredients map[string]map[string]struct{}) map[string]string {
	reduced := map[string]string{}
	numberOfAllergens := len(possibleIngredients)
	for len(reduced) != numberOfAllergens {
		for a, ings := range possibleIngredients {
			if len(ings) == 1 {
				for i := range ings {
					reduced[i] = a
				}
			}
		}

		for _, a := range reduced {
			delete(possibleIngredients, a)
		}

		for a, ings := range possibleIngredients {
			toDel := []string{}
			for i := range ings {
				if _, ok := reduced[i]; ok {
					toDel = append(toDel, i)
				}
			}

			for _, i := range toDel {
				delete(ings, i)
			}

			possibleIngredients[a] = ings
		}
	}
	return reduced
}

func B(lines []string) {
	count := 0

	possibleIngredients := map[string]map[string]struct{}{} // allergen -> ingredients

	for _, l := range lines {
		ingredients, allergens := parseLine(l)

		for _, a := range allergens {
			if _, ok := possibleIngredients[a]; !ok {
				possibleIngredients[a] = map[string]struct{}{}
				for _, i := range ingredients {
					possibleIngredients[a][i] = struct{}{}
				}
			} else {
				temp := map[string]struct{}{}
				for _, i := range ingredients {
					temp[i] = struct{}{}
				}

				toRemove := []string{}
				for i := range temp {
					if _, ok := possibleIngredients[a][i]; !ok {
						toRemove = append(toRemove, i)
					}
				}

				for _, s := range toRemove {
					delete(temp, s)
				}
				possibleIngredients[a] = temp
			}
		}
	}

	reducedIngredients := reduceIngredients(possibleIngredients)

	fmt.Println(getCanonicalList(reducedIngredients))

	fmt.Println(count)
}

func getCanonicalList(reducedIngredients map[string]string) string {
	list := []string{}

	for i := range reducedIngredients {
		list = append(list, i)
	}

	sort.Slice(list, func(i, j int) bool {
		return reducedIngredients[list[i]] < reducedIngredients[list[j]]
	})

	return strings.Join(list, ",")
}
