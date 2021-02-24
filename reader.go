package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pizza struct {
	id          uint
	ingredients []Ingredient
	score       uint
}

type Ingredient uint

var lastIngredient Ingredient = 0
var ingredientIdMap map[string]Ingredient
var ingredientCount map[Ingredient]uint

func ReadInput(fileName string) ([3]uint, []Pizza, map[Ingredient]uint) {
	ingredientCount = make(map[Ingredient]uint)
	ingredientIdMap = make(map[string]Ingredient)
	inp, err := os.Open(fileName)
	reader := bufio.NewReader(inp)
	check(err)

	var pizzaCount, teamsOf2, teamsOf3, teamsOf4 uint
	_, err = fmt.Fscanf(reader, "%d %d %d %d\n", &pizzaCount, &teamsOf2, &teamsOf3, &teamsOf4)
	check(err)

	pizzas := make([]Pizza, 0, pizzaCount)
	for i := uint(0); i < pizzaCount; i++ {
		pizzas = append(pizzas, readPizza(reader, i))
	}

	err = inp.Close()
	check(err)

	for i := uint(0); i < pizzaCount; i++ {
		pizzas[i].score = calcPizzaScore(pizzas[i])
	}

	return [3]uint{teamsOf2, teamsOf3, teamsOf4}, pizzas, ingredientCount
}

func readPizza(reader *bufio.Reader, id uint) Pizza {
	line, _, err := reader.ReadLine()
	check(err)

	parts := strings.Split(string(line), " ")
	ingCount, err := strconv.Atoi(parts[0])
	check(err)

	return Pizza{
		id:          id,
		ingredients: mapIngrs(parts[1 : ingCount+1]),
	}
}

func mapIngrs(asString []string) []Ingredient {
	mapped := make([]Ingredient, len(asString), len(asString))

	for i := 0; i < len(asString); i++ {
		if id, ok := ingredientIdMap[asString[i]]; ok {
			mapped[i] = id
			ingredientCount[id] = ingredientCount[id] + 1
		} else {
			ingredientIdMap[asString[i]] = lastIngredient
			ingredientCount[lastIngredient] = 1
			mapped[i] = lastIngredient
			lastIngredient++
		}
	}

	return mapped
}

func calcPizzaScore(pizza Pizza) uint {
	var score uint = 0

	for _, iId := range pizza.ingredients {
		score += ingredientCount[iId]
	}

	return score
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
