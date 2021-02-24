package main

import (
	"fmt"
	"sort"
)

type Delivery struct {
	teamSize int
	pizzaIds []int
}

//const fileName = "a_example"
//const fileName = "b_little_bit_of_everything"
//const fileName = "c_many_ingredients"
//const fileName = "d_many_pizzas"
const fileName = "e_many_teams"

func main() {
	teams, allPizzas, ingrCount := ReadInput(fileName + ".in")
	fmt.Println(teams, len(allPizzas), ingrCount)

	sort.Slice(allPizzas, func(i, j int) bool {
		return allPizzas[i].score < allPizzas[j].score
	})

	deliveries := make([]Delivery, 0)

	for {
		delivPizzas := make([]Pizza, 1, 4)
		delivPizzaIds := make([]int, 1, 4)
		delivPizzas[0] = allPizzas[0]
		delivPizzaIds[0] = int(allPizzas[0].id)
		allPizzas = allPizzas[1:]

		for {
			nextPizzaIndex, found := findMatch(delivPizzas, allPizzas)

			if !found {
				break
			}

			delivPizzas = append(delivPizzas, allPizzas[nextPizzaIndex])
			delivPizzaIds = append(delivPizzaIds, int(allPizzas[nextPizzaIndex].id))
			allPizzas = rem(allPizzas, nextPizzaIndex)

			if len(delivPizzas) == 5 {
				fmt.Println(teams, len(allPizzas), ingrCount)
				panic("Shouldn't happen")
			}

			if teams[len(delivPizzas)-2] > 0 {
				break
			}
		}

		deliveries = append(deliveries, Delivery{
			teamSize: len(delivPizzas),
			pizzaIds: delivPizzaIds,
		})

		teams[len(delivPizzas)-2] = teams[len(delivPizzas)-2] - 1

		if noMoreDeliveries(len(allPizzas), teams) {
			break
		}
	}

	fmt.Println(deliveries)
	write(fileName+".out", deliveries)
}

// return index in available array
func findMatch(pizza []Pizza, available []Pizza) (int, bool) {
	// todo implement
	return len(available) - 1, true
}

func noMoreDeliveries(leftPizzas int, teams [3]uint) bool {
	if leftPizzas < 2 {
		return true
	} else if leftPizzas < 3 && teams[0] == 0 {
		return true
	} else if leftPizzas < 4 && teams[1] == 0 && teams[0] == 0 {
		return true
	} else if leftPizzas < 5 && teams[2] == 0 && teams[1] == 0 && teams[0] == 0 {
		return true
	} else if teams[2] == 0 && teams[1] == 0 && teams[0] == 0 {
		return true
	}

	return false
}

func rem(slice []Pizza, i int) []Pizza {
	return append(slice[:i], slice[i+1:]...)
}
