package main

import "fmt"

//const fileName = "a.txt"
//const fileName = "b.txt"
//const fileName = "c.txt"
//const fileName = "d.txt"
//const fileName = "e.txt"
const fileName = "f.txt"

type Schedule struct {
	intersectionId int
	lights         []TrafficLight
}

type TrafficLight struct {
	streetName string
	time       int
}

func main() {
	ReadInput(fileName)

	schedules := make([]Schedule, 0)

	for _, intersection := range intersections {
		s := Schedule{
			intersectionId: intersection.id,
			lights:         make([]TrafficLight, 0),
		}

		for _, street := range intersection.in {
			s.lights = append(s.lights, TrafficLight{
				streetName: street.title,
				time:       1,
			})
		}

		schedules = append(schedules, s)
	}

	fmt.Println(schedules)
	write(fileName+".out", schedules)
}
