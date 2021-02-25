package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Car struct {
	streetNames []string
}

type Intersection struct {
	id  int
	in  []Street
	out []Street
}

type Street struct {
	title  string
	length int
	from   *Intersection
	to     *Intersection
}

var intersections = make(map[int]*Intersection)
var cars = make([]Car, 0)

func ReadInput(fileName string) {
	inp, err := os.Open(fileName)
	reader := bufio.NewReader(inp)
	check(err)

	var D, I, S, V, F int
	_, err = fmt.Fscanf(reader, "%d %d %d %d %d\n", &D, &I, &S, &V, &F)
	check(err)

	for i := 0; i < S; i++ {
		line, _, err := reader.ReadLine()
		check(err)

		parts := strings.Split(string(line), " ")
		B, err := strconv.Atoi(parts[0])
		check(err)

		E, err := strconv.Atoi(parts[1])
		check(err)

		title := parts[2]
		length, err := strconv.Atoi(parts[3])
		check(err)

		bIntersection := getOrCreate(B)
		eIntersection := getOrCreate(E)

		s := Street{
			title:  title,
			length: length,
			from:   bIntersection,
			to:     eIntersection,
		}

		bIntersection.out = append(bIntersection.out, s)
		eIntersection.in = append(eIntersection.in, s)
	}

	for i := 0; i < V; i++ {
		line, _, err := reader.ReadLine()
		check(err)

		parts := strings.Split(string(line), " ")

		P, err := strconv.Atoi(parts[0])
		check(err)

		streets := make([]string, P, P)
		for j := 0; j < P; j++ {
			streets[j] = parts[j+1]
		}

		cars = append(cars, Car{streetNames: streets})
	}

	fmt.Println(cars)
	printGraph()
}

func getOrCreate(id int) *Intersection {
	i, ok := intersections[id]

	if ok {
		return i
	}

	newI := Intersection{
		id:  id,
		in:  make([]Street, 0),
		out: make([]Street, 0),
	}

	intersections[id] = &newI

	return &newI
}

func printGraph() {
	for _, v := range intersections {
		fmt.Printf("%d in: %+v  out: %+v\n", v.id, v.in, v.out)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
