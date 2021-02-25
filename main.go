package main

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

		if len(intersection.in) == 0 {
			s.lights = append(s.lights, TrafficLight{
				streetName: intersection.in[0].title,
				time:       1,
			})
		} else {

			countMap := make(map[Street]int)

			for _, street := range intersection.in {
				if stritiNaInti[street.title] > 0 {
					countMap[street] = stritiNaInti[street.title]
				}
			}

			ratio := CalculateTrafficRate(countMap)
			ratios := NormalizeRatio(ratio)
			secondOrder := make([]TrafficLight, 0)
			for _, street := range intersection.in {
				if stritiNaInti[street.title] > 0 {
					countMap[street] = stritiNaInti[street.title]

					if _, ok := firstHits[street.title]; ok {
						s.lights = append(s.lights, TrafficLight{
							streetName: street.title,
							time:       ratios[street],
						})
					} else {
						secondOrder = append(secondOrder, TrafficLight{
							streetName: street.title,
							time:       ratios[street],
						})
					}
				}
			}

			for _, so := range secondOrder {
				s.lights = append(s.lights, so)
			}
		}

		if len(s.lights) > 0 {
			schedules = append(schedules, s)
		}
	}

	//fmt.Println(schedules)
	write(fileName+".out", schedules)
}
